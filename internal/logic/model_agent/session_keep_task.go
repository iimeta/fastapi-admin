package model_agent

import (
	"context"
	"errors"
	"fmt"
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

	agentSetKeys, err := redis.Keys(ctx, "session:agent:set:*")
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

			userId, modelName, memberOk := parseSessionKeepMember(item.Member)
			if !memberOk {
				if _, remErr := redis.ZRem(ctx, agentSetKey, item.Member); remErr != nil {
					logger.Error(ctx, remErr)
					continue
				}
				if _, remErr := redis.ZRem(ctx, "session:agent:global", item.Member); remErr != nil {
					logger.Error(ctx, remErr)
				}
				cleaned++
				continue
			}

			if int64(item.Score) > expireBefore {
				continue
			}

			removed, cleanErr := s.cleanupSessionKeepEntry(ctx, userId, modelName, agentId)
			if cleanErr != nil {
				logger.Error(ctx, cleanErr)
				continue
			}
			if removed {
				cleaned++
			}
		}
	}

	globalMembers, err := redis.ZRange(ctx, "session:agent:global", 0, -1)
	if err != nil {
		logger.Error(ctx, err)
	} else {
		for _, member := range globalMembers {

			userId, modelName, ok := parseSessionKeepMember(member)
			if !ok {
				if _, remErr := redis.ZRem(ctx, "session:agent:global", member); remErr != nil {
					logger.Error(ctx, remErr)
					continue
				}
				cleaned++
				continue
			}

			valueKey := fmt.Sprintf("session:agent:u:%d:m:%s", userId, modelName)
			agentId, getErr := redis.GetStr(ctx, valueKey)
			if getErr != nil {
				logger.Error(ctx, getErr)
				continue
			}

			if agentId != "" {
				continue
			}

			if _, remErr := redis.ZRem(ctx, "session:agent:global", member); remErr != nil {
				logger.Error(ctx, remErr)
				continue
			}

			userSetKeys, keyErr := redis.Keys(ctx, fmt.Sprintf("session:agent:user:set:%d:a:*", userId))
			if keyErr != nil {
				logger.Error(ctx, keyErr)
				continue
			}

			for _, key := range userSetKeys {
				if _, remErr := redis.ZRem(ctx, key, modelName); remErr != nil {
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

func (s *sModelAgent) cleanupSessionKeepEntry(ctx context.Context, userId int, modelName, agentId string) (bool, error) {

	member := fmt.Sprintf("%d:%s", userId, modelName)
	valueKey := fmt.Sprintf("session:agent:u:%d:m:%s", userId, modelName)
	userSetKey := fmt.Sprintf("session:agent:user:set:%d:a:%s", userId, agentId)
	failKey := fmt.Sprintf("session:agent:fail:u:%d:m:%s:a:%s", userId, modelName, agentId)

	currentAgentId, err := redis.GetStr(ctx, valueKey)
	if err != nil {
		return false, err
	}

	if _, err = redis.ZRem(ctx, fmt.Sprintf("session:agent:set:%s", agentId), member); err != nil {
		return false, err
	}

	if _, err = redis.ZRem(ctx, userSetKey, modelName); err != nil {
		return false, err
	}

	if currentAgentId == "" || currentAgentId == agentId {
		if _, err = redis.ZRem(ctx, "session:agent:global", member); err != nil {
			return false, err
		}
	}

	deleteKeys := make([]string, 0, 2)
	if currentAgentId == "" || currentAgentId == agentId {
		deleteKeys = append(deleteKeys, valueKey)
	}
	deleteKeys = append(deleteKeys, failKey)

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

func parseSessionKeepMember(member string) (int, string, bool) {

	parts := strings.SplitN(member, ":", 2)
	if len(parts) != 2 {
		return 0, "", false
	}

	var userId int
	if _, err := fmt.Sscanf(parts[0], "%d", &userId); err != nil || userId <= 0 {
		return 0, "", false
	}

	return userId, parts[1], true
}
