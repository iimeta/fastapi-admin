package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
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

	if _, err := dao.Model.Insert(ctx, &do.Model{
		Corp:    params.Corp,
		Name:    params.Name,
		Type:    params.Type,
		BaseUrl: params.BaseUrl,
		Path:    params.Path,
		Proxy:   params.Proxy,
		Keys:    params.Keys,
		Status:  params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新模型
func (s *sModel) Update(ctx context.Context, params model.ModelUpdateReq) error {

	if err := dao.Model.UpdateById(ctx, params.Id, &do.Model{
		Corp:    params.Corp,
		Name:    params.Name,
		Type:    params.Type,
		BaseUrl: params.BaseUrl,
		Path:    params.Path,
		Proxy:   params.Proxy,
		Keys:    params.Keys,
		Status:  params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除模型
func (s *sModel) Delete(ctx context.Context, id string) error {

	if err := dao.Model.DeleteById(ctx, id); err != nil {
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

	return &model.Model{
		Id:        m.Id,
		Corp:      m.Corp,
		Name:      m.Name,
		Type:      m.Type,
		BaseUrl:   m.BaseUrl,
		Path:      m.Path,
		Proxy:     m.Proxy,
		Keys:      m.Keys,
		Status:    m.Status,
		Creator:   m.Creator,
		Updater:   m.Updater,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}
