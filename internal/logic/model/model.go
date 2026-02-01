package model

import (
	"context"
	"fmt"
	"regexp"
	"slices"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	mcommon "github.com/iimeta/fastapi-admin/v2/internal/model/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	sutil "github.com/iimeta/fastapi-sdk/v2/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sModel struct{}

func init() {
	service.RegisterModel(New())
}

func New() service.IModel {
	return &sModel{}
}

// 新建模型
func (s *sModel) Create(ctx context.Context, params model.ModelCreateReq) error {

	if s.IsNameExist(ctx, params.Name) {
		return errors.Newf("模型名称 \"%s\" 已存在", params.Name)
	}

	m := &do.Model{
		ProviderId:           params.ProviderId,
		Name:                 gstr.Trim(params.Name),
		Model:                gstr.Trim(params.Model),
		Type:                 params.Type,
		BaseUrl:              gstr.Trim(params.BaseUrl),
		Path:                 gstr.Trim(params.Path),
		IsEnablePresetConfig: params.IsEnablePresetConfig,
		PresetConfig:         params.PresetConfig,
		Pricing:              common.ConvModelPricingToRatio(params.Pricing),
		RequestDataFormat:    params.RequestDataFormat,
		ResponseDataFormat:   params.ResponseDataFormat,
		IsPublic:             params.IsPublic,
		IsEnableModelAgent:   params.IsEnableModelAgent,
		LbStrategy:           params.LbStrategy,
		IsEnableForward:      params.IsEnableForward,
		ForwardConfig:        params.ForwardConfig,
		IsEnableFallback:     params.IsEnableFallback,
		FallbackConfig:       params.FallbackConfig,
		Remark:               params.Remark,
		Status:               params.Status,
	}

	id, err := dao.Model.Insert(ctx, m)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.Model.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.IsPublic && len(params.Groups) == 0 {

		params.Groups, err = service.Group().PublicGroups(ctx)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if len(params.Groups) == 0 {
			if id, err := service.Group().Create(ctx, model.GroupCreateReq{
				Name:     "系统自动创建分组",
				Discount: 1,
				Models:   []string{},
				IsPublic: true,
				Remark:   "此分组为系统自动创建, 可对其修改和删除",
				Status:   1,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			} else {
				params.Groups = append(params.Groups, id)
			}
		}
	}

	if len(params.Groups) > 0 {
		for _, group := range params.Groups {

			oldData, err := dao.Group.FindById(ctx, group)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			newData, err := dao.Group.FindOneAndUpdateById(ctx, group, bson.M{
				"models": append(oldData.Models, id),
			})
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: oldData,
				NewData: newData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	for _, modelAgentId := range params.ModelAgents {

		oldData, err := dao.ModelAgent.FindById(ctx, modelAgentId)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.ModelAgent.UpdateById(ctx, oldData.Id, &do.ModelAgent{
			ProviderId:           oldData.ProviderId,
			Name:                 oldData.Name,
			BaseUrl:              oldData.BaseUrl,
			Path:                 oldData.Path,
			Weight:               oldData.Weight,
			Models:               append(oldData.Models, id),
			IsEnableModelReplace: oldData.IsEnableModelReplace,
			ReplaceModels:        oldData.ReplaceModels,
			TargetModels:         oldData.TargetModels,
			IsNeverDisable:       oldData.IsNeverDisable,
			LbStrategy:           oldData.LbStrategy,
			Remark:               oldData.Remark,
			Status:               oldData.Status,
			IsAutoDisabled:       oldData.IsAutoDisabled,
			AutoDisabledReason:   oldData.AutoDisabledReason,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		modelAgent, err := service.ModelAgent().Detail(ctx, modelAgentId)
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
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_MODEL, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新模型
func (s *sModel) Update(ctx context.Context, params model.ModelUpdateReq) error {

	if s.IsNameExist(ctx, params.Name, params.Id) {
		return errors.Newf("模型名称 \"%s\" 已存在", params.Name)
	}

	oldData, err := dao.Model.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	m := &do.Model{
		ProviderId:           params.ProviderId,
		Name:                 gstr.Trim(params.Name),
		Model:                gstr.Trim(params.Model),
		Type:                 params.Type,
		BaseUrl:              gstr.Trim(params.BaseUrl),
		Path:                 gstr.Trim(params.Path),
		IsEnablePresetConfig: params.IsEnablePresetConfig,
		PresetConfig:         params.PresetConfig,
		Pricing:              common.ConvModelPricingToRatio(params.Pricing),
		RequestDataFormat:    params.RequestDataFormat,
		ResponseDataFormat:   params.ResponseDataFormat,
		IsPublic:             params.IsPublic,
		IsEnableModelAgent:   params.IsEnableModelAgent,
		LbStrategy:           params.LbStrategy,
		IsEnableForward:      params.IsEnableForward,
		ForwardConfig:        params.ForwardConfig,
		IsEnableFallback:     params.IsEnableFallback,
		FallbackConfig:       params.FallbackConfig,
		Remark:               params.Remark,
		Status:               params.Status,
	}

	newData, err := dao.Model.FindOneAndUpdateById(ctx, params.Id, m)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	groups, err := service.Group().GetGroupsByModels(ctx, oldData.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	oldGroups := gset.NewStrSetFrom(groups)
	newGroups := gset.NewStrSetFrom(params.Groups)

	if !oldGroups.Equal(newGroups) {

		allGroups := gset.NewStrSet()
		allGroups.Add(oldGroups.Slice()...)
		allGroups.Add(newGroups.Slice()...)

		for _, group := range allGroups.Slice() {

			// 新的有, 旧的没有, 说明新增了
			if newGroups.Contains(group) && !oldGroups.Contains(group) {

				oldData, err := dao.Group.FindById(ctx, group)
				if err != nil {
					logger.Error(ctx, err)
					return err
				}

				newData, err := dao.Group.FindOneAndUpdateById(ctx, group, bson.M{
					"models": append(oldData.Models, params.Id),
				})
				if err != nil {
					logger.Error(ctx, err)
					return err
				}

				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: oldData,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
					return err
				}

			} else if oldGroups.Contains(group) && !newGroups.Contains(group) { // 旧的有, 新的没有, 说明移除了

				oldData, err := dao.Group.FindById(ctx, group)
				if err != nil {
					logger.Error(ctx, err)
					return err
				}

				oldGroups := gset.NewStrSetFrom(oldData.Models)
				oldGroups.Remove(params.Id)

				newData, err := dao.Group.FindOneAndUpdateById(ctx, group, bson.M{
					"models": oldGroups.Slice(),
				})
				if err != nil {
					logger.Error(ctx, err)
					return err
				}

				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: oldData,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
					return err
				}
			}
		}
	}

	modelAgents, err := dao.ModelAgent.Find(ctx, bson.M{"models": bson.M{"$in": []string{params.Id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	oldModelAgents := make([]string, 0)
	for _, oldData := range modelAgents {

		oldModelAgents = append(oldModelAgents, oldData.Id)

		if !slices.Contains(params.ModelAgents, oldData.Id) {

			modelAgent := &do.ModelAgent{
				ProviderId:           oldData.ProviderId,
				Name:                 oldData.Name,
				BaseUrl:              oldData.BaseUrl,
				Path:                 oldData.Path,
				Weight:               oldData.Weight,
				Models:               oldData.Models,
				IsEnableModelReplace: oldData.IsEnableModelReplace,
				ReplaceModels:        oldData.ReplaceModels,
				TargetModels:         oldData.TargetModels,
				IsNeverDisable:       oldData.IsNeverDisable,
				LbStrategy:           oldData.LbStrategy,
				Remark:               oldData.Remark,
				Status:               oldData.Status,
				IsAutoDisabled:       oldData.IsAutoDisabled,
				AutoDisabledReason:   oldData.AutoDisabledReason,
			}

			for i, modelId := range modelAgent.Models {
				if modelId == params.Id {
					modelAgent.Models = util.Delete(modelAgent.Models, i)
					break
				}
			}

			if err = dao.ModelAgent.UpdateById(ctx, oldData.Id, modelAgent); err != nil {
				logger.Error(ctx, err)
				return err
			}

			newData, err := service.ModelAgent().Detail(ctx, oldData.Id)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: oldData,
				NewData: newData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	for _, modelAgentId := range params.ModelAgents {

		if !slices.Contains(oldModelAgents, modelAgentId) {

			oldData, err := dao.ModelAgent.FindById(ctx, modelAgentId)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			if err = dao.ModelAgent.UpdateById(ctx, oldData.Id, &do.ModelAgent{
				ProviderId:           oldData.ProviderId,
				Name:                 oldData.Name,
				BaseUrl:              oldData.BaseUrl,
				Path:                 oldData.Path,
				Weight:               oldData.Weight,
				Models:               append(oldData.Models, params.Id),
				IsEnableModelReplace: oldData.IsEnableModelReplace,
				ReplaceModels:        oldData.ReplaceModels,
				TargetModels:         oldData.TargetModels,
				IsNeverDisable:       oldData.IsNeverDisable,
				LbStrategy:           oldData.LbStrategy,
				Remark:               oldData.Remark,
				Status:               oldData.Status,
				IsAutoDisabled:       oldData.IsAutoDisabled,
				AutoDisabledReason:   oldData.AutoDisabledReason,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}

			modelAgent, err := service.ModelAgent().Detail(ctx, modelAgentId)
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
		}
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_MODEL, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改模型状态
func (s *sModel) ChangeStatus(ctx context.Context, params model.ModelChangeStatusReq) error {

	newData, err := dao.Model.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_MODEL, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除模型
func (s *sModel) Delete(ctx context.Context, id string) error {

	oldData, err := dao.Model.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	apps, err := dao.App.Find(ctx, bson.M{"models": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, app := range apps {

		appModelsReq := model.AppModelsReq{
			AppId:  app.AppId,
			Models: []string{},
		}

		for _, m := range app.Models {
			if m != id {
				appModelsReq.Models = append(appModelsReq.Models, m)
			}
		}

		if err = service.App().Models(ctx, appModelsReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	appKeys, err := dao.AppKey.Find(ctx, bson.M{"models": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, key := range appKeys {

		keyModelsReq := model.AppKeyModelsReq{
			Id:     key.Id,
			Models: []string{},
		}

		for _, m := range key.Models {
			if m != id {
				keyModelsReq.Models = append(keyModelsReq.Models, m)
			}
		}

		if err = service.AppKey().Models(ctx, keyModelsReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	keys, err := dao.Key.Find(ctx, bson.M{"models": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, key := range keys {

		keyModelsReq := model.KeyModelsReq{
			Id:     key.Id,
			Models: []string{},
		}

		for _, m := range key.Models {
			if m != id {
				keyModelsReq.Models = append(keyModelsReq.Models, m)
			}
		}

		if err = service.Key().Models(ctx, keyModelsReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_MODEL, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: oldData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	groups, err := service.Group().GetGroupsByModels(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, group := range groups {

		oldData, err := dao.Group.FindById(ctx, group)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		oldGroups := gset.NewStrSetFrom(oldData.Models)
		oldGroups.Remove(id)

		newData, err := dao.Group.FindOneAndUpdateById(ctx, group, bson.M{
			"models": oldGroups.Slice(),
		})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			OldData: oldData,
			NewData: newData,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

// 模型详情
func (s *sModel) Detail(ctx context.Context, id string) (*model.Model, error) {

	m, err := dao.Model.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelAgentList, err := dao.ModelAgent.Find(ctx, bson.M{"models": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelAgents := make([]string, 0)
	modelAgentNames := make([]string, 0)
	for _, modelAgent := range modelAgentList {
		modelAgents = append(modelAgents, modelAgent.Id)
		modelAgentNames = append(modelAgentNames, modelAgent.Name)
	}

	providerName := m.ProviderId
	providerCode := m.ProviderId
	if provider, err := dao.Provider.FindById(ctx, m.ProviderId); err == nil && provider != nil {
		providerName = provider.Name
		providerCode = provider.Code
	}

	groups, err := service.Group().List(ctx, model.GroupListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groupIds := make([]string, 0)
	groupNames := make([]string, 0)
	for _, group := range groups {
		if slices.Contains(group.Models, m.Id) {
			groupIds = append(groupIds, group.Id)
			groupNames = append(groupNames, group.Name)
		}
	}

	detail := &model.Model{
		Id:                   m.Id,
		ProviderId:           m.ProviderId,
		ProviderName:         providerName,
		ProviderCode:         providerCode,
		Name:                 m.Name,
		Model:                m.Model,
		Type:                 m.Type,
		BaseUrl:              m.BaseUrl,
		Path:                 m.Path,
		Groups:               groupIds,
		GroupNames:           groupNames,
		IsEnablePresetConfig: m.IsEnablePresetConfig,
		PresetConfig:         m.PresetConfig,
		Pricing:              common.ConvModelPricingToPrice(m.Pricing),
		RequestDataFormat:    m.RequestDataFormat,
		ResponseDataFormat:   m.ResponseDataFormat,
		IsPublic:             m.IsPublic,
		IsEnableModelAgent:   m.IsEnableModelAgent,
		LbStrategy:           m.LbStrategy,
		ModelAgents:          modelAgents,
		ModelAgentNames:      modelAgentNames,
		IsEnableForward:      m.IsEnableForward,
		ForwardConfig:        m.ForwardConfig,
		IsEnableFallback:     m.IsEnableFallback,
		FallbackConfig:       m.FallbackConfig,
		Remark:               m.Remark,
		Status:               m.Status,
		CreatedAt:            util.FormatDateTime(m.CreatedAt),
		UpdatedAt:            util.FormatDateTime(m.UpdatedAt),
	}

	if detail.ForwardConfig != nil {

		if detail.ForwardConfig.TargetModel != "" {
			if modelNames, err := s.ModelNames(ctx, []string{detail.ForwardConfig.TargetModel}); err != nil {
				logger.Error(ctx, err)
			} else if len(modelNames) > 0 {
				detail.ForwardConfig.TargetModelName = modelNames[0]
			}
		}

		if detail.ForwardConfig.DecisionModel != "" {
			if modelNames, err := s.ModelNames(ctx, []string{detail.ForwardConfig.DecisionModel}); err != nil {
				logger.Error(ctx, err)
			} else if len(modelNames) > 0 {
				detail.ForwardConfig.DecisionModelName = modelNames[0]
			}
		}

		if detail.ForwardConfig.TargetModels != nil && len(detail.ForwardConfig.TargetModels) > 0 {
			if modelNames, err := s.ModelNames(ctx, detail.ForwardConfig.TargetModels); err != nil {
				logger.Error(ctx, err)
			} else {
				detail.ForwardConfig.TargetModelNames = modelNames
			}
		}
	}

	if detail.FallbackConfig != nil {

		if detail.FallbackConfig.ModelAgent != "" {
			if modelAgent, _ := dao.ModelAgent.FindById(ctx, detail.FallbackConfig.ModelAgent); modelAgent != nil {
				detail.FallbackConfig.ModelAgentName = modelAgent.Name
			}
		}

		if detail.FallbackConfig.Model != "" {
			if model, _ := dao.Model.FindById(ctx, detail.FallbackConfig.Model); model != nil {
				detail.FallbackConfig.ModelName = model.Name
			}
		}
	}

	return detail, nil
}

// 模型分页列表
func (s *sModel) Page(ctx context.Context, params model.ModelPageReq) (*model.ModelPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {

		reseller, err := service.Reseller().GetResellerByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(reseller.Groups) == 0 {
			return nil, nil
		}

		if params.Group != "" {

			if !slices.Contains(reseller.Groups, params.Group) {
				return nil, nil
			}

			models, err := service.Group().GetModelsByGroups(ctx, params.Group)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}

			if len(models) == 0 {
				return nil, nil
			}

			filter["_id"] = bson.M{
				"$in": models,
			}

		} else {

			models, err := service.Group().GetModelsByGroups(ctx, reseller.Groups...)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}

			if len(models) == 0 {
				return nil, nil
			}

			filter["_id"] = bson.M{
				"$in": models,
			}
		}
	}

	if service.Session().IsUserRole(ctx) {

		user, err := service.User().GetUserByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(user.Groups) == 0 {
			return nil, nil
		}

		if params.Group != "" {

			if !slices.Contains(user.Groups, params.Group) {
				return nil, nil
			}

			models, err := service.Group().GetModelsByGroups(ctx, params.Group)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}

			if len(models) == 0 {
				return nil, nil
			}

			filter["_id"] = bson.M{
				"$in": models,
			}

		} else {

			models, err := service.Group().GetModelsByGroups(ctx, user.Groups...)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}

			if len(models) == 0 {
				return nil, nil
			}

			filter["_id"] = bson.M{
				"$in": models,
			}
		}
	}

	if params.Group != "" && service.Session().IsAdminRole(ctx) {

		models, err := service.Group().GetModelsByGroups(ctx, params.Group)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(models) == 0 {
			return nil, nil
		}

		filter["_id"] = bson.M{
			"$in": models,
		}
	}

	if params.ProviderId != "" {
		filter["provider_id"] = params.ProviderId
	}

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Name),
		}
	}

	if params.Model != "" {
		filter["model"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Model),
		}
	}

	if params.Type != 0 {
		filter["type"] = params.Type
	}

	if params.Remark != "" {
		filter["remark"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Remark),
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	results, err := dao.Model.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at", "name"}})
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

	groups, err := service.Group().List(ctx, model.GroupListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Model, 0)
	for _, result := range results {

		providerName := result.ProviderId
		providerCode := result.ProviderId
		if providerMap[result.ProviderId] != nil {
			providerName = providerMap[result.ProviderId].Name
			providerCode = providerMap[result.ProviderId].Code
		}

		groupNames := make([]string, 0)
		for _, group := range groups {
			if slices.Contains(group.Models, result.Id) {
				groupNames = append(groupNames, group.Name)
			}
		}

		model := &model.Model{
			Id:                 result.Id,
			ProviderId:         result.ProviderId,
			ProviderName:       providerName,
			ProviderCode:       providerCode,
			Name:               result.Name,
			Model:              result.Model,
			Type:               result.Type,
			GroupNames:         groupNames,
			Pricing:            common.ConvModelPricingToPrice(result.Pricing),
			RequestDataFormat:  result.RequestDataFormat,
			ResponseDataFormat: result.ResponseDataFormat,
			IsPublic:           result.IsPublic,
			Remark:             result.Remark,
			Status:             result.Status,
			CreatedAt:          util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:          util.FormatDateTimeMonth(result.UpdatedAt),
		}

		if result.IsEnableModelAgent && service.Session().IsAdminRole(ctx) {
			model.IsEnableModelAgent = result.IsEnableModelAgent
			model.LbStrategy = result.LbStrategy
		}

		items = append(items, model)
	}

	return &model.ModelPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 模型列表
func (s *sModel) List(ctx context.Context, params model.ModelListReq) ([]*model.Model, error) {

	filter := bson.M{}

	if params.ProviderId != "" {
		filter["provider_id"] = params.ProviderId
	}

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Name),
		}
	}

	if params.Model != "" {
		filter["model"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Model),
		}
	}

	if params.Type != 0 {
		filter["type"] = params.Type
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if len(params.Models) > 0 {
		filter["_id"] = bson.M{
			"$in": params.Models,
		}
	}

	if service.Session().IsResellerRole(ctx) {

		reseller, err := service.Reseller().GetResellerByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		modelSet := gset.NewStrSet()

		if len(reseller.Groups) > 0 {
			models, err := service.Group().GetModelsByGroups(ctx, reseller.Groups...)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			modelSet.Add(models...)
		}

		if len(modelSet.Slice()) == 0 {
			return nil, nil
		}

		if len(params.Models) > 0 {
			filter["_id"] = bson.M{
				"$in": gset.NewStrSetFrom(params.Models).Intersect(modelSet).Slice(),
			}
		} else {
			filter["_id"] = bson.M{
				"$in": modelSet.Slice(),
			}
		}
	}

	if service.Session().IsUserRole(ctx) {

		user, err := service.User().GetUserByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		modelSet := gset.NewStrSet()

		if len(user.Groups) > 0 {
			models, err := service.Group().GetModelsByGroups(ctx, user.Groups...)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			modelSet.Add(models...)
		}

		if len(modelSet.Slice()) == 0 {
			return nil, nil
		}

		if len(params.Models) > 0 {
			filter["_id"] = bson.M{
				"$in": gset.NewStrSetFrom(params.Models).Intersect(modelSet).Slice(),
			}
		} else {
			filter["_id"] = bson.M{
				"$in": modelSet.Slice(),
			}
		}
	}

	results, err := dao.Model.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"-updated_at", "name"}})
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

	items := make([]*model.Model, 0)
	for _, result := range results {

		providerName := result.ProviderId
		if providerMap[result.ProviderId] != nil {
			providerName = providerMap[result.ProviderId].Name
		}

		model := &model.Model{
			Id:           result.Id,
			ProviderId:   result.ProviderId,
			ProviderName: providerName,
			Name:         result.Name,
			Model:        result.Model,
			Type:         result.Type,
			Status:       result.Status,
		}

		if service.Session().IsAdminRole(ctx) {
			model.IsEnableFallback = result.IsEnableFallback
			model.FallbackConfig = result.FallbackConfig
		}

		items = append(items, model)
	}

	return items, nil
}

// 模型批量操作
func (s *sModel) BatchOperate(ctx context.Context, params model.ModelBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_AGENT:

		results, err := dao.Model.Find(ctx, bson.M{
			"_id": bson.M{
				"$in": params.Ids,
			},
		})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		groups, err := service.Group().List(ctx, model.GroupListReq{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, result := range results {

			groupIds := make([]string, 0)
			for _, group := range groups {
				if slices.Contains(group.Models, result.Id) {
					groupIds = append(groupIds, group.Id)
				}
			}

			m := model.ModelUpdateReq{
				Id:                   result.Id,
				ProviderId:           result.ProviderId,
				Name:                 result.Name,
				Model:                result.Model,
				Type:                 result.Type,
				BaseUrl:              result.BaseUrl,
				Path:                 result.Path,
				IsEnablePresetConfig: result.IsEnablePresetConfig,
				PresetConfig:         result.PresetConfig,
				Pricing:              common.ConvModelPricingToPrice(result.Pricing),
				RequestDataFormat:    result.RequestDataFormat,
				ResponseDataFormat:   result.ResponseDataFormat,
				IsPublic:             result.IsPublic,
				Groups:               groupIds,
				LbStrategy:           result.LbStrategy,
				IsEnableForward:      result.IsEnableForward,
				ForwardConfig:        result.ForwardConfig,
				IsEnableFallback:     result.IsEnableFallback,
				FallbackConfig:       result.FallbackConfig,
				Remark:               result.Remark,
				Status:               result.Status,
			}

			if params.Value == "all" {
				m.IsEnableModelAgent = true
				m.LbStrategy = params.LbStrategy
				m.ModelAgents = params.ModelAgents
			} else {
				m.IsEnableModelAgent = gconv.Bool(params.Value)
				if m.IsEnableModelAgent {

					modelAgents, err := dao.ModelAgent.Find(ctx, bson.M{"models": bson.M{"$in": []string{result.Id}}})
					if err != nil {
						logger.Error(ctx, err)
						continue
					}

					if len(modelAgents) == 0 {
						continue
					}

					for _, modelAgent := range modelAgents {
						m.ModelAgents = append(m.ModelAgents, modelAgent.Id)
					}
				}
			}

			if err = s.Update(ctx, m); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

	case consts.ACTION_FORWARD:

		results, err := dao.Model.Find(ctx, bson.M{
			"_id": bson.M{
				"$in": params.Ids,
			},
		})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		groups, err := service.Group().List(ctx, model.GroupListReq{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, result := range results {

			groupIds := make([]string, 0)
			for _, group := range groups {
				if slices.Contains(group.Models, result.Id) {
					groupIds = append(groupIds, group.Id)
				}
			}

			m := model.ModelUpdateReq{
				Id:                   result.Id,
				ProviderId:           result.ProviderId,
				Name:                 result.Name,
				Model:                result.Model,
				Type:                 result.Type,
				BaseUrl:              result.BaseUrl,
				Path:                 result.Path,
				IsEnablePresetConfig: result.IsEnablePresetConfig,
				PresetConfig:         result.PresetConfig,
				Pricing:              common.ConvModelPricingToPrice(result.Pricing),
				RequestDataFormat:    result.RequestDataFormat,
				ResponseDataFormat:   result.ResponseDataFormat,
				IsPublic:             result.IsPublic,
				Groups:               groupIds,
				IsEnableModelAgent:   result.IsEnableModelAgent,
				LbStrategy:           result.LbStrategy,
				ForwardConfig:        result.ForwardConfig,
				IsEnableFallback:     result.IsEnableFallback,
				FallbackConfig:       result.FallbackConfig,
				Remark:               result.Remark,
				Status:               result.Status,
			}

			if params.Value == "all" {

				if m.ForwardConfig == nil {
					m.ForwardConfig = new(mcommon.ForwardConfig)
				}

				m.IsEnableForward = true
				m.ForwardConfig.ForwardRule = 1
				m.ForwardConfig.TargetModel = params.TargetModel

			} else {
				m.IsEnableForward = gconv.Bool(params.Value)
				if m.IsEnableForward && (m.ForwardConfig == nil ||
					(m.ForwardConfig.ForwardRule == 1 && m.ForwardConfig.TargetModel == "") ||
					(m.ForwardConfig.ForwardRule == 2 && len(m.ForwardConfig.TargetModels) == 0)) {
					continue
				}
			}

			if modelAgents, err := dao.ModelAgent.Find(ctx, bson.M{"models": bson.M{"$in": []string{result.Id}}}); err != nil {
				logger.Error(ctx, err)
			} else {
				for _, modelAgent := range modelAgents {
					m.ModelAgents = append(m.ModelAgents, modelAgent.Id)
				}
			}

			if err = s.Update(ctx, m); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

	case consts.ACTION_FALLBACK:

		results, err := dao.Model.Find(ctx, bson.M{
			"_id": bson.M{
				"$in": params.Ids,
			},
		})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		groups, err := service.Group().List(ctx, model.GroupListReq{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, result := range results {

			groupIds := make([]string, 0)
			for _, group := range groups {
				if slices.Contains(group.Models, result.Id) {
					groupIds = append(groupIds, group.Id)
				}
			}

			m := model.ModelUpdateReq{
				Id:                   result.Id,
				ProviderId:           result.ProviderId,
				Name:                 result.Name,
				Model:                result.Model,
				Type:                 result.Type,
				BaseUrl:              result.BaseUrl,
				Path:                 result.Path,
				IsEnablePresetConfig: result.IsEnablePresetConfig,
				PresetConfig:         result.PresetConfig,
				Pricing:              common.ConvModelPricingToPrice(result.Pricing),
				RequestDataFormat:    result.RequestDataFormat,
				ResponseDataFormat:   result.ResponseDataFormat,
				IsPublic:             result.IsPublic,
				Groups:               groupIds,
				IsEnableModelAgent:   result.IsEnableModelAgent,
				LbStrategy:           result.LbStrategy,
				IsEnableForward:      result.IsEnableForward,
				ForwardConfig:        result.ForwardConfig,
				FallbackConfig:       result.FallbackConfig,
				Remark:               result.Remark,
				Status:               result.Status,
			}

			if params.Value == "all" {
				m.IsEnableFallback = true
				m.FallbackConfig = params.FallbackConfig
			} else {
				m.IsEnableFallback = gconv.Bool(params.Value)
				if m.IsEnableFallback && m.FallbackConfig == nil {
					continue
				}
			}

			if modelAgents, err := dao.ModelAgent.Find(ctx, bson.M{"models": bson.M{"$in": []string{result.Id}}}); err != nil {
				logger.Error(ctx, err)
			} else {
				for _, modelAgent := range modelAgents {
					m.ModelAgents = append(m.ModelAgents, modelAgent.Id)
				}
			}

			if err = s.Update(ctx, m); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.ModelChangeStatusReq{
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

// 模型树
func (s *sModel) Tree(ctx context.Context, params model.ModelTreeReq) ([]*model.Tree, error) {

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {

		reseller, err := service.Reseller().GetResellerByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		modelSet := gset.NewStrSet()

		if len(reseller.Groups) > 0 {
			models, err := service.Group().GetModelsByGroups(ctx, reseller.Groups...)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			modelSet.Add(models...)
		}

		if len(modelSet.Slice()) == 0 {
			return nil, nil
		}

		filter["_id"] = bson.M{
			"$in": modelSet.Slice(),
		}
	}

	if service.Session().IsUserRole(ctx) {

		user, err := service.User().GetUserByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		modelSet := gset.NewStrSet()

		if len(user.Groups) > 0 {
			models, err := service.Group().GetModelsByGroups(ctx, user.Groups...)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			modelSet.Add(models...)
		}

		if len(modelSet.Slice()) == 0 {
			return nil, nil
		}

		filter["_id"] = bson.M{
			"$in": modelSet.Slice(),
		}
	}

	results, err := dao.Model.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"-updated_at", "name"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	treeData := make(map[string][]model.Tree) // [ProviderId:Type][][Model]

	for _, result := range results {

		providerTree := treeData[fmt.Sprintf("%s:%d", result.ProviderId, result.Type)]
		if providerTree == nil {
			providerTree = make([]model.Tree, 0)
		}

		providerTree = append(providerTree, model.Tree{
			Title: result.Model,
			Value: result.Id,
			Key:   result.Id,
		})

		treeData[fmt.Sprintf("%s:%d", result.ProviderId, result.Type)] = providerTree
	}

	providers, err := service.Provider().List(ctx, model.ProviderListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Tree, 0)
	for _, provider := range providers {

		children := make([]model.Tree, 0)

		for _, typ := range consts.MODEL_TYPES {
			if modelTree := treeData[fmt.Sprintf("%s:%d", provider.Id, typ)]; modelTree != nil {
				children = append(children, model.Tree{
					Title:    consts.MODEL_TYPE[typ],
					Value:    fmt.Sprintf("%s:%d", provider.Id, typ),
					Key:      fmt.Sprintf("%s:%d", provider.Id, typ),
					Children: modelTree,
				})
			}
		}

		if len(children) > 0 {
			items = append(items, &model.Tree{
				Title:    provider.Name,
				Value:    provider.Id,
				Key:      provider.Id,
				Children: children,
			})
		}
	}

	return items, nil
}

// 模型权限列表
func (s *sModel) Permissions(ctx context.Context, params model.ModelPermissionsReq) ([]*model.Model, error) {

	modelListReq := model.ModelListReq{
		ProviderId: params.ProviderId,
		Name:       params.Name,
		Model:      params.Model,
		Type:       params.Type,
		Status:     params.Status,
	}

	switch params.Action {
	case consts.ACTION_APP:

		app, err := service.App().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(app.Models) == 0 {
			return nil, nil
		}

		modelListReq.Models = app.Models

	case consts.ACTION_APP_KEY:

		appKey, err := service.AppKey().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(appKey.Models) == 0 {
			return nil, nil
		}

		modelListReq.Models = appKey.Models

	case consts.ACTION_KEY:

		key, err := service.Key().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(key.Models) == 0 {
			return nil, nil
		}

		modelListReq.Models = key.Models

	case consts.ACTION_AGENT:

		modelAgent, err := service.ModelAgent().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(modelAgent.Models) == 0 {
			return nil, nil
		}

		modelListReq.Models = modelAgent.Models

	case consts.ACTION_FALLBACK:

		modelAgent, err := service.ModelAgent().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(modelAgent.FallbackModels) == 0 {
			return nil, nil
		}

		modelListReq.Models = modelAgent.FallbackModels

	case consts.ACTION_GROUP:

		group, err := service.Group().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(group.Models) == 0 {
			return nil, nil
		}

		modelListReq.Models = group.Models

	default:
		return nil, nil
	}

	return s.List(ctx, modelListReq)
}

// 模型初始化同步
func (s *sModel) InitSync(ctx context.Context, params model.ModelInitSyncReq) error {

	result := &model.ModelsRes{}
	if _, err := sutil.HttpGet(ctx, params.Url, g.MapStrStr{"Authorization": "Bearer " + params.Key}, g.MapStrAny{"is_fastapi": true}, &result, config.Cfg.Http.Timeout, config.Cfg.Http.ProxyUrl, nil); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if result.Data == nil || (len(result.Data) > 0 && result.Data[0].FastAPI == nil) {
		return errors.New("模型接口数据格式不支持, 请联系作者...")
	}

	providers, err := dao.Provider.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	providerNameMap := make(map[string]string)
	providerCodeMap := make(map[string]string)
	for _, provider := range providers {
		providerNameMap[provider.Name] = provider.Id
		providerCodeMap[provider.Code] = provider.Id
	}

	models, err := dao.Model.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	modelMap := make(map[string]*entity.Model)
	for _, model := range models {
		modelMap[model.Name] = model
	}

	modelAgentId := ""
	if params.IsConfigModelAgent {

		if modelAgent, err := dao.ModelAgent.FindOne(ctx, bson.M{"name": "智元 Fast API"}); err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {

				if providerCodeMap["FastAPI"] == "" {
					if providerCodeMap["FastAPI"], err = service.Provider().Create(ctx, model.ProviderCreateReq{
						Name:     "智元 Fast API",
						Code:     "FastAPI",
						IsPublic: true,
						Status:   1,
					}); err != nil {
						logger.Error(ctx, err)
						return err
					}
				}

				if modelAgentId, err = service.ModelAgent().Create(ctx, model.ModelAgentCreateReq{
					ProviderId:   providerCodeMap["FastAPI"],
					Name:         "智元 Fast API",
					BaseUrl:      gstr.Replace(params.Url, "/models", ""),
					Models:       []string{},
					Key:          params.Key,
					IsAgentsOnly: true,
					Status:       1,
				}); err != nil {
					logger.Error(ctx, err)
					return err
				}

			} else {
				logger.Error(ctx, err)
				return err
			}

		} else {
			modelAgentId = modelAgent.Id
		}
	}

	groups, err := service.Group().PublicGroups(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if len(groups) == 0 {
		if _, err = service.Group().Create(ctx, model.GroupCreateReq{
			Name:     "系统自动创建分组",
			Discount: 1,
			Models:   []string{},
			IsPublic: true,
			Remark:   "此分组为系统自动创建, 可对其修改和删除",
			Status:   1,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	for _, data := range result.Data {

		provider := providerNameMap[data.FastAPI.Provider]

		if provider == "" {
			if provider, err = service.Provider().Create(ctx, model.ProviderCreateReq{
				Name:     data.FastAPI.Provider,
				Code:     data.FastAPI.Code,
				IsPublic: true,
				Status:   1,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
			providerNameMap[data.FastAPI.Provider] = provider
			providerCodeMap[data.FastAPI.Code] = provider
		}

		if modelMap[data.Id] == nil {

			modelCreateReq := model.ModelCreateReq{
				ProviderId:         provider,
				Name:               data.Id,
				Model:              data.FastAPI.Model,
				Type:               data.FastAPI.Type,
				Pricing:            data.FastAPI.Pricing,
				RequestDataFormat:  1,
				ResponseDataFormat: 1,
				IsPublic:           true,
				ModelAgents:        []string{},
				Remark:             data.FastAPI.Remark,
				Status:             1,
			}

			if params.IsConfigModelAgent && modelAgentId != "" {
				modelCreateReq.IsEnableModelAgent = true
				modelCreateReq.LbStrategy = 1
				modelCreateReq.ModelAgents = append(modelCreateReq.ModelAgents, modelAgentId)
			}

			if err = s.Create(ctx, modelCreateReq); err != nil {
				logger.Error(ctx, err)
				return err
			}

			modelMap[data.Id] = &entity.Model{Name: data.Id, Model: data.FastAPI.Model}

		} else if params.IsCoverPrice && modelMap[data.Id].Id != "" {

			detail, err := s.Detail(ctx, modelMap[data.Id].Id)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			modelUpdateReq := model.ModelUpdateReq{
				Id:                   detail.Id,
				ProviderId:           detail.ProviderId,
				Name:                 detail.Name,
				Model:                detail.Model,
				Type:                 detail.Type,
				BaseUrl:              detail.BaseUrl,
				Path:                 detail.Path,
				IsEnablePresetConfig: detail.IsEnablePresetConfig,
				PresetConfig:         detail.PresetConfig,
				Pricing:              data.FastAPI.Pricing,
				RequestDataFormat:    detail.RequestDataFormat,
				ResponseDataFormat:   detail.ResponseDataFormat,
				IsPublic:             detail.IsPublic,
				Groups:               detail.Groups,
				IsEnableModelAgent:   detail.IsEnableModelAgent,
				LbStrategy:           detail.LbStrategy,
				ModelAgents:          detail.ModelAgents,
				IsEnableForward:      detail.IsEnableForward,
				ForwardConfig:        detail.ForwardConfig,
				IsEnableFallback:     detail.IsEnableFallback,
				FallbackConfig:       detail.FallbackConfig,
				Remark:               detail.Remark,
				Status:               detail.Status,
			}

			if params.IsConfigModelAgent && modelAgentId != "" && !slices.Contains(modelUpdateReq.ModelAgents, modelAgentId) {
				modelUpdateReq.ModelAgents = append(modelUpdateReq.ModelAgents, modelAgentId)
			}

			if err = s.Update(ctx, modelUpdateReq); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 公开的模型Ids
func (s *sModel) PublicModels(ctx context.Context) ([]string, error) {

	results, err := dao.Model.Find(ctx, bson.M{"is_public": true}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	ids := make([]string, 0)
	for _, result := range results {
		ids = append(ids, result.Id)
	}

	return ids, nil
}

// 根据模型Ids查询模型名称
func (s *sModel) ModelNames(ctx context.Context, models []string) ([]string, error) {

	if models == nil || len(models) == 0 {
		return nil, nil
	}

	results, err := dao.Model.Find(ctx, bson.M{"_id": bson.M{"$in": models}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelNames := make([]string, 0)

	modelMap := util.ToMap(results, func(t *entity.Model) string {
		return t.Id
	})

	for _, id := range models {
		if modelMap[id] != nil {
			modelNames = append(modelNames, modelMap[id].Name)
		}
	}

	return modelNames, nil
}

// 模型名称是否存在
func (s *sModel) IsNameExist(ctx context.Context, name string, id ...string) bool {

	model, err := dao.Model.FindOne(ctx, bson.M{"name": gstr.Trim(name)})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false
		}
		logger.Error(ctx, err)
		return true
	}

	if model != nil {
		if len(id) > 0 && model.Id == id[0] {
			return false
		}
		return true
	}

	return false
}
