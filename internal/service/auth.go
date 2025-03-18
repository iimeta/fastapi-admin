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
	IAuth interface {
		// 注册
		Register(ctx context.Context, params model.RegisterReq, channel ...string) error
		// 登录
		Login(ctx context.Context, params model.LoginReq) (res *model.LoginRes, err error)
		// 退出登录
		Logout(ctx context.Context) error
		// 找回密码
		Forget(ctx context.Context, params model.ForgetReq) error
		// 生成用户Token
		GenUserToken(ctx context.Context, user *model.User, isSaveSession bool) (token string, err error)
		// 根据Token获取用户信息
		GetUserByToken(ctx context.Context, token string) (*model.User, error)
		// 根据Token更新用户信息
		UpdateUserByToken(ctx context.Context, token string, user *model.User) error
		// 生成管理员Token
		GenAdminToken(ctx context.Context, admin *model.SysAdmin, isSaveSession bool) (token string, err error)
		// 根据Token获取管理员信息
		GetAdminByToken(ctx context.Context, token string) (*model.SysAdmin, error)
		// 根据Token更新管理员信息
		UpdateAdminByToken(ctx context.Context, token string, admin *model.SysAdmin) error
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
