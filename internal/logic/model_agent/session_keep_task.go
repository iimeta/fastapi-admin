package model_agent

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/model/common"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// 会话保持清理任务
func (s *sModelAgent) SessionKeepCleanupTask(ctx context.Context) {

	logger.Info(ctx, "sModelAgent SessionKeepCleanupTask start")

	now := gtime.TimestampMilli()

	mutex := s.healthCheckRedsync.NewMutex(consts.TASK_MODEL_AGENT_SESSION_KEEP_CLEAN_LOCK_KEY, redsync.WithExpiry(config.Cfg.ModelAgentSessionKeepTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sModelAgent SessionKeepCleanupTask", err)
		logger.Debugf(ctx, "sModelAgent SessionKeepCleanupTask end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sModelAgent SessionKeepCleanupTask lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sModelAgent SessionKeepCleanupTask unlock")
		}
		logger.Debugf(ctx, "sModelAgent SessionKeepCleanupTask end time: %d", gtime.TimestampMilli()-now)
	}()

	agentSetKeys, err := redis.Keys(ctx, fmt.Sprintf(consts.SESSION_KEEP_AGENT_SET_PREFIX, "*"))
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	var cleaned int64

	for _, agentSetKey := range agentSetKeys {

		agentId, ok := parseSessionKeepAgentSetKey(agentSetKey)
		if !ok {
			continue
		}

		ttlSeconds, ttlErr := s.sessionKeepTTLSecondsByAgent(ctx, agentId)
		if ttlErr != nil {
			logger.Error(ctx, ttlErr)
			continue
		}

		if ttlSeconds <= 0 {
			continue
		}

		items, rangeErr := redis.ZRangeWithScores(ctx, agentSetKey, 0, -1)
		if rangeErr != nil {
			logger.Error(ctx, rangeErr)
			continue
		}

		expireBefore := time.Now().Unix() - ttlSeconds
		for _, item := range items {

			sk := recoverSessionKeepMember(item.Member)
			if sk == nil {
				if _, remErr := redis.ZRem(ctx, agentSetKey, item.Member); remErr != nil {
					logger.Error(ctx, remErr)
					continue
				}
				if _, remErr := redis.ZRem(ctx, consts.SESSION_KEEP_GLOBAL_SET, item.Member); remErr != nil {
					logger.Error(ctx, remErr)
				}
				cleaned++
				continue
			}

			if int64(item.Score) > expireBefore {
				continue
			}

			removed, cleanErr := s.cleanupSessionKeepEntry(ctx, sk, agentId)
			if cleanErr != nil {
				logger.Error(ctx, cleanErr)
				continue
			}
			if removed {
				cleaned++
			}
		}
	}

	globalMembers, err := redis.ZRange(ctx, consts.SESSION_KEEP_GLOBAL_SET, 0, -1)
	if err != nil {
		logger.Error(ctx, err)
	} else {
		for _, member := range globalMembers {

			sk := recoverSessionKeepMember(member)
			if sk == nil {
				if _, remErr := redis.ZRem(ctx, consts.SESSION_KEEP_GLOBAL_SET, member); remErr != nil {
					logger.Error(ctx, remErr)
					continue
				}
				cleaned++
				continue
			}

			valueKey := fmt.Sprintf(consts.SESSION_KEEP_VALUE_PREFIX, sk.raw)
			value, getErr := redis.GetStr(ctx, valueKey)
			if getErr != nil {
				logger.Error(ctx, getErr)
				continue
			}

			if value != "" {
				continue
			}

			if _, remErr := redis.ZRem(ctx, consts.SESSION_KEEP_GLOBAL_SET, member); remErr != nil {
				logger.Error(ctx, remErr)
				continue
			}

			userSetKeys, keyErr := redis.Keys(ctx, fmt.Sprintf(consts.SESSION_KEEP_USER_SET_SCAN_BY_USER, sk.userId))
			if keyErr != nil {
				logger.Error(ctx, keyErr)
				continue
			}

			for _, key := range userSetKeys {
				if _, remErr := redis.ZRem(ctx, key, member); remErr != nil {
					logger.Error(ctx, remErr)
				}
			}

			cleaned++
		}
	}

	logger.Infof(ctx, "sModelAgent SessionKeepCleanupTask cleaned: %d", cleaned)

	if _, err = redis.Set(ctx, consts.TASK_MODEL_AGENT_SESSION_KEEP_CLEAN_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sModelAgent) cleanupSessionKeepEntry(ctx context.Context, sk *sessionKeepMember, agentId string) (bool, error) {

	member := sk.raw
	valueKey := fmt.Sprintf(consts.SESSION_KEEP_VALUE_PREFIX, sk.raw)
	userSetKey := fmt.Sprintf(consts.SESSION_KEEP_USER_AGENT_SET, sk.userId, agentId)
	failKey := fmt.Sprintf(consts.SESSION_KEEP_FAIL_PREFIX, sk.raw, agentId)

	currentValue, err := redis.GetStr(ctx, valueKey)
	if err != nil {
		return false, err
	}

	currentAgentId := currentValue
	if idx := strings.Index(currentValue, ":"); idx > 0 {
		currentAgentId = currentValue[:idx]
	}

	if _, err = redis.ZRem(ctx, fmt.Sprintf(consts.SESSION_KEEP_AGENT_SET_PREFIX, agentId), member); err != nil {
		return false, err
	}

	if _, err = redis.ZRem(ctx, userSetKey, member); err != nil {
		return false, err
	}

	if currentValue == "" || currentAgentId == agentId {
		if _, err = redis.ZRem(ctx, consts.SESSION_KEEP_GLOBAL_SET, member); err != nil {
			return false, err
		}
	}

	deleteKeys := make([]string, 0, 2)
	if currentValue == "" || currentAgentId == agentId {
		deleteKeys = append(deleteKeys, valueKey)
	}
	deleteKeys = append(deleteKeys, failKey)

	keyFailKeys, keyFailErr := redis.Keys(ctx, fmt.Sprintf(consts.SESSION_KEEP_KEY_FAIL_SCAN_BY_ENTRY, sk.raw, agentId))
	if keyFailErr == nil && len(keyFailKeys) > 0 {
		deleteKeys = append(deleteKeys, keyFailKeys...)
	}

	if len(deleteKeys) > 0 {
		if _, err = redis.Del(ctx, deleteKeys...); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *sModelAgent) sessionKeepTTLSecondsByAgent(ctx context.Context, agentId string) (int64, error) {

	cfg := sessionKeepConfigCopy(config.Cfg.ModelAgentSessionKeep)
	if agentId == "" {
		return sessionKeepTTLSeconds(cfg), nil
	}

	modelAgent, err := dao.ModelAgent.FindById(ctx, agentId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return sessionKeepTTLSeconds(cfg), nil
		}
		return 0, err
	}

	if modelAgent != nil && modelAgent.IsEnableSessionKeep && modelAgent.SessionKeepConfig != nil {
		cfg = mergeSessionKeepConfig(cfg, modelAgent.SessionKeepConfig)
	}

	return sessionKeepTTLSeconds(cfg), nil
}

func sessionKeepConfigCopy(cfg *common.ModelAgentSessionKeep) *common.ModelAgentSessionKeep {
	if cfg == nil {
		return &common.ModelAgentSessionKeep{}
	}

	copyCfg := *cfg
	return &copyCfg
}

func mergeSessionKeepConfig(base *common.ModelAgentSessionKeep, override *common.ModelAgentSessionKeep) *common.ModelAgentSessionKeep {

	cfg := sessionKeepConfigCopy(base)
	if override == nil {
		return cfg
	}

	cfg.Open = true

	if override.Mode != "" {
		cfg.Mode = override.Mode
	}
	if override.Ttl > 0 {
		cfg.Ttl = override.Ttl
	}
	if override.FailTtl > 0 {
		cfg.FailTtl = override.FailTtl
	}
	if override.FailSwitchThreshold > 0 {
		cfg.FailSwitchThreshold = override.FailSwitchThreshold
	}
	if override.UserLimit > 0 {
		cfg.UserLimit = override.UserLimit
	}
	if override.AgentLimit > 0 {
		cfg.AgentLimit = override.AgentLimit
	}
	if override.GlobalLimit > 0 {
		cfg.GlobalLimit = override.GlobalLimit
	}
	if len(override.Rules) > 0 {
		cfg.Rules = override.Rules
	}
	if override.EnableSystemPromptHash {
		cfg.EnableSystemPromptHash = true
	}

	return cfg
}

func sessionKeepTTLSeconds(cfg *common.ModelAgentSessionKeep) int64 {
	if cfg == nil || !cfg.Open || cfg.Ttl <= 0 {
		return 0
	}
	return int64(cfg.Ttl)
}

func parseSessionKeepAgentSetKey(key string) (string, bool) {

	// 对应 consts.SESSION_KEEP_AGENT_SET_PREFIX
	const prefix = "session:agent:set:"
	if !strings.HasPrefix(key, prefix) {
		return "", false
	}

	agentId := strings.TrimPrefix(key, prefix)
	if agentId == "" {
		return "", false
	}

	return agentId, true
}

type sessionKeepMember struct {
	raw    string
	userId int
}

func recoverSessionKeepMember(raw string) *sessionKeepMember {

	if strings.HasPrefix(raw, "u:") {
		// user mode: "u:{userId}:m:{modelName}"
		parts := strings.SplitN(raw, ":", 4)
		if len(parts) >= 4 {
			userId, _ := strconv.Atoi(parts[1])
			if userId > 0 {
				return &sessionKeepMember{raw: raw, userId: userId}
			}
		}
		return nil
	}

	if strings.HasPrefix(raw, "r:") {
		// rule mode: "r:{userId}:{ruleName}:{modelName}:{value}"
		parts := strings.SplitN(raw, ":", 3)
		if len(parts) >= 3 {
			userId, _ := strconv.Atoi(parts[1])
			if userId > 0 {
				return &sessionKeepMember{raw: raw, userId: userId}
			}
		}
		return nil
	}

	return nil
}
