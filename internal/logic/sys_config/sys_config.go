package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
)

type sSysConfig struct{}

func init() {
	service.RegisterSysConfig(New())
}

func New() service.ISysConfig {
	return &sSysConfig{}
}

// 新建应用
func (s *sSysConfig) Create(ctx context.Context, params model.SysConfigCreateReq) error {

	if _, err := dao.SysConfig.Insert(ctx, &do.SysConfig{
		Name:        params.Name,
		Type:        params.Type,
		Models:      params.Models,
		Keys:        params.Keys,
		IpWhitelist: params.IpWhitelist,
		IpBlacklist: params.IpBlacklist,
		Remark:      params.Remark,
		Status:      params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新应用
func (s *sSysConfig) Update(ctx context.Context, params model.SysConfigUpdateReq) error {

	if err := dao.SysConfig.UpdateById(ctx, params.Id, &do.SysConfig{
		Name:        params.Name,
		Type:        params.Type,
		Models:      params.Models,
		Keys:        params.Keys,
		IpWhitelist: params.IpWhitelist,
		IpBlacklist: params.IpBlacklist,
		Remark:      params.Remark,
		Status:      params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除应用
func (s *sSysConfig) Delete(ctx context.Context, id string) error {

	if err := dao.SysConfig.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用详情
func (s *sSysConfig) Detail(ctx context.Context, id string) (*model.SysConfig, error) {

	app, err := dao.SysConfig.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.SysConfig{
		Id:          app.Id,
		Name:        app.Name,
		Type:        app.Type,
		Models:      app.Models,
		Keys:        app.Keys,
		IpWhitelist: app.IpWhitelist,
		IpBlacklist: app.IpBlacklist,
		Remark:      app.Remark,
		Status:      app.Status,
		Creator:     app.Creator,
		Updater:     app.Updater,
		CreatedAt:   app.CreatedAt,
		UpdatedAt:   app.UpdatedAt,
	}, nil
}

// 应用分页列表
func (s *sSysConfig) Page(ctx context.Context, params model.SysConfigPageReq) (*model.SysConfigPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	results, err := dao.SysConfig.FindByPage(ctx, paging, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.SysConfig, 0)
	for _, result := range results {
		items = append(items, &model.SysConfig{
			Id:          result.Id,
			Name:        result.Name,
			Type:        result.Type,
			Models:      result.Models,
			Keys:        result.Keys,
			IpWhitelist: result.IpWhitelist,
			IpBlacklist: result.IpBlacklist,
			Remark:      result.Remark,
			Status:      result.Status,
			Creator:     result.Creator,
			Updater:     result.Updater,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		})
	}

	return &model.SysConfigPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:      paging.Page,
			PageSize:  paging.PageSize,
			Total:     paging.Total,
			PageCount: paging.PageCount,
		},
	}, nil
}
