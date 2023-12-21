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
	IKey interface {
		// 新建密钥
		Create(ctx context.Context, params model.KeyCreateReq) error
		// 更新密钥
		Update(ctx context.Context, params model.KeyUpdateReq) error
		// 删除密钥
		Delete(ctx context.Context, id string) error
		// 密钥详情
		Detail(ctx context.Context, id string) (*model.Key, error)
		// 密钥分页列表
		Page(ctx context.Context, params model.KeyPageReq) (*model.KeyPageRes, error)
		// 密钥列表
		List(ctx context.Context, params model.KeyListReq) ([]*model.Key, error)
	}
)

var (
	localKey IKey
)

func Key() IKey {
	if localKey == nil {
		panic("implement not found for interface IKey, forgot register?")
	}
	return localKey
}

func RegisterKey(i IKey) {
	localKey = i
}
