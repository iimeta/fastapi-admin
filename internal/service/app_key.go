// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

type (
	IAppKey interface {
		// 新建应用密钥
		Create(ctx context.Context, params model.AppKeyCreateReq) (string, error)
		// 应用密钥配置
		Config(ctx context.Context, params model.AppKeyConfigReq) (k string, err error)
		// 更改应用密钥状态
		ChangeStatus(ctx context.Context, params model.AppKeyChangeStatusReq) error
		// 删除应用密钥
		Delete(ctx context.Context, id string) error
		// 应用密钥详情
		Detail(ctx context.Context, id string) (*model.AppKey, error)
		// 应用密钥分页列表
		Page(ctx context.Context, params model.AppKeyPageReq) (*model.AppKeyPageRes, error)
		// 应用密钥模型权限
		Models(ctx context.Context, params model.AppKeyModelsReq) error
		// 应用密钥绑定分组
		Group(ctx context.Context, params model.AppKeyGroupReq) error
		// 应用密钥批量操作
		BatchOperate(ctx context.Context, params model.AppKeyBatchOperateReq) (keys string, err error)
		// 应用密钥导出
		Export(ctx context.Context, params model.AppKeyExportReq) (string, error)
	}
)

var (
	localAppKey IAppKey
)

func AppKey() IAppKey {
	if localAppKey == nil {
		panic("implement not found for interface IAppKey, forgot register?")
	}
	return localAppKey
}

func RegisterAppKey(i IAppKey) {
	localAppKey = i
}
