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

type sSysRole struct{}

func init() {
	service.RegisterSysRole(New())
}

func New() service.ISysRole {
	return &sSysRole{}
}

// 新建应用
func (s *sSysRole) Create(ctx context.Context, params model.SysRoleCreateReq) error {

	if _, err := dao.SysRole.Insert(ctx, &do.SysRole{
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
func (s *sSysRole) Update(ctx context.Context, params model.SysRoleUpdateReq) error {

	if err := dao.SysRole.UpdateById(ctx, params.Id, &do.SysRole{
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
func (s *sSysRole) Delete(ctx context.Context, id string) error {

	if err := dao.SysRole.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用详情
func (s *sSysRole) Detail(ctx context.Context, id string) (*model.SysRole, error) {

	app, err := dao.SysRole.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.SysRole{
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
func (s *sSysRole) Page(ctx context.Context, params model.SysRolePageReq) (*model.SysRolePageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	results, err := dao.SysRole.FindByPage(ctx, paging, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.SysRole, 0)
	for _, result := range results {
		items = append(items, &model.SysRole{
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

	return &model.SysRolePageRes{
		Items: items,
		Paging: &model.Paging{
			Page:      paging.Page,
			PageSize:  paging.PageSize,
			Total:     paging.Total,
			PageCount: paging.PageCount,
		},
	}, nil
}
