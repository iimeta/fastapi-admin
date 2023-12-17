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

// 新建菜单
func (s *sSysMenu) Create(ctx context.Context, params model.SysMenuCreateReq) error {

	if _, err := dao.SysMenu.Insert(ctx, &do.SysMenu{
		Pid:    params.Pid,
		Name:   params.Name,
		Perm:   params.Perm,
		Type:   params.Type,
		Route:  params.Route,
		Sort:   params.Sort,
		Level:  params.Level,
		Hidden: params.Hidden,
		Remark: params.Remark,
		Status: params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新菜单
func (s *sSysMenu) Update(ctx context.Context, params model.SysMenuUpdateReq) error {

	if err := dao.SysMenu.UpdateById(ctx, params.Id, &do.SysMenu{
		Pid:    params.Pid,
		Name:   params.Name,
		Perm:   params.Perm,
		Type:   params.Type,
		Route:  params.Route,
		Sort:   params.Sort,
		Level:  params.Level,
		Hidden: params.Hidden,
		Remark: params.Remark,
		Status: params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除菜单
func (s *sSysMenu) Delete(ctx context.Context, id string) error {

	if err := dao.SysMenu.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 菜单详情
func (s *sSysMenu) Detail(ctx context.Context, id string) (*model.SysMenu, error) {

	menu, err := dao.SysMenu.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.SysMenu{
		Id:        menu.Id,
		Pid:       menu.Pid,
		Name:      menu.Name,
		Perm:      menu.Perm,
		Type:      menu.Type,
		Route:     menu.Route,
		Sort:      menu.Sort,
		Level:     menu.Level,
		Hidden:    menu.Hidden,
		Remark:    menu.Remark,
		Status:    menu.Status,
		Creator:   menu.Creator,
		Updater:   menu.Updater,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}, nil
}

// 菜单分页列表
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
			Id:        result.Id,
			Pid:       result.Pid,
			Name:      result.Name,
			Perm:      result.Perm,
			Type:      result.Type,
			Route:     result.Route,
			Sort:      result.Sort,
			Level:     result.Level,
			Hidden:    result.Hidden,
			Remark:    result.Remark,
			Status:    result.Status,
			Creator:   result.Creator,
			Updater:   result.Updater,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		})
	}

	return &model.SysMenuPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}
