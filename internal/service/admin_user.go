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
	IAdminUser interface {
		// 新建用户
		Create(ctx context.Context, params model.UserCreateReq) error
		// 更新用户
		Update(ctx context.Context, params model.UserUpdateReq) error
		// 更改用户额度过期时间
		ChangeQuotaExpire(ctx context.Context, params model.UserChangeQuotaExpireReq) error
		// 更改用户状态
		ChangeStatus(ctx context.Context, params model.UserChangeStatusReq) error
		// 删除用户
		Delete(ctx context.Context, id string) error
		// 用户详情
		Detail(ctx context.Context, id string) (*model.User, error)
		// 用户分页列表
		Page(ctx context.Context, params model.UserPageReq) (*model.UserPageRes, error)
		// 用户列表
		List(ctx context.Context, params model.UserListReq) ([]*model.User, error)
		// 用户充值
		Recharge(ctx context.Context, params model.UserRechargeReq) error
		// 用户模型权限
		Models(ctx context.Context, params model.UserModelsReq) error
		// 用户批量操作
		BatchOperate(ctx context.Context, params model.UserBatchOperateReq) error
	}
)

var (
	localAdminUser IAdminUser
)

func AdminUser() IAdminUser {
	if localAdminUser == nil {
		panic("implement not found for interface IAdminUser, forgot register?")
	}
	return localAdminUser
}

func RegisterAdminUser(i IAdminUser) {
	localAdminUser = i
}
