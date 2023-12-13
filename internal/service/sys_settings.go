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
	ISysSettings interface {
		// 新建应用
		Create(ctx context.Context, params model.SysSettingsCreateReq) error
		// 更新应用
		Update(ctx context.Context, params model.SysSettingsUpdateReq) error
		// 删除应用
		Delete(ctx context.Context, id string) error
		// 应用详情
		Detail(ctx context.Context, id string) (*model.SysSettings, error)
		// 应用分页列表
		Page(ctx context.Context, params model.SysSettingsPageReq) (*model.SysSettingsPageRes, error)
	}
)

var (
	localSysSettings ISysSettings
)

func SysSettings() ISysSettings {
	if localSysSettings == nil {
		panic("implement not found for interface ISysSettings, forgot register?")
	}
	return localSysSettings
}

func RegisterSysSettings(i ISysSettings) {
	localSysSettings = i
}
