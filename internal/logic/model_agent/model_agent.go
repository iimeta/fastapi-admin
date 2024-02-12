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

type sModelAgent struct{}

func init() {
	service.RegisterModelAgent(New())
}

func New() service.IModelAgent {
	return &sModelAgent{}
}

// 新建模型代理
func (s *sModelAgent) Create(ctx context.Context, params model.ModelAgentCreateReq) error {

	if s.IsNameExist(ctx, params.Name) {
		return errors.Newf("模型代理名称 \"%s\" 已存在", params.Name)
	}

	if _, err := dao.ModelAgent.Insert(ctx, &do.ModelAgent{
		Name:    gstr.Trim(params.Name),
		BaseUrl: params.BaseUrl,
		Path:    params.Path,
		Weight:  params.Weight,
		Remark:  params.Remark,
		Status:  params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新模型代理
func (s *sModelAgent) Update(ctx context.Context, params model.ModelAgentUpdateReq) error {

	if s.IsNameExist(ctx, params.Name, params.Id) {
		return errors.Newf("模型代理名称 \"%s\" 已存在", params.Name)
	}

	if err := dao.ModelAgent.UpdateById(ctx, params.Id, &do.ModelAgent{
		Name:    gstr.Trim(params.Name),
		BaseUrl: params.BaseUrl,
		Path:    params.Path,
		Weight:  params.Weight,
		Remark:  params.Remark,
		Status:  params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除模型代理
func (s *sModelAgent) Delete(ctx context.Context, id string) error {

	if _, err := dao.ModelAgent.DeleteById(ctx, id); err != nil {
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

	return &model.ModelAgent{
		Id:        modelAgent.Id,
		Name:      modelAgent.Name,
		BaseUrl:   modelAgent.BaseUrl,
		Path:      modelAgent.Path,
		Weight:    modelAgent.Weight,
		Remark:    modelAgent.Remark,
		Status:    modelAgent.Status,
		Creator:   modelAgent.Creator,
		Updater:   modelAgent.Updater,
		CreatedAt: util.FormatDatetime(modelAgent.CreatedAt),
		UpdatedAt: util.FormatDatetime(modelAgent.UpdatedAt),
	}, nil
}

// 模型代理分页列表
func (s *sModelAgent) Page(ctx context.Context, params model.ModelAgentPageReq) (*model.ModelAgentPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.Name != "" {
		filter["name"] = params.Name
	}

	if params.BaseUrl != "" {
		filter["base_url"] = params.BaseUrl
	}

	results, err := dao.ModelAgent.FindByPage(ctx, paging, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.ModelAgent, 0)
	for _, result := range results {
		items = append(items, &model.ModelAgent{
			Id:        result.Id,
			Name:      result.Name,
			BaseUrl:   result.BaseUrl,
			Path:      result.Path,
			Weight:    result.Weight,
			Remark:    result.Remark,
			Status:    result.Status,
			Creator:   result.Creator,
			Updater:   result.Updater,
			CreatedAt: util.FormatDatetime(result.CreatedAt),
			UpdatedAt: util.FormatDatetime(result.UpdatedAt),
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

	results, err := dao.ModelAgent.Find(ctx, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.ModelAgent, 0)
	for _, result := range results {
		items = append(items, &model.ModelAgent{
			Id:        result.Id,
			Name:      result.Name,
			BaseUrl:   result.BaseUrl,
			Path:      result.Path,
			Weight:    result.Weight,
			Remark:    result.Remark,
			Status:    result.Status,
			Creator:   result.Creator,
			Updater:   result.Updater,
			CreatedAt: util.FormatDatetime(result.CreatedAt),
			UpdatedAt: util.FormatDatetime(result.UpdatedAt),
		})
	}

	return items, nil
}

// 模型代理名称是否存在
func (s *sModelAgent) IsNameExist(ctx context.Context, name string, id ...string) bool {

	model, err := dao.ModelAgent.FindOne(ctx, bson.M{"name": gstr.Trim(name)})
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
