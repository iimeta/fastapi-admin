package model_agent

import (
	"context"
	"regexp"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sModelAgent struct{}

func init() {
	service.RegisterModelAgent(New())
}

func New() service.IModelAgent {
	return &sModelAgent{}
}

// 新建模型代理
func (s *sModelAgent) Create(ctx context.Context, params model.ModelAgentCreateReq) (string, error) {

	if s.IsNameExist(ctx, params.Name) {
		return "", errors.Newf("模型代理名称 \"%s\" 已存在", params.Name)
	}

	id, err := dao.ModelAgent.Insert(ctx, &do.ModelAgent{
		ProviderId:           params.ProviderId,
		Name:                 gstr.Trim(params.Name),
		BaseUrl:              gstr.Trim(params.BaseUrl),
		Path:                 gstr.Trim(params.Path),
		Weight:               params.Weight,
		Models:               params.Models,
		IsEnableModelReplace: params.IsEnableModelReplace,
		ReplaceModels:        params.ReplaceModels,
		TargetModels:         params.TargetModels,
		IsNeverDisable:       params.IsNeverDisable,
		LbStrategy:           params.LbStrategy,
		Remark:               params.Remark,
		Status:               params.Status,
	})

	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if params.Key != "" {
		if err = service.Key().Create(ctx, model.KeyCreateReq{
			ProviderId:     params.ProviderId,
			Key:            params.Key,
			Weight:         params.Weight,
			Models:         params.Models,
			ModelAgents:    []string{id},
			IsAgentsOnly:   params.IsAgentsOnly,
			IsNeverDisable: params.IsNeverDisableKey,
			Remark:         params.Remark,
			Status:         params.Status,
		}, true); err != nil {
			logger.Error(ctx, err)
			return "", err
		}
	}

	modelAgent, err := s.Detail(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return id, nil
}

// 更新模型代理
func (s *sModelAgent) Update(ctx context.Context, params model.ModelAgentUpdateReq) error {

	if s.IsNameExist(ctx, params.Name, params.Id) {
		return errors.Newf("模型代理名称 \"%s\" 已存在", params.Name)
	}

	oldData, err := s.Detail(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.ModelAgent.UpdateById(ctx, params.Id, &do.ModelAgent{
		ProviderId:           params.ProviderId,
		Name:                 gstr.Trim(params.Name),
		BaseUrl:              gstr.Trim(params.BaseUrl),
		Path:                 gstr.Trim(params.Path),
		Weight:               params.Weight,
		Models:               params.Models,
		IsEnableModelReplace: params.IsEnableModelReplace,
		ReplaceModels:        params.ReplaceModels,
		TargetModels:         params.TargetModels,
		IsNeverDisable:       params.IsNeverDisable,
		LbStrategy:           params.LbStrategy,
		Remark:               params.Remark,
		Status:               params.Status,
		IsAutoDisabled:       oldData.IsAutoDisabled,
		AutoDisabledReason:   oldData.AutoDisabledReason,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	var oldKeyList []*entity.Key

	if oldData.Key != "" {

		oldKeys := gstr.Split(gstr.Trim(oldData.Key), "\n")

		oldKeyList, err = service.Key().DetailListByKey(ctx, oldKeys)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if err = dao.Key.UpdateMany(ctx, bson.M{"model_agents": bson.M{"$in": []string{params.Id}}}, bson.M{
		"$pull": bson.M{
			"model_agents": params.Id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.Key != "" {
		if err = service.Key().Create(ctx, model.KeyCreateReq{
			ProviderId:     params.ProviderId,
			Key:            params.Key,
			Weight:         params.Weight,
			Models:         params.Models,
			ModelAgents:    []string{params.Id},
			IsAgentsOnly:   params.IsAgentsOnly,
			IsNeverDisable: params.IsNeverDisableKey,
			Remark:         params.Remark,
			Status:         params.Status,
		}, true); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if oldKeyList != nil && len(oldKeyList) > 0 {

		oldKeyMap := util.ToMap(oldKeyList, func(t *entity.Key) string {
			return t.Key
		})

		newKeys := gstr.Split(gstr.Trim(params.Key), "\n")
		newKeyMap := util.ToMap(newKeys, func(t string) string {
			return t
		})

		updateKeys := make([]string, 0)
		for _, key := range oldKeyList {
			if newKeyMap[key.Key] == "" {
				updateKeys = append(updateKeys, key.Key)
			}
		}

		if len(updateKeys) > 0 {

			keys, err := service.Key().DetailListByKey(ctx, updateKeys)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			for _, key := range keys {
				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_KEY, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: oldKeyMap[key.Key],
					NewData: key,
				}); err != nil {
					logger.Error(ctx, err)
					return err
				}
			}
		}
	}

	modelAgent, err := s.Detail(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改模型代理状态
func (s *sModelAgent) ChangeStatus(ctx context.Context, params model.ModelAgentChangeStatusReq) error {

	if err := dao.ModelAgent.UpdateById(ctx, params.Id, bson.M{
		"status":               params.Status,
		"is_auto_disabled":     false,
		"auto_disabled_reason": "",
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	modelAgent, err := s.Detail(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err := redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除模型代理
func (s *sModelAgent) Delete(ctx context.Context, id string) error {

	modelAgent, err := s.Detail(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.ModelAgent.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.Key.UpdateMany(ctx, bson.M{"model_agents": bson.M{"$in": []string{id}}}, bson.M{
		"$pull": bson.M{
			"model_agents": id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 模型代理详情
func (s *sModelAgent) Detail(ctx context.Context, id string) (*model.ModelAgent, error) {

	modelAgent, err := dao.ModelAgent.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	providerName := modelAgent.ProviderId
	if provider, err := dao.Provider.FindById(ctx, modelAgent.ProviderId); err == nil && provider != nil {
		providerName = provider.Name
	}

	modelList, err := dao.Model.Find(ctx, bson.M{"_id": bson.M{"$in": modelAgent.Models}}, &dao.FindOptions{SortFields: []string{"-updated_at", "name"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelNames := make([]string, 0)
	for _, model := range modelList {
		modelNames = append(modelNames, model.Name)
	}

	fallbackModelList, err := dao.Model.Find(ctx, bson.M{"fallback_config.model_agent": id}, &dao.FindOptions{SortFields: []string{"-updated_at", "name"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	fallbackModels := make([]string, 0)
	fallbackModelNames := make([]string, 0)

	for _, model := range fallbackModelList {
		fallbackModels = append(fallbackModels, model.Id)
		fallbackModelNames = append(fallbackModelNames, model.Name)
	}

	keyList, err := dao.Key.Find(ctx, bson.M{"model_agents": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	keys := make([]string, 0)
	for _, key := range keyList {
		keys = append(keys, key.Key)
	}

	return &model.ModelAgent{
		Id:                   modelAgent.Id,
		ProviderId:           modelAgent.ProviderId,
		ProviderName:         providerName,
		Name:                 modelAgent.Name,
		BaseUrl:              modelAgent.BaseUrl,
		Path:                 modelAgent.Path,
		Weight:               modelAgent.Weight,
		Models:               modelAgent.Models,
		ModelNames:           modelNames,
		FallbackModels:       fallbackModels,
		FallbackModelNames:   fallbackModelNames,
		IsEnableModelReplace: modelAgent.IsEnableModelReplace,
		ReplaceModels:        modelAgent.ReplaceModels,
		TargetModels:         modelAgent.TargetModels,
		IsNeverDisable:       modelAgent.IsNeverDisable,
		LbStrategy:           modelAgent.LbStrategy,
		Key:                  gstr.Join(keys, "\n"),
		Remark:               modelAgent.Remark,
		Status:               modelAgent.Status,
		IsAutoDisabled:       modelAgent.IsAutoDisabled,
		AutoDisabledReason:   modelAgent.AutoDisabledReason,
		Creator:              modelAgent.Creator,
		Updater:              modelAgent.Updater,
		CreatedAt:            util.FormatDateTime(modelAgent.CreatedAt),
		UpdatedAt:            util.FormatDateTime(modelAgent.UpdatedAt),
	}, nil
}

// 模型代理分页列表
func (s *sModelAgent) Page(ctx context.Context, params model.ModelAgentPageReq) (*model.ModelAgentPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.ProviderId != "" {
		filter["provider_id"] = params.ProviderId
	}

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Name),
		}
	}

	if params.BaseUrl != "" {
		filter["base_url"] = bson.M{
			"$regex": regexp.QuoteMeta(params.BaseUrl),
		}
	}

	if len(params.Models) > 0 {

		modelAgents, err := dao.ModelAgent.Find(ctx, bson.M{"models": bson.M{"$in": params.Models}})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(modelAgents) == 0 {
			return nil, nil
		}

		modelAgentIds := make([]string, 0)
		for _, modelAgent := range modelAgents {
			modelAgentIds = append(modelAgentIds, modelAgent.Id)
		}

		filter["_id"] = bson.M{
			"$in": modelAgentIds,
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if params.Remark != "" {
		filter["remark"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Remark),
		}
	}

	results, err := dao.ModelAgent.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	providers, err := dao.Provider.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	providerMap := util.ToMap(providers, func(t *entity.Provider) string {
		return t.Id
	})

	modelList, err := service.Model().List(ctx, model.ModelListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelMap := make(map[string][]string)
	modelNameMap := make(map[string][]string)

	fallbackModelMap := make(map[string][]string)
	fallbackModelNameMap := make(map[string][]string)

	for _, model := range modelList {

		for _, id := range model.ModelAgents {
			modelMap[id] = append(modelMap[id], model.Id)
			modelNameMap[id] = append(modelNameMap[id], model.Name)
		}

		if model.IsEnableFallback && model.FallbackConfig.ModelAgent != "" {
			fallbackModelMap[model.FallbackConfig.ModelAgent] = append(fallbackModelMap[model.FallbackConfig.ModelAgent], model.Id)
			fallbackModelNameMap[model.FallbackConfig.ModelAgent] = append(fallbackModelNameMap[model.FallbackConfig.ModelAgent], model.Name)
		}
	}

	items := make([]*model.ModelAgent, 0)
	for _, result := range results {

		providerName := result.ProviderId
		if providerMap[result.ProviderId] != nil {
			providerName = providerMap[result.ProviderId].Name
		}

		items = append(items, &model.ModelAgent{
			Id:                 result.Id,
			ProviderId:         result.ProviderId,
			ProviderName:       providerName,
			Name:               result.Name,
			BaseUrl:            result.BaseUrl,
			Path:               result.Path,
			Weight:             result.Weight,
			LbStrategy:         result.LbStrategy,
			Models:             modelMap[result.Id],
			ModelNames:         modelNameMap[result.Id],
			FallbackModels:     fallbackModelMap[result.Id],
			FallbackModelNames: fallbackModelNameMap[result.Id],
			Remark:             result.Remark,
			Status:             result.Status,
			CreatedAt:          util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:          util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.ModelAgentPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 模型代理列表
func (s *sModelAgent) List(ctx context.Context, params model.ModelAgentListReq) ([]*model.ModelAgent, error) {

	filter := bson.M{}

	results, err := dao.ModelAgent.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelList, err := service.Model().List(ctx, model.ModelListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelMap := make(map[string][]string)
	modelNameMap := make(map[string][]string)

	for _, model := range modelList {
		for _, id := range model.ModelAgents {
			modelMap[id] = append(modelMap[id], model.Id)
			modelNameMap[id] = append(modelNameMap[id], model.Name)
		}
	}

	items := make([]*model.ModelAgent, 0)
	for _, result := range results {
		items = append(items, &model.ModelAgent{
			Id:         result.Id,
			ProviderId: result.ProviderId,
			Name:       result.Name,
			BaseUrl:    result.BaseUrl,
			Path:       result.Path,
			Models:     modelMap[result.Id],
			ModelNames: modelNameMap[result.Id],
			Status:     result.Status,
		})
	}

	return items, nil
}

// 模型代理批量操作
func (s *sModelAgent) BatchOperate(ctx context.Context, params model.ModelAgentBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.ModelAgentChangeStatusReq{
				Id:     id,
				Status: gconv.Int(params.Value),
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	case consts.ACTION_DELETE:
		for _, id := range params.Ids {
			if err := s.Delete(ctx, id); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 模型代理名称是否存在
func (s *sModelAgent) IsNameExist(ctx context.Context, name string, id ...string) bool {

	modelAgent, err := dao.ModelAgent.FindOne(ctx, bson.M{"name": gstr.Trim(name)})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false
		}
		logger.Error(ctx, err)
		return true
	}

	if modelAgent != nil {
		if len(id) > 0 && modelAgent.Id == id[0] {
			return false
		}
		return true
	}

	return false
}
