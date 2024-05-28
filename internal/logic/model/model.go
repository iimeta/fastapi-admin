package model

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
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
		Corp:               params.Corp,
		Name:               gstr.Trim(params.Name),
		Model:              gstr.Trim(params.Model),
		Type:               params.Type,
		BaseUrl:            params.BaseUrl,
		Path:               params.Path,
		Prompt:             params.Prompt,
		BillingMethod:      params.BillingMethod,
		PromptRatio:        params.PromptRatio,
		CompletionRatio:    params.CompletionRatio,
		FixedQuota:         params.FixedQuota,
		DataFormat:         params.DataFormat,
		IsPublic:           params.IsPublic,
		IsEnableModelAgent: params.IsEnableModelAgent,
		ModelAgents:        params.ModelAgents,
		IsForward:          params.IsForward,
		Remark:             params.Remark,
		Status:             params.Status,
	}

	if params.ForwardConfig != nil {
		m.ForwardConfig = &do.ForwardConfig{
			ForwardRule:   params.ForwardConfig.ForwardRule,
			MatchRule:     params.ForwardConfig.MatchRule,
			TargetModel:   params.ForwardConfig.TargetModel,
			DecisionModel: params.ForwardConfig.DecisionModel,
			Keywords:      params.ForwardConfig.Keywords,
			TargetModels:  params.ForwardConfig.TargetModels,
		}
	}

	id, err := dao.Model.Insert(ctx, m)
	if err != nil {
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
		Corp:               params.Corp,
		Name:               gstr.Trim(params.Name),
		Model:              gstr.Trim(params.Model),
		Type:               params.Type,
		BaseUrl:            params.BaseUrl,
		Path:               params.Path,
		Prompt:             params.Prompt,
		BillingMethod:      params.BillingMethod,
		PromptRatio:        params.PromptRatio,
		CompletionRatio:    params.CompletionRatio,
		FixedQuota:         params.FixedQuota,
		DataFormat:         params.DataFormat,
		IsPublic:           params.IsPublic,
		IsEnableModelAgent: params.IsEnableModelAgent,
		ModelAgents:        params.ModelAgents,
		IsForward:          params.IsForward,
		Remark:             params.Remark,
		Status:             params.Status,
	}

	if params.ForwardConfig != nil {
		m.ForwardConfig = &do.ForwardConfig{
			ForwardRule:   params.ForwardConfig.ForwardRule,
			MatchRule:     params.ForwardConfig.MatchRule,
			TargetModel:   params.ForwardConfig.TargetModel,
			DecisionModel: params.ForwardConfig.DecisionModel,
			Keywords:      params.ForwardConfig.Keywords,
			TargetModels:  params.ForwardConfig.TargetModels,
		}
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

	if err = dao.User.UpdateMany(ctx, bson.M{"models": bson.M{"$in": []string{id}}}, bson.M{
		"$pull": bson.M{
			"models": id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.App.UpdateMany(ctx, bson.M{"models": bson.M{"$in": []string{id}}}, bson.M{
		"$pull": bson.M{
			"models": id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.Key.UpdateMany(ctx, bson.M{"models": bson.M{"$in": []string{id}}}, bson.M{
		"$pull": bson.M{
			"models": id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
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

	corp, err := dao.Corp.FindById(ctx, m.Corp)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	corpName := m.Corp
	if corp != nil {
		corpName = corp.Name
	}

	detail := &model.Model{
		Id:                 m.Id,
		Corp:               corpName,
		Name:               m.Name,
		Model:              m.Model,
		Type:               m.Type,
		BaseUrl:            m.BaseUrl,
		Path:               m.Path,
		Prompt:             m.Prompt,
		BillingMethod:      m.BillingMethod,
		PromptRatio:        m.PromptRatio,
		PromptPrice:        util.PriceConv(m.PromptRatio),
		CompletionRatio:    m.CompletionRatio,
		CompletionPrice:    util.PriceConv(m.CompletionRatio),
		FixedQuota:         m.FixedQuota,
		FixedPrice:         util.QuotaConv(m.FixedQuota),
		DataFormat:         m.DataFormat,
		IsPublic:           m.IsPublic,
		IsEnableModelAgent: m.IsEnableModelAgent,
		ModelAgents:        m.ModelAgents,
		ModelAgentNames:    modelAgentNames,
		IsForward:          m.IsForward,
		Remark:             m.Remark,
		Status:             m.Status,
		CreatedAt:          util.FormatDateTime(m.CreatedAt),
		UpdatedAt:          util.FormatDateTime(m.UpdatedAt),
	}

	if m.ForwardConfig != nil {

		detail.ForwardConfig = &model.ForwardConfig{
			ForwardRule:   m.ForwardConfig.ForwardRule,
			MatchRule:     m.ForwardConfig.MatchRule,
			TargetModel:   m.ForwardConfig.TargetModel,
			DecisionModel: m.ForwardConfig.DecisionModel,
			Keywords:      m.ForwardConfig.Keywords,
			TargetModels:  m.ForwardConfig.TargetModels,
		}

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

		models := service.Session().GetUser(ctx).Models
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
			Id:              result.Id,
			Corp:            corpName,
			Name:            result.Name,
			Model:           result.Model,
			Type:            result.Type,
			BillingMethod:   result.BillingMethod,
			PromptRatio:     result.PromptRatio,
			PromptPrice:     util.PriceConv(result.PromptRatio),
			CompletionRatio: result.CompletionRatio,
			CompletionPrice: util.PriceConv(result.CompletionRatio),
			FixedQuota:      result.FixedQuota,
			FixedPrice:      util.QuotaConv(result.FixedQuota),
			DataFormat:      result.DataFormat,
			IsPublic:        result.IsPublic,
			Status:          result.Status,
			CreatedAt:       util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:       util.FormatDateTimeMonth(result.UpdatedAt),
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
			Id:              result.Id,
			Corp:            result.Corp,
			Name:            result.Name,
			Model:           result.Model,
			Type:            result.Type,
			BillingMethod:   result.BillingMethod,
			PromptRatio:     result.PromptRatio,
			CompletionRatio: result.CompletionRatio,
			FixedQuota:      result.FixedQuota,
			DataFormat:      result.DataFormat,
			Status:          result.Status,
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
				Id:              result.Id,
				Corp:            result.Corp,
				Name:            result.Name,
				Model:           result.Model,
				Type:            result.Type,
				BaseUrl:         result.BaseUrl,
				Path:            result.Path,
				Prompt:          result.Prompt,
				BillingMethod:   result.BillingMethod,
				PromptRatio:     result.PromptRatio,
				CompletionRatio: result.CompletionRatio,
				FixedQuota:      result.FixedQuota,
				DataFormat:      result.DataFormat,
				IsPublic:        result.IsPublic,
				ModelAgents:     result.ModelAgents,
				IsForward:       result.IsForward,
				Remark:          result.Remark,
				Status:          result.Status,
			}

			if result.ForwardConfig != nil {
				m.ForwardConfig = &model.ForwardConfig{
					ForwardRule:   result.ForwardConfig.ForwardRule,
					MatchRule:     result.ForwardConfig.MatchRule,
					TargetModel:   result.ForwardConfig.TargetModel,
					DecisionModel: result.ForwardConfig.DecisionModel,
					Keywords:      result.ForwardConfig.Keywords,
					TargetModels:  result.ForwardConfig.TargetModels,
				}
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
				Id:                 result.Id,
				Corp:               result.Corp,
				Name:               result.Name,
				Model:              result.Model,
				Type:               result.Type,
				BaseUrl:            result.BaseUrl,
				Path:               result.Path,
				Prompt:             result.Prompt,
				BillingMethod:      result.BillingMethod,
				PromptRatio:        result.PromptRatio,
				CompletionRatio:    result.CompletionRatio,
				FixedQuota:         result.FixedQuota,
				DataFormat:         result.DataFormat,
				IsPublic:           result.IsPublic,
				IsEnableModelAgent: result.IsEnableModelAgent,
				ModelAgents:        result.ModelAgents,
				Remark:             result.Remark,
				Status:             result.Status,
			}

			if result.ForwardConfig != nil {
				m.ForwardConfig = &model.ForwardConfig{
					ForwardRule:   result.ForwardConfig.ForwardRule,
					MatchRule:     result.ForwardConfig.MatchRule,
					TargetModel:   result.ForwardConfig.TargetModel,
					DecisionModel: result.ForwardConfig.DecisionModel,
					Keywords:      result.ForwardConfig.Keywords,
					TargetModels:  result.ForwardConfig.TargetModels,
				}
			}

			if params.Value == "all" {

				if m.ForwardConfig == nil {
					m.ForwardConfig = new(model.ForwardConfig)
				}

				m.IsForward = true
				m.ForwardConfig.ForwardRule = 1
				m.ForwardConfig.TargetModel = params.TargetModel

			} else {
				m.IsForward = gconv.Bool(params.Value)
				if m.IsForward && (m.ForwardConfig == nil ||
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
