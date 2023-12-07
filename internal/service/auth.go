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
		// 注册接口
		Register(ctx context.Context, params model.RegisterReq) error
		// 登录接口
		Login(ctx context.Context, params model.LoginReq) (res *model.LoginRes, err error)
		// 退出登录接口
		Logout(ctx context.Context) error
		// 账号找回接口
		Forget(ctx context.Context, params model.ForgetReq) error
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
