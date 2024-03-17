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
	ISysAdmin interface {
		// 新建管理员
		Create(ctx context.Context, params model.SysAdminCreateReq) error
		// 更新管理员
		Update(ctx context.Context, params model.SysAdminUpdateReq) error
		// 删除管理员
		Delete(ctx context.Context, id string) error
		// 管理员详情
		Detail(ctx context.Context, id string) (*model.SysAdmin, error)
		// 管理员分页列表
		Page(ctx context.Context, params model.SysAdminPageReq) (*model.SysAdminPageRes, error)
	}
)

var (
	localSysAdmin ISysAdmin
)

func SysAdmin() ISysAdmin {
	if localSysAdmin == nil {
		panic("implement not found for interface ISysAdmin, forgot register?")
	}
	return localSysAdmin
}

func RegisterSysAdmin(i ISysAdmin) {
	localSysAdmin = i
}
