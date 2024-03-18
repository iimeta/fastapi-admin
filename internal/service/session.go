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
	ISession interface {
		// 保存用户会话信息
		SaveUser(ctx context.Context, token string, user *model.User) error
		// 保存管理员会话信息
		SaveAdmin(ctx context.Context, token string, admin *model.SysAdmin) error
		// 获取会话中Token
		GetToken(ctx context.Context) string
		// 获取会话中用户主键ID
		GetUid(ctx context.Context) string
		// 获取会话中UserId
		GetUserId(ctx context.Context) int
		// 获取会话中角色
		GetRole(ctx context.Context) string
		// 获取会话中创建人
		GetCreator(ctx context.Context) string
		// 获取会话中用户信息
		GetUser(ctx context.Context) *model.User
		// 获取会话中管理员信息
		GetAdmin(ctx context.Context) *model.SysAdmin
		// 判断获取会话中角色是否为用户
		IsUserRole(ctx context.Context) bool
		// 判断获取会话中角色是否为管理员
		IsAdminRole(ctx context.Context) bool
		// 更新用户会话信息
		UpdateUserSession(ctx context.Context, user *model.User) error
		// 更新管理员会话信息
		UpdateAdminSession(ctx context.Context, admin *model.SysAdmin) error
	}
)

var (
	localSession ISession
)

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
