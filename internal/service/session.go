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
		// 保存会话信息
		Save(ctx context.Context, user *model.User) error
		// 获取会话中用户主键ID
		GetUid(ctx context.Context) string
		// 获取会话中UserId
		GetUserId(ctx context.Context) int
		// 获取会话中用户信息
		GetUser(ctx context.Context) *model.User
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
