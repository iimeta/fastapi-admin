package model

import (
	"context"
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
	"slices"
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
		BaseUrl:              params.BaseUrl,
		Path:                 params.Path,
		IsEnablePresetConfig: params.IsEnablePresetConfig,
		PresetConfig:         params.PresetConfig,
		TextQuota:            params.TextQuota,
		ImageQuotas:          params.ImageQuotas,
		MultimodalQuota:      params.MultimodalQuota,
		MidjourneyQuotas:     params.MidjourneyQuotas,
		DataFormat:           params.DataFormat,
		IsPublic:             params.IsPublic,
		IsEnableModelAgent:   params.IsEnableModelAgent,
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
				Action:  consts.ACTION_MODELS,
				OldData: user,
				NewData: newUserData,
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
				Id:      modelAgent.Id,
				Corp:    modelAgent.Corp,
				Name:    modelAgent.Name,
				BaseUrl: modelAgent.BaseUrl,
				Path:    modelAgent.Path,
				Weight:  modelAgent.Weight,
				Models:  append(modelAgent.Models, params.Id),
				Key:     modelAgent.Key,
				Remark:  modelAgent.Remark,
				Status:  modelAgent.Status,
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
				Id:      modelAgent.Id,
				Corp:    modelAgent.Corp,
				Name:    modelAgent.Name,
				BaseUrl: modelAgent.BaseUrl,
				Path:    modelAgent.Path,
				Weight:  modelAgent.Weight,
				Models:  modelAgent.Models,
				Key:     modelAgent.Key,
				Remark:  modelAgent.Remark,
				Status:  modelAgent.Status,
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
		BaseUrl:              params.BaseUrl,
		Path:                 params.Path,
		IsEnablePresetConfig: params.IsEnablePresetConfig,
		PresetConfig:         params.PresetConfig,
		TextQuota:            params.TextQuota,
		ImageQuotas:          params.ImageQuotas,
		MultimodalQuota:      params.MultimodalQuota,
		MidjourneyQuotas:     params.MidjourneyQuotas,
		DataFormat:           params.DataFormat,
		IsPublic:             params.IsPublic,
		IsEnableModelAgent:   params.IsEnableModelAgent,
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
				Action:  consts.ACTION_MODELS,
				OldData: user,
				NewData: newUserData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

	} else if !oldData.IsPublic && newData.IsPublic { // 旧数据是私有, 新数据改为了公开

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
				Action:  consts.ACTION_MODELS,
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

	users, err := dao.User.Find(ctx, bson.M{"models": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, user := range users {

		userModelsReq := model.UserModelsReq{
			UserId: user.UserId,
			Models: []string{},
		}

		for _, m := range user.Models {
			if m != id {
				userModelsReq.Models = append(userModelsReq.Models, m)
			}
		}

		if err = service.AdminUser().Models(ctx, userModelsReq); err != nil {
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
		IsEnablePresetConfig: m.IsEnablePresetConfig,
		PresetConfig:         m.PresetConfig,
		TextQuota:            m.TextQuota,
		ImageQuotas:          m.ImageQuotas,
		MultimodalQuota:      m.MultimodalQuota,
		MidjourneyQuotas:     m.MidjourneyQuotas,
		DataFormat:           m.DataFormat,
		IsPublic:             m.IsPublic,
		IsEnableModelAgent:   m.IsEnableModelAgent,
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
			modelNames, err := s.ModelNames(ctx, []string{detail.ForwardConfig.TargetModel})
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			detail.ForwardConfig.TargetModelName = modelNames[0]
		}

		if detail.ForwardConfig.DecisionModel != "" {
			modelNames, err := s.ModelNames(ctx, []string{detail.ForwardConfig.DecisionModel})
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			detail.ForwardConfig.DecisionModelName = modelNames[0]
		}

		if detail.ForwardConfig.TargetModels != nil && len(detail.ForwardConfig.TargetModels) > 0 {
			modelNames, err := s.ModelNames(ctx, detail.ForwardConfig.TargetModels)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			detail.ForwardConfig.TargetModelNames = modelNames
		}
	}

	if detail.FallbackConfig != nil {

		if detail.FallbackConfig.FallbackModel != "" {
			modelNames, err := s.ModelNames(ctx, []string{detail.FallbackConfig.FallbackModel})
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			detail.FallbackConfig.FallbackModelName = modelNames[0]
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

	if service.Session().IsUserRole(ctx) {

		user, err := service.User().GetUserByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(user.Models) == 0 {
			return nil, nil
		}

		filter["_id"] = bson.M{
			"$in": user.Models,
		}
	}

	if params.Corp != "" {
		filter["corp"] = params.Corp
	}

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex": params.Name,
		}
	}

	if params.Model != "" {
		filter["model"] = bson.M{
			"$regex": params.Model,
		}
	}

	if params.Type != 0 {
		filter["type"] = params.Type
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	results, err := dao.Model.FindByPage(ctx, paging, filter, "status", "-updated_at")
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

		items = append(items, &model.Model{
			Id:               result.Id,
			Corp:             result.Corp,
			CorpName:         corpName,
			Name:             result.Name,
			Model:            result.Model,
			Type:             result.Type,
			TextQuota:        result.TextQuota,
			ImageQuotas:      result.ImageQuotas,
			MultimodalQuota:  result.MultimodalQuota,
			MidjourneyQuotas: result.MidjourneyQuotas,
			DataFormat:       result.DataFormat,
			IsPublic:         result.IsPublic,
			Status:           result.Status,
			CreatedAt:        util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:        util.FormatDateTimeMonth(result.UpdatedAt),
		})
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

	if service.Session().IsUserRole(ctx) {

		models := service.Session().GetUser(ctx).Models
		if len(models) == 0 {
			return nil, nil
		}

		filter["_id"] = bson.M{
			"$in": models,
		}
	}

	results, err := dao.Model.Find(ctx, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Model, 0)
	for _, result := range results {

		model := &model.Model{
			Id:         result.Id,
			Corp:       result.Corp,
			Name:       result.Name,
			Model:      result.Model,
			Type:       result.Type,
			DataFormat: result.DataFormat,
			Status:     result.Status,
		}

		if service.Session().IsAdminRole(ctx) {
			model.ModelAgents = result.ModelAgents
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

		for _, result := range results {

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
				ImageQuotas:          result.ImageQuotas,
				MultimodalQuota:      result.MultimodalQuota,
				MidjourneyQuotas:     result.MidjourneyQuotas,
				DataFormat:           result.DataFormat,
				IsPublic:             result.IsPublic,
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

		for _, result := range results {

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
				ImageQuotas:          result.ImageQuotas,
				MultimodalQuota:      result.MultimodalQuota,
				MidjourneyQuotas:     result.MidjourneyQuotas,
				DataFormat:           result.DataFormat,
				IsPublic:             result.IsPublic,
				IsEnableModelAgent:   result.IsEnableModelAgent,
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

		for _, result := range results {

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
				ImageQuotas:          result.ImageQuotas,
				MultimodalQuota:      result.MultimodalQuota,
				MidjourneyQuotas:     result.MidjourneyQuotas,
				DataFormat:           result.DataFormat,
				IsPublic:             result.IsPublic,
				IsEnableModelAgent:   result.IsEnableModelAgent,
				ModelAgents:          result.ModelAgents,
				IsEnableForward:      result.IsEnableForward,
				ForwardConfig:        result.ForwardConfig,
				FallbackConfig:       result.FallbackConfig,
				Remark:               result.Remark,
				Status:               result.Status,
			}

			if params.Value == "all" {

				if m.FallbackConfig == nil {
					m.FallbackConfig = new(common.FallbackConfig)
				}

				m.IsEnableFallback = true
				m.FallbackConfig.FallbackModel = params.FallbackModel

			} else {
				m.IsEnableFallback = gconv.Bool(params.Value)
				if m.IsEnableFallback && (m.FallbackConfig == nil || m.FallbackConfig.FallbackModel == "") {
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

// 模型初始化
func (s *sModel) Init(ctx context.Context, params model.ModelInitReq) error {

	result := &model.ModelsRes{}
	if err := util.HttpGet(ctx, params.Url, g.MapStrStr{"Authorization": "Bearer " + params.Key}, nil, &result); err != nil {
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

	corpMap := make(map[string]string)
	for _, corp := range corps {
		corpMap[corp.Code] = corp.Id
	}

	modelAgentId := ""
	if params.IsConfigModelAgent {

		if modelAgent, err := dao.ModelAgent.FindOne(ctx, bson.M{"name": "FastAPI"}); err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {

				if corpMap["FastAPI"] == "" {
					if corpMap["FastAPI"], err = service.Corp().Create(ctx, model.CorpCreateReq{
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
					Corp:         corpMap["FastAPI"],
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

		code := data.FastAPI.Code
		if params.IsConfigModelAgent {
			code += "2FastAPI"
		}

		if corpMap[code] == "" {
			if corpMap[code], err = service.Corp().Create(ctx, model.CorpCreateReq{
				Name:     data.FastAPI.Corp,
				Code:     code,
				IsPublic: true,
				Status:   1,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		modelCreateReq := model.ModelCreateReq{
			Corp:             corpMap[code],
			Name:             data.Id,
			Model:            data.Id,
			Type:             data.FastAPI.Type,
			BaseUrl:          data.FastAPI.BaseUrl,
			Path:             data.FastAPI.Path,
			TextQuota:        data.FastAPI.TextQuota,
			ImageQuotas:      data.FastAPI.ImageQuotas,
			MultimodalQuota:  data.FastAPI.MultimodalQuota,
			MidjourneyQuotas: data.FastAPI.MidjourneyQuotas,
			DataFormat:       1,
			IsPublic:         true,
			Status:           1,
		}

		if params.IsConfigModelAgent && modelAgentId != "" {
			modelCreateReq.IsEnableModelAgent = true
			modelCreateReq.ModelAgents = append(modelCreateReq.ModelAgents, modelAgentId)
		}

		if err = s.Create(ctx, modelCreateReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

// 公开的模型Ids
func (s *sModel) PublicModels(ctx context.Context) ([]string, error) {

	filter := bson.M{
		"is_public": true,
		"status":    1,
	}

	results, err := dao.Model.Find(ctx, filter, "-updated_at")
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
