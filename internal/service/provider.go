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
	IProvider interface {
		// 新建提供商
		Create(ctx context.Context, params model.ProviderCreateReq) (string, error)
		// 更新提供商
		Update(ctx context.Context, params model.ProviderUpdateReq) error
		// 更改提供商公开状态
		ChangePublic(ctx context.Context, params model.ProviderChangePublicReq) error
		// 更改提供商状态
		ChangeStatus(ctx context.Context, params model.ProviderChangeStatusReq) error
		// 删除提供商
		Delete(ctx context.Context, id string) error
		// 提供商详情
		Detail(ctx context.Context, id string) (*model.Provider, error)
		// 提供商分页列表
		Page(ctx context.Context, params model.ProviderPageReq) (*model.ProviderPageRes, error)
		// 提供商列表
		List(ctx context.Context, params model.ProviderListReq) ([]*model.Provider, error)
		// 提供商批量操作
		BatchOperate(ctx context.Context, params model.ProviderBatchOperateReq) error
	}
)

var (
	localProvider IProvider
)

func Provider() IProvider {
	if localProvider == nil {
		panic("implement not found for interface IProvider, forgot register?")
	}
	return localProvider
}

func RegisterProvider(i IProvider) {
	localProvider = i
}
