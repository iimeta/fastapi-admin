package task_video

import (
	"context"
	"slices"
	"strings"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	endpointVideoGenerations = "/v1/videos"
	endpointVideoRemix       = "/v1/videos"
)

// 失败重试且需要重新提交时, 按 API 同样规则实时重选模型代理和密钥
// 分组开启模型代理则走分组池(再按代理绑定的模型过滤), 否则走模型自己绑定的代理
// 查不到可用代理/密钥时返回错误, 由调用方继续用当前代理重试
func (s *sTaskVideo) switchRetryAgent(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, errorAgentIds, errorKeys []string, attempt int) (agentChanged bool, err error) {

	agent, key, err := s.pickRetryAgentAndKey(ctx, taskVideo, logVideo, errorAgentIds, errorKeys, attempt)
	if err != nil {
		return false, err
	}

	if agent.Id == logVideo.ModelAgentId && key.Key == logVideo.Key {
		return false, nil
	}

	agentChanged = agent.Id != logVideo.ModelAgentId

	logger.Infof(ctx, "sTaskVideo switchRetryAgent task: %s model agent %s -> %s", taskVideo.Id, logVideo.ModelAgentId, agent.Id)

	snapshot := slimModelAgent(agent)

	logVideo.ModelAgentId = agent.Id
	logVideo.ModelAgent = snapshot
	logVideo.Key = key.Key

	taskVideo.ModelAgentId = agent.Id
	taskVideo.ModelAgent = snapshot
	filter := videoWorkerFilter(taskVideo)
	taskVideo.JobId = videoJobResubmit

	if _, err = dao.TaskVideo.FindOneAndUpdate(ctx, filter, bson.M{
		"job_id":         videoJobResubmit,
		"model_agent_id": agent.Id,
		"model_agent":    snapshot,
	}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, errors.New("task already handled")
		}
		logger.Error(ctx, err)
		return false, err
	}

	if err = dao.LogVideo.UpdateById(ctx, logVideo.Id, bson.M{
		"model_agent_id": agent.Id,
		"model_agent":    snapshot,
		"key":            key.Key,
	}); err != nil {
		logger.Error(ctx, err)
	}

	return agentChanged, nil
}

func (s *sTaskVideo) pickRetryAgentAndKey(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, errorAgentIds, errorKeys []string, attempt int) (*entity.ModelAgent, *entity.Key, error) {

	agents, lbStrategy, err := s.listCandidateAgents(ctx, taskVideo, logVideo)
	if err != nil {
		return nil, nil, err
	}
	if len(agents) == 0 {
		return nil, nil, errors.New("no available model agent")
	}

	preferred := filterAgentsByIds(agents, errorAgentIds)
	if len(preferred) == 0 {
		preferred = agents
	}

	if agent, key := s.pickAgentWithKey(ctx, preferred, lbStrategy, logVideo.ModelAgentId, errorKeys, attempt); agent != nil && key != nil {
		return agent, key, nil
	}

	if agent, key := s.pickAgentWithKey(ctx, agents, lbStrategy, "", nil, attempt); agent != nil && key != nil {
		return agent, key, nil
	}

	return nil, nil, errors.New("no available model agent key")
}

func (s *sTaskVideo) pickAgentWithKey(ctx context.Context, agents []*entity.ModelAgent, lbStrategy int, currentAgentId string, errorKeys []string, attempt int) (*entity.ModelAgent, *entity.Key) {

	remaining := append([]*entity.ModelAgent(nil), agents...)

	for len(remaining) > 0 {

		agent := pickOneAgent(remaining, lbStrategy, currentAgentId, attempt)
		if agent == nil {
			break
		}

		key, err := s.pickRetryKey(ctx, agent, errorKeys, attempt)
		if err != nil {
			logger.Errorf(ctx, "sTaskVideo pickRetryKey agent: %s, error: %v", agent.Id, err)
		} else if key != nil {
			return agent, key
		}

		remaining = filterAgentsByIds(remaining, []string{agent.Id})
		currentAgentId = ""
		attempt++
	}

	return nil, nil
}

func (s *sTaskVideo) listCandidateAgents(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo) ([]*entity.ModelAgent, int, error) {

	modelId := logVideo.RealModelId
	if modelId == "" {
		modelId = logVideo.ModelId
	}
	if modelId == "" {
		return nil, 1, errors.New("model id is empty")
	}

	endpoint := endpointVideoGenerations
	if taskVideo.Action == "remix" {
		endpoint = endpointVideoRemix
	}

	lbStrategy := 1

	if logVideo.Spend.GroupId != "" {

		group, err := dao.Group.FindById(ctx, logVideo.Spend.GroupId)
		if err != nil {
			if !errors.Is(err, mongo.ErrNoDocuments) {
				logger.Error(ctx, err)
			}
		} else if group != nil && group.IsEnableModelAgent {

			if group.LbStrategy > 0 {
				lbStrategy = group.LbStrategy
			}

			if len(group.ModelAgents) == 0 {
				return nil, lbStrategy, nil
			}

			list, err := dao.ModelAgent.FindByIds(ctx, group.ModelAgents)
			if err != nil {
				return nil, lbStrategy, err
			}

			agents := make([]*entity.ModelAgent, 0, len(list))
			for _, agent := range list {
				if agent.Status != 1 {
					continue
				}
				if !slices.Contains(agent.Models, modelId) && !slices.Contains(agent.AbnormalModels, modelId) {
					continue
				}
				if !matchAgentEndpoint(agent, endpoint) {
					continue
				}
				agents = append(agents, agent)
			}

			return agents, lbStrategy, nil
		}
	}

	if m, err := dao.Model.FindById(ctx, modelId); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
		}
	} else if m != nil && m.LbStrategy > 0 {
		lbStrategy = m.LbStrategy
	}

	list, err := dao.ModelAgent.Find(ctx, bson.M{
		"models": bson.M{"$in": []string{modelId}},
		"status": 1,
	})
	if err != nil {
		return nil, lbStrategy, err
	}

	agents := make([]*entity.ModelAgent, 0, len(list))
	for _, agent := range list {
		if !matchAgentEndpoint(agent, endpoint) {
			continue
		}
		agents = append(agents, agent)
	}

	return agents, lbStrategy, nil
}

func (s *sTaskVideo) pickRetryKey(ctx context.Context, agent *entity.ModelAgent, errorKeys []string, attempt int) (*entity.Key, error) {

	keys, err := dao.Key.Find(ctx, bson.M{
		"model_agents": agent.Id,
		"status":       1,
	})
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return nil, errors.New("no available model agent key")
	}

	preferred := filterKeysByValue(keys, errorKeys)
	if len(preferred) == 0 {
		preferred = keys
	}

	return pickOneKey(preferred, agent.LbStrategy, attempt), nil
}

func matchAgentEndpoint(agent *entity.ModelAgent, endpoint string) bool {
	if len(agent.Endpoints) == 0 || slices.Contains(agent.Endpoints, endpoint) {
		return true
	}
	for _, ep := range agent.Endpoints {
		if strings.Contains(ep, "video") || strings.Contains(ep, "generations/tasks") {
			return true
		}
	}
	return false
}

func filterAgentsByIds(agents []*entity.ModelAgent, excludeIds []string) []*entity.ModelAgent {

	if len(excludeIds) == 0 {
		return agents
	}

	filtered := make([]*entity.ModelAgent, 0, len(agents))
	for _, agent := range agents {
		if !slices.Contains(excludeIds, agent.Id) {
			filtered = append(filtered, agent)
		}
	}

	return filtered
}

func filterKeysByValue(keys []*entity.Key, excludeKeys []string) []*entity.Key {

	if len(excludeKeys) == 0 {
		return keys
	}

	filtered := make([]*entity.Key, 0, len(keys))
	for _, key := range keys {
		if !slices.Contains(excludeKeys, key.Key) {
			filtered = append(filtered, key)
		}
	}

	return filtered
}

func pickOneAgent(agents []*entity.ModelAgent, lbStrategy int, currentAgentId string, attempt int) *entity.ModelAgent {

	if len(agents) == 0 {
		return nil
	}

	list := agents
	if currentAgentId != "" && len(agents) > 1 {
		filtered := filterAgentsByIds(agents, []string{currentAgentId})
		if len(filtered) > 0 {
			list = filtered
		}
	}

	if len(list) == 1 {
		return list[0]
	}

	if lbStrategy == 2 {
		return pickAgentByWeight(list)
	}

	if attempt < 0 {
		attempt = 0
	}

	return list[attempt%len(list)]
}

func pickOneKey(keys []*entity.Key, lbStrategy int, attempt int) *entity.Key {

	if len(keys) == 0 {
		return nil
	}
	if len(keys) == 1 {
		return keys[0]
	}

	if lbStrategy == 2 {
		return pickKeyByWeight(keys)
	}

	if attempt < 0 {
		attempt = 0
	}

	return keys[attempt%len(keys)]
}

func pickAgentByWeight(agents []*entity.ModelAgent) *entity.ModelAgent {

	total := 0
	weights := make([]int, len(agents))
	for i, agent := range agents {
		w := agent.Weight
		if w <= 0 {
			w = 1
		}
		weights[i] = w
		total += w
	}

	n := int(gtime.TimestampMilli() % int64(total))
	for i, w := range weights {
		if n < w {
			return agents[i]
		}
		n -= w
	}

	return agents[len(agents)-1]
}

func pickKeyByWeight(keys []*entity.Key) *entity.Key {

	total := 0
	weights := make([]int, len(keys))
	for i, key := range keys {
		w := key.Weight
		if w <= 0 {
			w = 1
		}
		weights[i] = w
		total += w
	}

	n := int(gtime.TimestampMilli() % int64(total))
	for i, w := range weights {
		if n < w {
			return keys[i]
		}
		n -= w
	}

	return keys[len(keys)-1]
}

func slimModelAgent(agent *entity.ModelAgent) *entity.ModelAgent {
	return &entity.ModelAgent{
		ProviderId: agent.ProviderId,
		Name:       agent.Name,
		BaseUrl:    agent.BaseUrl,
		Path:       agent.Path,
		Weight:     agent.Weight,
		Remark:     agent.Remark,
	}
}

func appendUnique(list []string, value string) []string {
	if value == "" || slices.Contains(list, value) {
		return list
	}
	return append(list, value)
}
