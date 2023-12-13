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
	ISysRole interface {
		// 新建应用
		Create(ctx context.Context, params model.SysRoleCreateReq) error
		// 更新应用
		Update(ctx context.Context, params model.SysRoleUpdateReq) error
		// 删除应用
		Delete(ctx context.Context, id string) error
		// 应用详情
		Detail(ctx context.Context, id string) (*model.SysRole, error)
		// 应用分页列表
		Page(ctx context.Context, params model.SysRolePageReq) (*model.SysRolePageRes, error)
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
