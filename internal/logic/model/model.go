package model

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
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

	if _, err := dao.Model.Insert(ctx, &do.Model{
		Corp:               params.Corp,
		Name:               gstr.Trim(params.Name),
		Model:              gstr.Trim(params.Model),
		Type:               params.Type,
		PromptRatio:        params.PromptRatio,
		CompletionRatio:    params.CompletionRatio,
		DataFormat:         params.DataFormat,
		IsEnableModelAgent: params.IsEnableModelAgent,
		ModelAgents:        params.ModelAgents,
		IsPublic:           params.IsPublic,
		Remark:             params.Remark,
		Status:             params.Status,
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

	if err := dao.Model.UpdateById(ctx, params.Id, &do.Model{
		Corp:               params.Corp,
		Name:               gstr.Trim(params.Name),
		Model:              gstr.Trim(params.Model),
		Type:               params.Type,
		PromptRatio:        params.PromptRatio,
		CompletionRatio:    params.CompletionRatio,
		DataFormat:         params.DataFormat,
		IsEnableModelAgent: params.IsEnableModelAgent,
		ModelAgents:        params.ModelAgents,
		IsPublic:           params.IsPublic,
		Remark:             params.Remark,
		Status:             params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除模型
func (s *sModel) Delete(ctx context.Context, id string) error {

	if _, err := dao.Model.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := dao.Key.UpdateMany(ctx, bson.M{"models": bson.M{"$in": []string{id}}}, bson.M{
		"$pull": bson.M{
			"models": id,
		},
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

	return &model.Model{
		Id:                 m.Id,
		Corp:               m.Corp,
		Name:               m.Name,
		Model:              m.Model,
		Type:               m.Type,
		PromptRatio:        m.PromptRatio,
		CompletionRatio:    m.CompletionRatio,
		DataFormat:         m.DataFormat,
		IsEnableModelAgent: m.IsEnableModelAgent,
		ModelAgents:        m.ModelAgents,
		ModelAgentNames:    modelAgentNames,
		IsPublic:           m.IsPublic,
		Remark:             m.Remark,
		Status:             m.Status,
		Creator:            m.Creator,
		Updater:            m.Updater,
		CreatedAt:          util.FormatDatetime(m.CreatedAt),
		UpdatedAt:          util.FormatDatetime(m.UpdatedAt),
	}, nil
}

// 模型分页列表
func (s *sModel) Page(ctx context.Context, params model.ModelPageReq) (*model.ModelPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.Corp != "" {
		filter["corp"] = params.Corp
	}

	if params.Name != "" {
		filter["name"] = params.Name
	}

	if params.Model != "" {
		filter["model"] = params.Model
	}

	if params.Type != 0 {
		filter["type"] = params.Type
	}

	results, err := dao.Model.FindByPage(ctx, paging, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Model, 0)
	for _, result := range results {
		items = append(items, &model.Model{
			Id:                 result.Id,
			Corp:               result.Corp,
			Name:               result.Name,
			Model:              result.Model,
			Type:               result.Type,
			PromptRatio:        result.PromptRatio,
			CompletionRatio:    result.CompletionRatio,
			DataFormat:         result.DataFormat,
			IsEnableModelAgent: result.IsEnableModelAgent,
			ModelAgents:        result.ModelAgents,
			IsPublic:           result.IsPublic,
			Remark:             result.Remark,
			Status:             result.Status,
			Creator:            result.Creator,
			Updater:            result.Updater,
			CreatedAt:          util.FormatDatetime(result.CreatedAt),
			UpdatedAt:          util.FormatDatetime(result.UpdatedAt),
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

	results, err := dao.Model.Find(ctx, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Model, 0)
	for _, result := range results {
		items = append(items, &model.Model{
			Id:                 result.Id,
			Corp:               result.Corp,
			Name:               result.Name,
			Model:              result.Model,
			Type:               result.Type,
			PromptRatio:        result.PromptRatio,
			CompletionRatio:    result.CompletionRatio,
			DataFormat:         result.DataFormat,
			IsEnableModelAgent: result.IsEnableModelAgent,
			ModelAgents:        result.ModelAgents,
			Remark:             result.Remark,
			Status:             result.Status,
			Creator:            result.Creator,
			Updater:            result.Updater,
			CreatedAt:          util.FormatDatetime(result.CreatedAt),
			UpdatedAt:          util.FormatDatetime(result.UpdatedAt),
		})
	}

	return items, nil
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
