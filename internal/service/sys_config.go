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
		// 更新配置
		Update(ctx context.Context, params model.SysConfigUpdateReq) error
		// 更改配置状态
		ChangeStatus(ctx context.Context, params model.SysConfigChangeStatusReq) error
		// 配置详情
		Detail(ctx context.Context) (*model.SysConfig, error)
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
