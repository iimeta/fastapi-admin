// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

type (
	IUser interface {
		// 用户更新信息
		UpdateInfo(ctx context.Context, params model.UserUpdateInfoReq) error
		// 用户更改密码
		ChangePassword(ctx context.Context, params model.UserChangePasswordReq) (err error)
		// 用户更改邮箱
		ChangeEmail(ctx context.Context, params model.UserChangeEmailReq) error
		// 用户更改头像
		ChangeAvatar(ctx context.Context, file *ghttp.UploadFile) error
		// 根据userId获取用户信息
		GetUserByUserId(ctx context.Context, userId int) (*model.User, error)
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
