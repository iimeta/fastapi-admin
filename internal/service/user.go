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
		// 修改用户信息
		ChangeDetail(ctx context.Context, params model.UserDetailUpdateReq) error
		// 修改密码接口
		ChangePassword(ctx context.Context, params model.UserPasswordUpdateReq) (err error)
		// 用户设置
		Setting(ctx context.Context) (*model.UserSettingRes, error)
		// 换绑手机号
		ChangeMobile(ctx context.Context, params model.UserMobileUpdateReq) error
		// 换绑邮箱
		ChangeEmail(ctx context.Context, params model.UserEmailUpdateReq) error
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
