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
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/common"
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
		Corp:                 params.Corp,
		Name:                 gstr.Trim(params.Name),
		Model:                gstr.Trim(params.Model),
		Type:                 params.Type,
		BaseUrl:              gstr.Trim(params.BaseUrl),
		Path:                 gstr.Trim(params.Path),
		IsEnablePresetConfig: params.IsEnablePresetConfig,
		PresetConfig:         params.PresetConfig,
		TextQuota:            params.TextQuota,
		ImageQuota:           params.ImageQuota,
		AudioQuota:           params.AudioQuota,
		MultimodalQuota:      params.MultimodalQuota,
		RealtimeQuota:        params.RealtimeQuota,
		MultimodalAudioQuota: params.MultimodalAudioQuota,
		MidjourneyQuotas:     params.MidjourneyQuotas,
		RequestDataFormat:    params.RequestDataFormat,
		ResponseDataFormat:   params.ResponseDataFormat,
		IsPublic:             params.IsPublic,
		IsEnableModelAgent:   params.IsEnableModelAgent,
		LbStrategy:           params.LbStrategy,
		ModelAgents:          params.ModelAgents,
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

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_MODEL, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.IsPublic {

		resellerList, err := dao.Reseller.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.Reseller.UpdateMany(ctx, bson.M{}, bson.M{
			"$push": bson.M{
				"models": id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, reseller := range resellerList {

			newResellerData := *reseller

			newResellerData.Models = append(newResellerData.Models, id)

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: reseller,
				NewData: newResellerData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		userList, err := dao.User.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.User.UpdateMany(ctx, bson.M{}, bson.M{
			"$push": bson.M{
				"models": id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, user := range userList {

			newUserData := *user

			newUserData.Models = append(newUserData.Models, id)

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: user,
				NewData: newUserData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
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

	for _, modelAgentId := range params.ModelAgents {

		if !slices.Contains(oldData.ModelAgents, modelAgentId) {

			modelAgent, err := service.ModelAgent().Detail(ctx, modelAgentId)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			if err = service.ModelAgent().Update(ctx, model.ModelAgentUpdateReq{
				Id:                   modelAgent.Id,
				Corp:                 modelAgent.Corp,
				Name:                 modelAgent.Name,
				BaseUrl:              modelAgent.BaseUrl,
				Path:                 modelAgent.Path,
				Weight:               modelAgent.Weight,
				Models:               append(modelAgent.Models, params.Id),
				IsEnableModelReplace: modelAgent.IsEnableModelReplace,
				ReplaceModels:        modelAgent.ReplaceModels,
				TargetModels:         modelAgent.TargetModels,
				IsNeverDisable:       modelAgent.IsNeverDisable,
				LbStrategy:           modelAgent.LbStrategy,
				Key:                  modelAgent.Key,
				Remark:               modelAgent.Remark,
				Status:               modelAgent.Status,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	for _, modelAgentId := range oldData.ModelAgents {

		if !slices.Contains(params.ModelAgents, modelAgentId) {

			modelAgent, err := service.ModelAgent().Detail(ctx, modelAgentId)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			for i, modelId := range modelAgent.Models {
				if modelId == params.Id {
					modelAgent.Models = util.Delete(modelAgent.Models, i)
				}
			}

			if err = service.ModelAgent().Update(ctx, model.ModelAgentUpdateReq{
				Id:                   modelAgent.Id,
				Corp:                 modelAgent.Corp,
				Name:                 modelAgent.Name,
				BaseUrl:              modelAgent.BaseUrl,
				Path:                 modelAgent.Path,
				Weight:               modelAgent.Weight,
				Models:               append(modelAgent.Models, params.Id),
				IsEnableModelReplace: modelAgent.IsEnableModelReplace,
				ReplaceModels:        modelAgent.ReplaceModels,
				TargetModels:         modelAgent.TargetModels,
				IsNeverDisable:       modelAgent.IsNeverDisable,
				LbStrategy:           modelAgent.LbStrategy,
				Key:                  modelAgent.Key,
				Remark:               modelAgent.Remark,
				Status:               modelAgent.Status,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	m := &do.Model{
		Corp:                 params.Corp,
		Name:                 gstr.Trim(params.Name),
		Model:                gstr.Trim(params.Model),
		Type:                 params.Type,
		BaseUrl:              gstr.Trim(params.BaseUrl),
		Path:                 gstr.Trim(params.Path),
		IsEnablePresetConfig: params.IsEnablePresetConfig,
		PresetConfig:         params.PresetConfig,
		TextQuota:            params.TextQuota,
		ImageQuota:           params.ImageQuota,
		AudioQuota:           params.AudioQuota,
		MultimodalQuota:      params.MultimodalQuota,
		RealtimeQuota:        params.RealtimeQuota,
		MultimodalAudioQuota: params.MultimodalAudioQuota,
		MidjourneyQuotas:     params.MidjourneyQuotas,
		RequestDataFormat:    params.RequestDataFormat,
		ResponseDataFormat:   params.ResponseDataFormat,
		IsPublic:             params.IsPublic,
		IsEnableModelAgent:   params.IsEnableModelAgent,
		LbStrategy:           params.LbStrategy,
		ModelAgents:          params.ModelAgents,
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

	// 旧数据是公开, 新数据改为了私有
	if oldData.IsPublic && !newData.IsPublic {

		resellerList, err := dao.Reseller.Find(ctx, bson.M{"models": bson.M{"$in": []string{params.Id}}})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.Reseller.UpdateMany(ctx, bson.M{"models": bson.M{"$in": []string{params.Id}}}, bson.M{
			"$pull": bson.M{
				"models": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, reseller := range resellerList {

			newResellerData := *reseller

			for i, id := range newResellerData.Models {
				if id == params.Id {
					newResellerData.Models = append(newResellerData.Models[:i], newResellerData.Models[i+1:]...)
					break
				}
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: reseller,
				NewData: newResellerData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		userList, err := dao.User.Find(ctx, bson.M{"models": bson.M{"$in": []string{params.Id}}})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.User.UpdateMany(ctx, bson.M{"models": bson.M{"$in": []string{params.Id}}}, bson.M{
			"$pull": bson.M{
				"models": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, user := range userList {

			newUserData := *user

			for i, id := range newUserData.Models {
				if id == params.Id {
					newUserData.Models = append(newUserData.Models[:i], newUserData.Models[i+1:]...)
					break
				}
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: user,
				NewData: newUserData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

	} else if !oldData.IsPublic && newData.IsPublic { // 旧数据是私有, 新数据改为了公开

		resellerList, err := dao.Reseller.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.Reseller.UpdateMany(ctx, bson.M{}, bson.M{
			"$addToSet": bson.M{
				"models": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, reseller := range resellerList {

			newResellerData := *reseller

			newResellerData.Models = gset.NewStrSetFrom(append(newResellerData.Models, params.Id)).Slice()

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: reseller,
				NewData: newResellerData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		userList, err := dao.User.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.User.UpdateMany(ctx, bson.M{}, bson.M{
			"$addToSet": bson.M{
				"models": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, user := range userList {

			newUserData := *user

			newUserData.Models = gset.NewStrSetFrom(append(newUserData.Models, params.Id)).Slice()

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: user,
				NewData: newUserData,
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

	resellers, err := dao.Reseller.Find(ctx, bson.M{"models": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, reseller := range resellers {

		resellerModelsReq := model.ResellerPermissionsReq{
			UserId: reseller.UserId,
			Models: []string{},
			Groups: reseller.Groups,
		}

		for _, m := range reseller.Models {
			if m != id {
				resellerModelsReq.Models = append(resellerModelsReq.Models, m)
			}
		}

		if err = service.AdminReseller().Permissions(ctx, resellerModelsReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	users, err := dao.User.Find(ctx, bson.M{"models": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, user := range users {

		userPermissionsReq := model.UserPermissionsReq{
			UserId: user.UserId,
			Models: []string{},
			Groups: user.Groups,
		}

		for _, m := range user.Models {
			if m != id {
				userPermissionsReq.Models = append(userPermissionsReq.Models, m)
			}
		}

		if err = service.AdminUser().Permissions(ctx, userPermissionsReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
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

	modelAgentNames := make([]string, 0)

	if len(m.ModelAgents) > 0 {

		modelAgentList, err := dao.ModelAgent.Find(ctx, bson.M{"_id": bson.M{"$in": m.ModelAgents}})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		for _, modelAgent := range modelAgentList {
			modelAgentNames = append(modelAgentNames, modelAgent.Name)
		}
	}

	corpName := m.Corp
	corpCode := m.Corp
	if corp, err := dao.Corp.FindById(ctx, m.Corp); err == nil && corp != nil {
		corpName = corp.Name
		corpCode = corp.Code
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
		Corp:                 m.Corp,
		CorpName:             corpName,
		CorpCode:             corpCode,
		Name:                 m.Name,
		Model:                m.Model,
		Type:                 m.Type,
		BaseUrl:              m.BaseUrl,
		Path:                 m.Path,
		Groups:               groupIds,
		GroupNames:           groupNames,
		IsEnablePresetConfig: m.IsEnablePresetConfig,
		PresetConfig:         m.PresetConfig,
		TextQuota:            m.TextQuota,
		ImageQuota:           m.ImageQuota,
		AudioQuota:           m.AudioQuota,
		MultimodalQuota:      m.MultimodalQuota,
		RealtimeQuota:        m.RealtimeQuota,
		MultimodalAudioQuota: m.MultimodalAudioQuota,
		MidjourneyQuotas:     m.MidjourneyQuotas,
		RequestDataFormat:    m.RequestDataFormat,
		ResponseDataFormat:   m.ResponseDataFormat,
		IsPublic:             m.IsPublic,
		IsEnableModelAgent:   m.IsEnableModelAgent,
		LbStrategy:           m.LbStrategy,
		ModelAgents:          m.ModelAgents,
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

		if params.Group != "" && slices.Contains(reseller.Groups, params.Group) {

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

			modelSet := gset.NewStrSetFrom(reseller.Models)

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
	}

	if service.Session().IsUserRole(ctx) {

		user, err := service.User().GetUserByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if params.Group != "" && slices.Contains(user.Groups, params.Group) {

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

			modelSet := gset.NewStrSetFrom(user.Models)

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

	if params.Corp != "" {
		filter["corp"] = params.Corp
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

	corps, err := dao.Corp.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	corpMap := util.ToMap(corps, func(t *entity.Corp) string {
		return t.Id
	})

	groups, err := service.Group().List(ctx, model.GroupListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Model, 0)
	for _, result := range results {

		corpName := result.Corp
		corpCode := result.Corp
		if corpMap[result.Corp] != nil {
			corpName = corpMap[result.Corp].Name
			corpCode = corpMap[result.Corp].Code
		}

		groupNames := make([]string, 0)
		for _, group := range groups {
			if slices.Contains(group.Models, result.Id) {
				groupNames = append(groupNames, group.Name)
			}
		}

		model := &model.Model{
			Id:                   result.Id,
			Corp:                 result.Corp,
			CorpName:             corpName,
			CorpCode:             corpCode,
			Name:                 result.Name,
			Model:                result.Model,
			Type:                 result.Type,
			GroupNames:           groupNames,
			TextQuota:            result.TextQuota,
			ImageQuota:           result.ImageQuota,
			AudioQuota:           result.AudioQuota,
			MultimodalQuota:      result.MultimodalQuota,
			RealtimeQuota:        result.RealtimeQuota,
			MultimodalAudioQuota: result.MultimodalAudioQuota,
			MidjourneyQuotas:     result.MidjourneyQuotas,
			RequestDataFormat:    result.RequestDataFormat,
			ResponseDataFormat:   result.ResponseDataFormat,
			IsPublic:             result.IsPublic,
			Remark:               result.Remark,
			Status:               result.Status,
			CreatedAt:            util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:            util.FormatDateTimeMonth(result.UpdatedAt),
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

	if params.Corp != "" {
		filter["corp"] = params.Corp
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

		modelSet := gset.NewStrSetFrom(reseller.Models)

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

		modelSet := gset.NewStrSetFrom(user.Models)

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

	corps, err := dao.Corp.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	corpMap := util.ToMap(corps, func(t *entity.Corp) string {
		return t.Id
	})

	items := make([]*model.Model, 0)
	for _, result := range results {

		corpName := result.Corp
		if corpMap[result.Corp] != nil {
			corpName = corpMap[result.Corp].Name
		}

		model := &model.Model{
			Id:       result.Id,
			Corp:     result.Corp,
			CorpName: corpName,
			Name:     result.Name,
			Model:    result.Model,
			Type:     result.Type,
			Status:   result.Status,
		}

		if service.Session().IsAdminRole(ctx) {
			model.ModelAgents = result.ModelAgents
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
				Corp:                 result.Corp,
				Name:                 result.Name,
				Model:                result.Model,
				Type:                 result.Type,
				BaseUrl:              result.BaseUrl,
				Path:                 result.Path,
				IsEnablePresetConfig: result.IsEnablePresetConfig,
				PresetConfig:         result.PresetConfig,
				TextQuota:            result.TextQuota,
				ImageQuota:           result.ImageQuota,
				AudioQuota:           result.AudioQuota,
				MultimodalQuota:      result.MultimodalQuota,
				RealtimeQuota:        result.RealtimeQuota,
				MultimodalAudioQuota: result.MultimodalAudioQuota,
				MidjourneyQuotas:     result.MidjourneyQuotas,
				RequestDataFormat:    result.RequestDataFormat,
				ResponseDataFormat:   result.ResponseDataFormat,
				IsPublic:             result.IsPublic,
				Groups:               groupIds,
				LbStrategy:           result.LbStrategy,
				ModelAgents:          result.ModelAgents,
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
				if m.IsEnableModelAgent && len(m.ModelAgents) == 0 {
					continue
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
				Corp:                 result.Corp,
				Name:                 result.Name,
				Model:                result.Model,
				Type:                 result.Type,
				BaseUrl:              result.BaseUrl,
				Path:                 result.Path,
				IsEnablePresetConfig: result.IsEnablePresetConfig,
				PresetConfig:         result.PresetConfig,
				TextQuota:            result.TextQuota,
				ImageQuota:           result.ImageQuota,
				AudioQuota:           result.AudioQuota,
				MultimodalQuota:      result.MultimodalQuota,
				RealtimeQuota:        result.RealtimeQuota,
				MultimodalAudioQuota: result.MultimodalAudioQuota,
				MidjourneyQuotas:     result.MidjourneyQuotas,
				RequestDataFormat:    result.RequestDataFormat,
				ResponseDataFormat:   result.ResponseDataFormat,
				IsPublic:             result.IsPublic,
				Groups:               groupIds,
				IsEnableModelAgent:   result.IsEnableModelAgent,
				LbStrategy:           result.LbStrategy,
				ModelAgents:          result.ModelAgents,
				ForwardConfig:        result.ForwardConfig,
				IsEnableFallback:     result.IsEnableFallback,
				FallbackConfig:       result.FallbackConfig,
				Remark:               result.Remark,
				Status:               result.Status,
			}

			if params.Value == "all" {

				if m.ForwardConfig == nil {
					m.ForwardConfig = new(common.ForwardConfig)
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
				Corp:                 result.Corp,
				Name:                 result.Name,
				Model:                result.Model,
				Type:                 result.Type,
				BaseUrl:              result.BaseUrl,
				Path:                 result.Path,
				IsEnablePresetConfig: result.IsEnablePresetConfig,
				PresetConfig:         result.PresetConfig,
				TextQuota:            result.TextQuota,
				ImageQuota:           result.ImageQuota,
				AudioQuota:           result.AudioQuota,
				MultimodalQuota:      result.MultimodalQuota,
				RealtimeQuota:        result.RealtimeQuota,
				MultimodalAudioQuota: result.MultimodalAudioQuota,
				MidjourneyQuotas:     result.MidjourneyQuotas,
				RequestDataFormat:    result.RequestDataFormat,
				ResponseDataFormat:   result.ResponseDataFormat,
				IsPublic:             result.IsPublic,
				Groups:               groupIds,
				IsEnableModelAgent:   result.IsEnableModelAgent,
				LbStrategy:           result.LbStrategy,
				ModelAgents:          result.ModelAgents,
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

		modelSet := gset.NewStrSetFrom(reseller.Models)

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

		modelSet := gset.NewStrSetFrom(user.Models)

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

	treeData := make(map[string][]model.Tree) // [Corp:Type][][Model]

	for _, result := range results {

		corpTree := treeData[fmt.Sprintf("%s:%d", result.Corp, result.Type)]
		if corpTree == nil {
			corpTree = make([]model.Tree, 0)
		}

		corpTree = append(corpTree, model.Tree{
			Title: result.Model,
			Value: result.Id,
			Key:   result.Id,
		})

		treeData[fmt.Sprintf("%s:%d", result.Corp, result.Type)] = corpTree
	}

	corps, err := service.Corp().List(ctx, model.CorpListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Tree, 0)
	for _, corp := range corps {

		children := make([]model.Tree, 0)

		for _, typ := range consts.MODEL_TYPES {
			if modelTree := treeData[fmt.Sprintf("%s:%d", corp.Id, typ)]; modelTree != nil {
				children = append(children, model.Tree{
					Title:    consts.MODEL_TYPE[typ],
					Value:    fmt.Sprintf("%s:%d", corp.Id, typ),
					Key:      fmt.Sprintf("%s:%d", corp.Id, typ),
					Children: modelTree,
				})
			}
		}

		if len(children) > 0 {
			items = append(items, &model.Tree{
				Title:    corp.Name,
				Value:    corp.Id,
				Key:      corp.Id,
				Children: children,
			})
		}
	}

	return items, nil
}

// 模型权限列表
func (s *sModel) Permissions(ctx context.Context, params model.ModelPermissionsReq) ([]*model.Model, error) {

	modelListReq := model.ModelListReq{
		Corp:   params.Corp,
		Name:   params.Name,
		Model:  params.Model,
		Type:   params.Type,
		Status: params.Status,
	}

	switch params.Action {
	case consts.ACTION_USER:

		user, err := service.AdminUser().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if user.Models == nil {
			return nil, nil
		}

		modelListReq.Models = user.Models

	case consts.ACTION_APP:

		app, err := service.App().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if app.Models == nil {
			return nil, nil
		}

		modelListReq.Models = app.Models

	case consts.ACTION_KEY:

		key, err := service.Key().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if key.Models == nil {
			return nil, nil
		}

		modelListReq.Models = key.Models

	case consts.ACTION_AGENT:

		modelAgent, err := service.ModelAgent().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if modelAgent.Models == nil {
			return nil, nil
		}

		modelListReq.Models = modelAgent.Models

	case consts.ACTION_FALLBACK:

		modelAgent, err := service.ModelAgent().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if modelAgent.FallbackModels == nil {
			return nil, nil
		}

		modelListReq.Models = modelAgent.FallbackModels

	case consts.ACTION_GROUP:

		group, err := service.Group().Detail(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if group.Models == nil {
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
	if err := util.HttpGet(ctx, params.Url, g.MapStrStr{"Authorization": "Bearer " + params.Key}, g.MapStrAny{"is_fastapi": true}, &result); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if result.Data == nil || (len(result.Data) > 0 && result.Data[0].FastAPI == nil) {
		return errors.New("模型接口数据格式不支持, 请联系作者...")
	}

	corps, err := dao.Corp.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	corpNameMap := make(map[string]string)
	corpCodeMap := make(map[string]string)
	for _, corp := range corps {
		corpNameMap[corp.Name] = corp.Id
		corpCodeMap[corp.Code] = corp.Id
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

		if modelAgent, err := dao.ModelAgent.FindOne(ctx, bson.M{"name": "FastAPI"}); err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {

				if corpCodeMap["FastAPI"] == "" {
					if corpCodeMap["FastAPI"], err = service.Corp().Create(ctx, model.CorpCreateReq{
						Name:     "FastAPI",
						Code:     "FastAPI",
						IsPublic: true,
						Status:   1,
					}); err != nil {
						logger.Error(ctx, err)
						return err
					}
				}

				if modelAgentId, err = service.ModelAgent().Create(ctx, model.ModelAgentCreateReq{
					Corp:         corpCodeMap["FastAPI"],
					Name:         "FastAPI",
					BaseUrl:      gstr.Replace(params.Url, "/models", ""),
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

	for _, data := range result.Data {

		corp := corpNameMap[data.FastAPI.Corp]
		if corp == "" {
			corp = corpCodeMap[data.FastAPI.Code]
		}

		if corp == "" {
			if corp, err = service.Corp().Create(ctx, model.CorpCreateReq{
				Name:     data.FastAPI.Corp,
				Code:     data.FastAPI.Code,
				IsPublic: true,
				Status:   1,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
			corpNameMap[data.FastAPI.Corp] = corp
			corpCodeMap[data.FastAPI.Code] = corp
		}

		if modelMap[data.Id] == nil {

			modelCreateReq := model.ModelCreateReq{
				Corp:                 corp,
				Name:                 data.Id,
				Model:                data.FastAPI.Model,
				Type:                 data.FastAPI.Type,
				TextQuota:            data.FastAPI.TextQuota,
				ImageQuota:           data.FastAPI.ImageQuota,
				AudioQuota:           data.FastAPI.AudioQuota,
				MultimodalQuota:      data.FastAPI.MultimodalQuota,
				RealtimeQuota:        data.FastAPI.RealtimeQuota,
				MultimodalAudioQuota: data.FastAPI.MultimodalAudioQuota,
				MidjourneyQuotas:     data.FastAPI.MidjourneyQuotas,
				RequestDataFormat:    1,
				ResponseDataFormat:   1,
				IsPublic:             true,
				ModelAgents:          []string{},
				Remark:               data.FastAPI.Remark,
				Status:               1,
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
				Corp:                 detail.Corp,
				Name:                 detail.Name,
				Model:                detail.Model,
				Type:                 detail.Type,
				BaseUrl:              detail.BaseUrl,
				Path:                 detail.Path,
				IsEnablePresetConfig: detail.IsEnablePresetConfig,
				PresetConfig:         detail.PresetConfig,
				TextQuota:            data.FastAPI.TextQuota,
				ImageQuota:           data.FastAPI.ImageQuota,
				AudioQuota:           data.FastAPI.AudioQuota,
				MultimodalQuota:      data.FastAPI.MultimodalQuota,
				RealtimeQuota:        data.FastAPI.RealtimeQuota,
				MultimodalAudioQuota: data.FastAPI.MultimodalAudioQuota,
				MidjourneyQuotas:     data.FastAPI.MidjourneyQuotas,
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
