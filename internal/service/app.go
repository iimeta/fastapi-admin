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
		Create(ctx context.Context, params model.AppCreateReq) error
		// 更新应用
		Update(ctx context.Context, params model.AppUpdateReq) error
		// 删除应用
		Delete(ctx context.Context, id string) error
		// 应用详情
		Detail(ctx context.Context, id string) (*model.App, error)
		// 应用分页列表
		Page(ctx context.Context, params model.AppPageReq) (*model.AppPageRes, error)
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
