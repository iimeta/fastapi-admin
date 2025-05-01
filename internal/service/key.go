// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/common"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
)

type (
	IKey interface {
		// 新建密钥
		Create(ctx context.Context, params model.KeyCreateReq, isModelAgent bool) error
		// 更新密钥
		Update(ctx context.Context, params model.KeyUpdateReq, isModelAgent bool) error
		// 更改密钥状态
		ChangeStatus(ctx context.Context, params model.KeyChangeStatusReq) error
		// 删除密钥
		Delete(ctx context.Context, id string) error
		// 密钥详情
		Detail(ctx context.Context, id string) (*model.Key, error)
		// 密钥分页列表
		Page(ctx context.Context, params model.KeyPageReq) (*model.KeyPageRes, error)
		// 密钥列表
		List(ctx context.Context, params model.KeyListReq) ([]*model.Key, error)
		// 密钥批量操作
		BatchOperate(ctx context.Context, params model.KeyBatchOperateReq) error
		// 根据Keys查询密钥详情列表
		DetailListByKey(ctx context.Context, keys []string) ([]*entity.Key, error)
		// 密钥模型权限
		Models(ctx context.Context, params model.KeyModelsReq) error
		// 密钥绑定分组
		Group(ctx context.Context, params model.KeyGroupReq) error
		// 检查任务
		CheckTask(ctx context.Context, enableError common.EnableError)
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
