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
	IChat interface {
		// 聊天详情
		Detail(ctx context.Context, id string) (*model.Chat, error)
		// 聊天分页列表
		Page(ctx context.Context, params model.ChatPageReq) (*model.ChatPageRes, error)
	}
)

var (
	localChat IChat
)

func Chat() IChat {
	if localChat == nil {
		panic("implement not found for interface IChat, forgot register?")
	}
	return localChat
}

func RegisterChat(i IChat) {
	localChat = i
}
