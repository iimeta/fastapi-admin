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
	ISysConfig interface {
		// 新建应用
		Create(ctx context.Context, params model.SysConfigCreateReq) error
		// 更新应用
		Update(ctx context.Context, params model.SysConfigUpdateReq) error
		// 删除应用
		Delete(ctx context.Context, id string) error
		// 应用详情
		Detail(ctx context.Context, id string) (*model.SysConfig, error)
		// 应用分页列表
		Page(ctx context.Context, params model.SysConfigPageReq) (*model.SysConfigPageRes, error)
	}
)

var (
	localSysConfig ISysConfig
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}
