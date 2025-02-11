// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/iimeta/fastapi-admin/internal/model"
)

type (
	ISysAdmin interface {
		// 管理员更新信息
		UpdateInfo(ctx context.Context, params model.UserUpdateInfoReq) error
		// 管理员更改密码
		ChangePassword(ctx context.Context, params model.UserChangePasswordReq) (err error)
		// 管理员更改邮箱
		ChangeEmail(ctx context.Context, params model.UserChangeEmailReq) error
		// 管理员更改头像
		ChangeAvatar(ctx context.Context, file *ghttp.UploadFile) error
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
