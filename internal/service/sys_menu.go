// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/internal/model"
)

type (
	ISysMenu interface {
		// 新建应用
		Create(ctx context.Context, params model.SysMenuCreateReq) error
		// 更新应用
		Update(ctx context.Context, params model.SysMenuUpdateReq) error
		// 删除应用
		Delete(ctx context.Context, id string) error
		// 应用详情
		Detail(ctx context.Context, id string) (*model.SysMenu, error)
		// 应用分页列表
		Page(ctx context.Context, params model.SysMenuPageReq) (*model.SysMenuPageRes, error)
	}
)

var (
	localSysMenu ISysMenu
)

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}
