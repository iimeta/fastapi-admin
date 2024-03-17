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
	IUser interface {
		// 用户信息
		Info(ctx context.Context) (*model.UserInfoRes, error)
		// 用户更新信息
		UpdateInfo(ctx context.Context, params model.UserUpdateInfoReq) error
		// 用户修改密码接口
		ChangePassword(ctx context.Context, params model.UserChangePasswordReq) (err error)
		// 用户修改邮箱
		ChangeEmail(ctx context.Context, params model.UserChangeEmailReq) error
		// 根据userId获取用户信息
		GetUserById(ctx context.Context, userId int) (*model.User, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
