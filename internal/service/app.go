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
	IApp interface {
		// 新建应用
		Create(ctx context.Context, params model.AppCreateReq) (string, error)
		// 更新应用
		Update(ctx context.Context, params model.AppUpdateReq) error
		// 更改应用状态
		ChangeStatus(ctx context.Context, params model.AppChangeStatusReq) error
		// 删除应用
		Delete(ctx context.Context, id string) error
		// 应用详情
		Detail(ctx context.Context, id string) (*model.App, error)
		// 应用分页列表
		Page(ctx context.Context, params model.AppPageReq) (*model.AppPageRes, error)
		// 应用列表
		List(ctx context.Context, params model.AppListReq) ([]*model.App, error)
		// 新建应用密钥
		CreateKey(ctx context.Context, params model.AppCreateKeyReq) (string, error)
		// 应用密钥配置
		KeyConfig(ctx context.Context, params model.AppKeyConfigReq) (k string, err error)
		// 应用模型权限
		Models(ctx context.Context, params model.AppModelsReq) error
		// 应用绑定分组
		Group(ctx context.Context, params model.AppGroupReq) error
		// 应用批量操作
		BatchOperate(ctx context.Context, params model.AppBatchOperateReq) error
		// 应用密钥批量操作
		KeyBatchOperate(ctx context.Context, params model.AppKeyBatchOperateReq) (keys string, err error)
		// 应用密钥导出
		KeyExport(ctx context.Context, params model.AppKeyExportReq) (string, error)
	}
)

var (
	localApp IApp
)

func App() IApp {
	if localApp == nil {
		panic("implement not found for interface IApp, forgot register?")
	}
	return localApp
}

func RegisterApp(i IApp) {
	localApp = i
}
