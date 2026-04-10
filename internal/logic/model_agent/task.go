package model_agent

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// 健康检查任务
func (s *sModelAgent) HealthCheckTask(ctx context.Context) {

	logger.Info(ctx, "sModelAgent HealthCheckTask start")

	now := gtime.TimestampMilli()

	mutex := s.healthCheckRedsync.NewMutex(consts.TASK_MODEL_AGENT_HEALTH_CHECK_LOCK_KEY, redsync.WithExpiry(config.Cfg.ModelAgentHealthCheckTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sModelAgent HealthCheckTask", err)
		logger.Debugf(ctx, "sModelAgent HealthCheckTask end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sModelAgent HealthCheckTask lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sModelAgent HealthCheckTask unlock")
		}
		logger.Debugf(ctx, "sModelAgent HealthCheckTask end time: %d", gtime.TimestampMilli()-now)
	}()

	if len(config.Cfg.ModelAgentHealthCheckTask.ModelAgents) == 0 || len(config.Cfg.ModelAgentHealthCheckTask.Models) == 0 {
		return
	}

	for _, modelAgentId := range config.Cfg.ModelAgentHealthCheckTask.ModelAgents {

		modelAgent, err := dao.ModelAgent.FindById(ctx, modelAgentId)
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		if !modelAgent.IsEnableHealthCheck || (modelAgent.Status == 2 && !modelAgent.IsAutoDisabled) {
			continue
		}

		if config.Cfg.ModelAgentHealthCheckTask.SmartCheck {
			// 智能模式
			s.smartCheckAgent(ctx, modelAgent, config.Cfg.ModelAgentHealthCheckTask)
			continue
		}

		s.healthCheck(ctx, modelAgent, config.Cfg.ModelAgentHealthCheckTask.Models, config.Cfg.ModelAgentHealthCheckTask.TestMethod, config.Cfg.ModelAgentHealthCheckTask.BaseUrl, config.Cfg.ModelAgentHealthCheckTask.Key)

		// 统计周期内的检查结果
		resultKey := fmt.Sprintf(consts.TASK_MODEL_AGENT_HEALTH_CHECK_RESULT_KEY, modelAgentId)
		statPeriodMs := (config.Cfg.ModelAgentHealthCheckTask.StatPeriod * time.Minute).Milliseconds()
		cutoff := gtime.TimestampMilli() - statPeriodMs

		results, err := redis.LRange(ctx, resultKey, 0, -1)
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		// 统计周期内的失败和成功次数
		var failCount, successCount int64
		for _, r := range results {

			parts := strings.Split(gconv.String(r), ",")
			if len(parts) < 2 {
				continue
			}

			ts := gconv.Int64(parts[0])
			if ts < cutoff {
				continue
			}

			if parts[1] == "0" {
				failCount++
			} else {
				successCount++
			}
		}

		// 禁用逻辑: 当前正常 + 非永不禁用 + 失败次数达标
		if modelAgent.Status == 1 && !modelAgent.IsNeverDisable && config.Cfg.ModelAgentHealthCheckTask.DisableCount > 0 && failCount >= config.Cfg.ModelAgentHealthCheckTask.DisableCount {

			reason := fmt.Sprintf("健康检查失败, 统计周期%d分钟内失败%d次", config.Cfg.ModelAgentHealthCheckTask.StatPeriod, failCount)

			if err = dao.ModelAgent.UpdateById(ctx, modelAgentId, bson.M{
				"status":               2,
				"is_auto_disabled":     true,
				"auto_disabled_reason": reason,
			}); err != nil {
				logger.Error(ctx, err)
				continue
			}

			logger.Infof(ctx, "sModelAgent HealthCheckTask 模型代理[%s %s]已自动禁用, 原因: %s", modelAgent.Name, modelAgentId, reason)

			// 清除代理级别检查记录, 避免恢复后被历史记录再次禁用
			if _, err = redis.Del(ctx, resultKey); err != nil {
				logger.Error(ctx, err)
			}

			if newData, err := s.Detail(ctx, modelAgentId); err != nil {
				logger.Error(ctx, err)
			} else {
				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
					Action:  consts.ACTION_STATUS,
					OldData: modelAgent,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
				}
			}
		}

		// 恢复逻辑: 当前已自动禁用 + 开启自动恢复 + 成功次数达标
		if modelAgent.Status == 2 && modelAgent.IsAutoDisabled && config.Cfg.ModelAgentHealthCheckTask.AutoRecover && config.Cfg.ModelAgentHealthCheckTask.RecoverCount > 0 && successCount >= config.Cfg.ModelAgentHealthCheckTask.RecoverCount {

			if err = dao.ModelAgent.UpdateById(ctx, modelAgentId, bson.M{
				"status":               1,
				"is_auto_disabled":     false,
				"auto_disabled_reason": "",
			}); err != nil {
				logger.Error(ctx, err)
				continue
			}

			logger.Infof(ctx, "sModelAgent HealthCheckTask 模型代理[%s %s]已自动恢复, 统计周期%d分钟内成功%d次", modelAgent.Name, modelAgentId, config.Cfg.ModelAgentHealthCheckTask.StatPeriod, successCount)

			// 清除代理级别检查记录, 避免被历史记录再次影响
			if _, err = redis.Del(ctx, resultKey); err != nil {
				logger.Error(ctx, err)
			}

			if newData, err := s.Detail(ctx, modelAgentId); err != nil {
				logger.Error(ctx, err)
			} else {
				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
					Action:  consts.ACTION_STATUS,
					OldData: modelAgent,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	if _, err := redis.Set(ctx, consts.TASK_MODEL_AGENT_HEALTH_CHECK_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sModelAgent) healthCheck(ctx context.Context, modelAgent *entity.ModelAgent, configModels []string, testMethod int, baseUrl, key string) {

	// 取 configModels 与 modelAgent.Models 的交集作为测试模型
	agentModelSet := make(map[string]bool)
	for _, m := range modelAgent.Models {
		agentModelSet[m] = true
	}

	matchedModels := make([]string, 0)
	for _, m := range configModels {
		if agentModelSet[m] {
			matchedModels = append(matchedModels, m)
		}
	}

	// 构建异常模型集合
	abnormalModelSet := make(map[string]bool)
	for _, m := range modelAgent.AbnormalModels {
		abnormalModelSet[m] = true
	}

	// 测试模型 = 交集模型 + 异常模型(用于恢复检测)
	testModels := make([]string, 0, len(matchedModels)+len(modelAgent.AbnormalModels))
	testModels = append(testModels, matchedModels...)
	testModels = append(testModels, modelAgent.AbnormalModels...)

	if len(testModels) == 0 {
		return
	}

	agentResultKey := fmt.Sprintf(consts.TASK_MODEL_AGENT_HEALTH_CHECK_RESULT_KEY, modelAgent.Id)
	statPeriod := config.Cfg.ModelAgentHealthCheckTask.StatPeriod

	for _, modelId := range testModels {

		res, err := s.TestModel(ctx, model.ModelAgentTestModelReq{
			ModelAgentId: modelAgent.Id,
			ModelId:      modelId,
			TestMethod:   testMethod,
			BaseUrl:      baseUrl,
			Key:          key,
		})

		result := "1" // 成功
		if err != nil || (res != nil && !res.Result) {
			result = "0" // 失败
		}

		record := fmt.Sprintf("%d,%s", gtime.TimestampMilli(), result)

		maxLen := int64(statPeriod) * 100
		if maxLen < 100 {
			maxLen = 100
		}

		ttl := int64(statPeriod) * 60 * 2
		if ttl < 600 {
			ttl = 600
		}

		if !abnormalModelSet[modelId] || result == "1" {

			if _, err = redis.LPush(ctx, agentResultKey, record); err != nil {
				logger.Error(ctx, err)
			}

			if err = redis.LTrim(ctx, agentResultKey, 0, maxLen); err != nil {
				logger.Error(ctx, err)
			}

			if _, err = redis.Expire(ctx, agentResultKey, ttl); err != nil {
				logger.Error(ctx, err)
			}
		}

		// 所有测试模型的结果都计入模型级别统计
		modelResultKey := fmt.Sprintf(consts.TASK_MODEL_AGENT_HEALTH_CHECK_MODEL_RESULT_KEY, modelAgent.Id, modelId)
		if _, err = redis.LPush(ctx, modelResultKey, record); err != nil {
			logger.Error(ctx, err)
		}

		if err = redis.LTrim(ctx, modelResultKey, 0, maxLen); err != nil {
			logger.Error(ctx, err)
		}

		if _, err = redis.Expire(ctx, modelResultKey, ttl); err != nil {
			logger.Error(ctx, err)
		}
	}

	// 模型级别: 统计每个模型的失败/成功次数, 判断移除和恢复
	statPeriodMs := (statPeriod * time.Minute).Milliseconds()
	cutoff := gtime.TimestampMilli() - statPeriodMs

	failedFromModels := make([]string, 0)
	recoveredFromAbnormal := make([]string, 0)

	for _, modelId := range testModels {

		modelResultKey := fmt.Sprintf(consts.TASK_MODEL_AGENT_HEALTH_CHECK_MODEL_RESULT_KEY, modelAgent.Id, modelId)

		results, err := redis.LRange(ctx, modelResultKey, 0, -1)
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		var failCount, successCount int64
		for _, r := range results {

			parts := strings.Split(gconv.String(r), ",")
			if len(parts) < 2 {
				continue
			}

			ts := gconv.Int64(parts[0])
			if ts < cutoff {
				continue
			}

			if parts[1] == "0" {
				failCount++
			} else {
				successCount++
			}
		}

		if abnormalModelSet[modelId] {
			// 异常模型: 成功次数达标则恢复
			if config.Cfg.ModelAgentHealthCheckTask.RecoverModelCount > 0 && successCount >= config.Cfg.ModelAgentHealthCheckTask.RecoverModelCount {
				recoveredFromAbnormal = append(recoveredFromAbnormal, modelId)
				logger.Infof(ctx, "sModelAgent HealthCheckTask 模型代理[%s %s]模型[%s]已恢复, 统计周期%d分钟内成功%d次", modelAgent.Name, modelAgent.Id, modelId, statPeriod, successCount)
			}
		} else {
			// 正常模型: 代理正常时, 失败次数达标则移除; 代理已禁用时不移除模型(失败可能是代理本身的问题)
			if modelAgent.Status == 1 && config.Cfg.ModelAgentHealthCheckTask.RemoveModelCount > 0 && failCount >= config.Cfg.ModelAgentHealthCheckTask.RemoveModelCount {
				failedFromModels = append(failedFromModels, modelId)
				logger.Infof(ctx, "sModelAgent HealthCheckTask 模型代理[%s %s]模型[%s]已标记异常, 统计周期%d分钟内失败%d次", modelAgent.Name, modelAgent.Id, modelId, statPeriod, failCount)
			}
		}
	}

	// 移除异常模型: 从Models移到AbnormalModels
	if modelAgent.IsRemoveAbnormalModel && len(failedFromModels) > 0 {

		var err error
		modelAgent, err = dao.ModelAgent.FindOneAndUpdateById(ctx, modelAgent.Id, bson.M{
			"$pull": bson.M{"models": bson.M{"$in": failedFromModels}},
			"$push": bson.M{"abnormal_models": bson.M{"$each": failedFromModels}},
		})
		if err != nil {
			logger.Error(ctx, err)
		} else {

			// 清除已移除模型的检查记录, 避免恢复后被历史记录再次移除
			for _, modelId := range failedFromModels {
				modelResultKey := fmt.Sprintf(consts.TASK_MODEL_AGENT_HEALTH_CHECK_MODEL_RESULT_KEY, modelAgent.Id, modelId)
				if _, err := redis.Del(ctx, modelResultKey); err != nil {
					logger.Error(ctx, err)
				}
			}

			if newData, err := s.Detail(ctx, modelAgent.Id); err != nil {
				logger.Error(ctx, err)
			} else {
				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: modelAgent,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	// 恢复模型: 从AbnormalModels移回Models
	if len(recoveredFromAbnormal) > 0 {

		if err := dao.ModelAgent.UpdateById(ctx, modelAgent.Id, bson.M{
			"$push": bson.M{"models": bson.M{"$each": recoveredFromAbnormal}},
			"$pull": bson.M{"abnormal_models": bson.M{"$in": recoveredFromAbnormal}},
		}); err != nil {
			logger.Error(ctx, err)
		} else {

			// 清除已恢复模型的检查记录, 避免被历史记录再次移除
			for _, modelId := range recoveredFromAbnormal {
				modelResultKey := fmt.Sprintf(consts.TASK_MODEL_AGENT_HEALTH_CHECK_MODEL_RESULT_KEY, modelAgent.Id, modelId)
				if _, err := redis.Del(ctx, modelResultKey); err != nil {
					logger.Error(ctx, err)
				}
			}

			if newData, err := s.Detail(ctx, modelAgent.Id); err != nil {
				logger.Error(ctx, err)
			} else {
				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: modelAgent,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}
}

// 智能检查 - 代理级别
func (s *sModelAgent) smartCheckAgent(ctx context.Context, modelAgent *entity.ModelAgent, cfg *common.ModelAgentHealthCheckTask) {

	successCount, failCount := s.readAgentCallStats(ctx, modelAgent.Id)
	totalCount := successCount + failCount

	// 无 API 调用数据 → 回退传统检查 (代理级别 + 模型级别全部走传统逻辑)
	if totalCount == 0 {
		logger.Infof(ctx, "SmartCheck 代理[%s %s] 无API调用数据, 回退主动探测", modelAgent.Name, modelAgent.Id)
		s.healthCheck(ctx, modelAgent, cfg.Models, cfg.TestMethod, cfg.BaseUrl, cfg.Key)
		s.traditionalAgentStats(ctx, modelAgent, cfg)
		return
	}

	// 禁用判断: 代理正常 + 非永不禁用 + API 失败次数 >= DisableCount
	if modelAgent.Status == 1 && !modelAgent.IsNeverDisable && cfg.DisableCount > 0 && failCount >= cfg.DisableCount {

		logger.Infof(ctx, "SmartCheck 代理[%s %s] API失败次数%d达到阈值%d, 开始主动验证", modelAgent.Name, modelAgent.Id, failCount, cfg.DisableCount)

		if s.smartVerify(ctx, modelAgent, cfg) {

			reason := fmt.Sprintf("智能检查: 统计周期%d分钟内API调用失败%d次(阈值%d), 主动验证确认异常", cfg.StatPeriod, failCount, cfg.DisableCount)

			if err := dao.ModelAgent.UpdateById(ctx, modelAgent.Id, bson.M{
				"status":               2,
				"is_auto_disabled":     true,
				"auto_disabled_reason": reason,
			}); err != nil {
				logger.Error(ctx, err)
			} else {
				logger.Infof(ctx, "SmartCheck 模型代理[%s %s]已自动禁用, 原因: %s", modelAgent.Name, modelAgent.Id, reason)
				if newData, err := s.Detail(ctx, modelAgent.Id); err != nil {
					logger.Error(ctx, err)
				} else {
					if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
						Action:  consts.ACTION_STATUS,
						OldData: modelAgent,
						NewData: newData,
					}); err != nil {
						logger.Error(ctx, err)
					}
				}
			}
		} else {
			logger.Infof(ctx, "SmartCheck 代理[%s %s] 主动验证通过, 暂不操作", modelAgent.Name, modelAgent.Id)
		}

		s.clearAgentCallStats(ctx, modelAgent.Id)
		return
	}

	// 恢复判断: 代理已自动禁用 + 开启自动恢复 + API 成功次数 >= RecoverCount
	if modelAgent.Status == 2 && modelAgent.IsAutoDisabled && cfg.AutoRecover && cfg.RecoverCount > 0 && successCount >= cfg.RecoverCount {

		if err := dao.ModelAgent.UpdateById(ctx, modelAgent.Id, bson.M{
			"status":               1,
			"is_auto_disabled":     false,
			"auto_disabled_reason": "",
		}); err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Infof(ctx, "SmartCheck 模型代理[%s %s]已自动恢复, 统计周期%d分钟内API成功%d次", modelAgent.Name, modelAgent.Id, cfg.StatPeriod, successCount)
			if newData, err := s.Detail(ctx, modelAgent.Id); err != nil {
				logger.Error(ctx, err)
			} else {
				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
					Action:  consts.ACTION_STATUS,
					OldData: modelAgent,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
		s.clearAgentCallStats(ctx, modelAgent.Id)
	}

	// 模型级别判断
	s.smartCheckModels(ctx, modelAgent, cfg)
}

// 主动验证: 对匹配的模型各跑一次 TestModel, 有任一失败则视为验证失败
func (s *sModelAgent) smartVerify(ctx context.Context, modelAgent *entity.ModelAgent, cfg *common.ModelAgentHealthCheckTask) (failed bool) {

	agentModelSet := make(map[string]bool)
	for _, m := range modelAgent.Models {
		agentModelSet[m] = true
	}

	matchedModels := make([]string, 0)
	for _, m := range cfg.Models {
		if agentModelSet[m] {
			matchedModels = append(matchedModels, m)
		}
	}

	if len(matchedModels) == 0 {
		return false
	}

	for _, modelId := range matchedModels {
		res, err := s.TestModel(ctx, model.ModelAgentTestModelReq{
			ModelAgentId: modelAgent.Id,
			ModelId:      modelId,
			TestMethod:   cfg.TestMethod,
			BaseUrl:      cfg.BaseUrl,
			Key:          cfg.Key,
		})
		if err != nil || (res != nil && !res.Result) {
			return true
		}
	}

	return false
}

// 智能检查 - 模型级别
func (s *sModelAgent) smartCheckModels(ctx context.Context, modelAgent *entity.ModelAgent, cfg *common.ModelAgentHealthCheckTask) {

	agentModelSet := make(map[string]bool)
	for _, m := range modelAgent.Models {
		agentModelSet[m] = true
	}

	matchedModels := make([]string, 0)
	for _, m := range cfg.Models {
		if agentModelSet[m] {
			matchedModels = append(matchedModels, m)
		}
	}

	failedFromModels := make([]string, 0)
	recoveredFromAbnormal := make([]string, 0)

	// 正常模型: 检查是否需要移除
	if modelAgent.IsRemoveAbnormalModel && modelAgent.Status == 1 && cfg.RemoveModelCount > 0 {
		for _, modelId := range matchedModels {
			_, failCount := s.readModelCallStats(ctx, modelAgent.Id, modelId)
			if failCount >= cfg.RemoveModelCount {
				res, err := s.TestModel(ctx, model.ModelAgentTestModelReq{
					ModelAgentId: modelAgent.Id,
					ModelId:      modelId,
					TestMethod:   cfg.TestMethod,
					BaseUrl:      cfg.BaseUrl,
					Key:          cfg.Key,
				})
				if err != nil || (res != nil && !res.Result) {
					failedFromModels = append(failedFromModels, modelId)
					logger.Infof(ctx, "SmartCheck 模型代理[%s %s]模型[%s]已标记异常, API失败%d次, 主动验证确认", modelAgent.Name, modelAgent.Id, modelId, failCount)
				}
				s.clearModelCallStats(ctx, modelAgent.Id, modelId)
			}
		}
	}

	// 异常模型: 主动探测检查是否可恢复
	// 异常模型已从 Models 移出, API 不再路由到该模型, 必须通过主动探测判断恢复
	if cfg.RecoverModelCount > 0 {

		statPeriodMs := (cfg.StatPeriod * time.Minute).Milliseconds()
		cutoff := gtime.TimestampMilli() - statPeriodMs

		for _, modelId := range modelAgent.AbnormalModels {

			res, err := s.TestModel(ctx, model.ModelAgentTestModelReq{
				ModelAgentId: modelAgent.Id,
				ModelId:      modelId,
				TestMethod:   cfg.TestMethod,
				BaseUrl:      cfg.BaseUrl,
				Key:          cfg.Key,
			})

			if err == nil && (res == nil || res.Result) {
				modelResultKey := fmt.Sprintf(consts.TASK_MODEL_AGENT_HEALTH_CHECK_MODEL_RESULT_KEY, modelAgent.Id, modelId)
				record := fmt.Sprintf("%d,%s", gtime.TimestampMilli(), "1")
				if _, err = redis.LPush(ctx, modelResultKey, record); err != nil {
					logger.Error(ctx, err)
				}
			}
		}

		for _, modelId := range modelAgent.AbnormalModels {

			modelResultKey := fmt.Sprintf(consts.TASK_MODEL_AGENT_HEALTH_CHECK_MODEL_RESULT_KEY, modelAgent.Id, modelId)

			results, err := redis.LRange(ctx, modelResultKey, 0, -1)
			if err != nil {
				logger.Error(ctx, err)
				continue
			}

			var successCount int64
			for _, r := range results {
				parts := strings.Split(gconv.String(r), ",")
				if len(parts) < 2 {
					continue
				}
				ts := gconv.Int64(parts[0])
				if ts < cutoff {
					continue
				}
				if parts[1] == "1" {
					successCount++
				}
			}

			if successCount >= cfg.RecoverModelCount {
				recoveredFromAbnormal = append(recoveredFromAbnormal, modelId)
				logger.Infof(ctx, "SmartCheck 模型代理[%s %s]模型[%s]已恢复, 主动探测成功%d次", modelAgent.Name, modelAgent.Id, modelId, successCount)
				if _, err := redis.Del(ctx, modelResultKey); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	// 移除异常模型
	if modelAgent.IsRemoveAbnormalModel && len(failedFromModels) > 0 {

		var err error
		modelAgent, err = dao.ModelAgent.FindOneAndUpdateById(ctx, modelAgent.Id, bson.M{
			"$pull": bson.M{"models": bson.M{"$in": failedFromModels}},
			"$push": bson.M{"abnormal_models": bson.M{"$each": failedFromModels}},
		})
		if err != nil {
			logger.Error(ctx, err)
		} else {
			if newData, err := s.Detail(ctx, modelAgent.Id); err != nil {
				logger.Error(ctx, err)
			} else {
				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: modelAgent,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	// 恢复模型
	if len(recoveredFromAbnormal) > 0 {
		if err := dao.ModelAgent.UpdateById(ctx, modelAgent.Id, bson.M{
			"$push": bson.M{"models": bson.M{"$each": recoveredFromAbnormal}},
			"$pull": bson.M{"abnormal_models": bson.M{"$in": recoveredFromAbnormal}},
		}); err != nil {
			logger.Error(ctx, err)
		} else {
			if newData, err := s.Detail(ctx, modelAgent.Id); err != nil {
				logger.Error(ctx, err)
			} else {
				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: modelAgent,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}
}

func (s *sModelAgent) readAgentCallStats(ctx context.Context, agentId string) (success, fail int64) {
	successKey := fmt.Sprintf(consts.SMART_HEALTH_AGENT_SUCCESS_KEY, agentId)
	failKey := fmt.Sprintf(consts.SMART_HEALTH_AGENT_FAIL_KEY, agentId)
	successVal, _ := redis.GetInt64(ctx, successKey)
	failVal, _ := redis.GetInt64(ctx, failKey)
	return successVal, failVal
}

func (s *sModelAgent) readModelCallStats(ctx context.Context, agentId, modelId string) (success, fail int64) {
	successKey := fmt.Sprintf(consts.SMART_HEALTH_MODEL_SUCCESS_KEY, agentId, modelId)
	failKey := fmt.Sprintf(consts.SMART_HEALTH_MODEL_FAIL_KEY, agentId, modelId)
	successVal, _ := redis.GetInt64(ctx, successKey)
	failVal, _ := redis.GetInt64(ctx, failKey)
	return successVal, failVal
}

func (s *sModelAgent) clearAgentCallStats(ctx context.Context, agentId string) {
	if _, err := redis.Del(ctx, fmt.Sprintf(consts.SMART_HEALTH_AGENT_SUCCESS_KEY, agentId)); err != nil {
		logger.Error(ctx, err)
	}
	if _, err := redis.Del(ctx, fmt.Sprintf(consts.SMART_HEALTH_AGENT_FAIL_KEY, agentId)); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sModelAgent) clearModelCallStats(ctx context.Context, agentId, modelId string) {
	if _, err := redis.Del(ctx, fmt.Sprintf(consts.SMART_HEALTH_MODEL_SUCCESS_KEY, agentId, modelId)); err != nil {
		logger.Error(ctx, err)
	}
	if _, err := redis.Del(ctx, fmt.Sprintf(consts.SMART_HEALTH_MODEL_FAIL_KEY, agentId, modelId)); err != nil {
		logger.Error(ctx, err)
	}
}

// 传统模式完整统计 (回退时使用, 包含代理级别和模型级别)
func (s *sModelAgent) traditionalAgentStats(ctx context.Context, modelAgent *entity.ModelAgent, cfg *common.ModelAgentHealthCheckTask) {

	// 代理级别: 统计周期内的检查结果
	resultKey := fmt.Sprintf(consts.TASK_MODEL_AGENT_HEALTH_CHECK_RESULT_KEY, modelAgent.Id)
	statPeriodMs := (cfg.StatPeriod * time.Minute).Milliseconds()
	cutoff := gtime.TimestampMilli() - statPeriodMs

	results, err := redis.LRange(ctx, resultKey, 0, -1)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	var failCount, successCount int64
	for _, r := range results {
		parts := strings.Split(gconv.String(r), ",")
		if len(parts) < 2 {
			continue
		}
		ts := gconv.Int64(parts[0])
		if ts < cutoff {
			continue
		}
		if parts[1] == "0" {
			failCount++
		} else {
			successCount++
		}
	}

	// 禁用逻辑
	if modelAgent.Status == 1 && !modelAgent.IsNeverDisable && cfg.DisableCount > 0 && failCount >= cfg.DisableCount {

		reason := fmt.Sprintf("健康检查失败, 统计周期%d分钟内失败%d次", cfg.StatPeriod, failCount)

		if err = dao.ModelAgent.UpdateById(ctx, modelAgent.Id, bson.M{
			"status":               2,
			"is_auto_disabled":     true,
			"auto_disabled_reason": reason,
		}); err != nil {
			logger.Error(ctx, err)
			return
		}

		logger.Infof(ctx, "sModelAgent HealthCheckTask 模型代理[%s %s]已自动禁用, 原因: %s", modelAgent.Name, modelAgent.Id, reason)

		if _, err = redis.Del(ctx, resultKey); err != nil {
			logger.Error(ctx, err)
		}

		if newData, err := s.Detail(ctx, modelAgent.Id); err != nil {
			logger.Error(ctx, err)
		} else {
			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
				Action:  consts.ACTION_STATUS,
				OldData: modelAgent,
				NewData: newData,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}
	}

	// 恢复逻辑
	if modelAgent.Status == 2 && modelAgent.IsAutoDisabled && cfg.AutoRecover && cfg.RecoverCount > 0 && successCount >= cfg.RecoverCount {

		if err = dao.ModelAgent.UpdateById(ctx, modelAgent.Id, bson.M{
			"status":               1,
			"is_auto_disabled":     false,
			"auto_disabled_reason": "",
		}); err != nil {
			logger.Error(ctx, err)
			return
		}

		logger.Infof(ctx, "sModelAgent HealthCheckTask 模型代理[%s %s]已自动恢复, 统计周期%d分钟内成功%d次", modelAgent.Name, modelAgent.Id, cfg.StatPeriod, successCount)

		if _, err = redis.Del(ctx, resultKey); err != nil {
			logger.Error(ctx, err)
		}

		if newData, err := s.Detail(ctx, modelAgent.Id); err != nil {
			logger.Error(ctx, err)
		} else {
			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
				Action:  consts.ACTION_STATUS,
				OldData: modelAgent,
				NewData: newData,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}
	}
}
