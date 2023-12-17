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

type sSysSettings struct{}

func init() {
	service.RegisterSysSettings(New())
}

func New() service.ISysSettings {
	return &sSysSettings{}
}

// 新建应用
func (s *sSysSettings) Create(ctx context.Context, params model.SysSettingsCreateReq) error {

	if _, err := dao.SysSettings.Insert(ctx, &do.SysSettings{
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
func (s *sSysSettings) Update(ctx context.Context, params model.SysSettingsUpdateReq) error {

	if err := dao.SysSettings.UpdateById(ctx, params.Id, &do.SysSettings{
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
func (s *sSysSettings) Delete(ctx context.Context, id string) error {

	if err := dao.SysSettings.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用详情
func (s *sSysSettings) Detail(ctx context.Context, id string) (*model.SysSettings, error) {

	app, err := dao.SysSettings.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.SysSettings{
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
func (s *sSysSettings) Page(ctx context.Context, params model.SysSettingsPageReq) (*model.SysSettingsPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	results, err := dao.SysSettings.FindByPage(ctx, paging, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.SysSettings, 0)
	for _, result := range results {
		items = append(items, &model.SysSettings{
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

	return &model.SysSettingsPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}
