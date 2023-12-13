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

type sSysMenu struct{}

func init() {
	service.RegisterSysMenu(New())
}

func New() service.ISysMenu {
	return &sSysMenu{}
}

// 新建应用
func (s *sSysMenu) Create(ctx context.Context, params model.SysMenuCreateReq) error {

	if _, err := dao.SysMenu.Insert(ctx, &do.SysMenu{
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
func (s *sSysMenu) Update(ctx context.Context, params model.SysMenuUpdateReq) error {

	if err := dao.SysMenu.UpdateById(ctx, params.Id, &do.SysMenu{
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
func (s *sSysMenu) Delete(ctx context.Context, id string) error {

	if err := dao.SysMenu.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用详情
func (s *sSysMenu) Detail(ctx context.Context, id string) (*model.SysMenu, error) {

	app, err := dao.SysMenu.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.SysMenu{
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
func (s *sSysMenu) Page(ctx context.Context, params model.SysMenuPageReq) (*model.SysMenuPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	results, err := dao.SysMenu.FindByPage(ctx, paging, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.SysMenu, 0)
	for _, result := range results {
		items = append(items, &model.SysMenu{
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

	return &model.SysMenuPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:      paging.Page,
			PageSize:  paging.PageSize,
			Total:     paging.Total,
			PageCount: paging.PageCount,
		},
	}, nil
}
