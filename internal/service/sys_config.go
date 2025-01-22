// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
)

type (
	ISysConfig interface {
		// 更新配置
		Update(ctx context.Context, params model.SysConfigUpdateReq) error
		// 更改配置状态
		ChangeStatus(ctx context.Context, params model.SysConfigChangeStatusReq) error
		// 配置详情
		Detail(ctx context.Context) (*model.SysConfig, error)
		// 重置配置
		Reset(ctx context.Context, params model.SysConfigResetReq) (*entity.SysConfig, error)
		// 默认配置
		Default(ctx context.Context) (*entity.SysConfig, error)
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
