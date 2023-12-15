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
		SaveUser(ctx context.Context, user *model.User) error
		// 保存管理员会话信息
		SaveAdmin(ctx context.Context, admin *model.SysAdmin) error
		// 获取会话中用户主键ID
		GetUid(ctx context.Context) string
		// 获取会话中UserId
		GetUserId(ctx context.Context) int
		// 获取会话中用户信息
		GetUser(ctx context.Context) *model.User
		// 获取会话中用户信息
		GetAdmin(ctx context.Context) *model.SysAdmin
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
