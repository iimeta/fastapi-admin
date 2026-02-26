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
	IOpen interface {
		// 站点配置
		Site(ctx context.Context, params model.SiteConfigDetailReq) *model.SiteConfig
		// 系统配置
		Config(ctx context.Context, params model.SysConfigReq) (*model.SysConfig, error)
		// 视频文件
		Video(ctx context.Context, fileName string) (string, error)
		// 文件
		File(ctx context.Context, fileName string) (string, error)
		// 用户协议
		UserAgreement(ctx context.Context, params model.SysConfigReq) (string, error)
		// 隐私政策
		PrivacyPolicy(ctx context.Context, params model.SysConfigReq) (string, error)
	}
)

var (
	localOpen IOpen
)

func Open() IOpen {
	if localOpen == nil {
		panic("implement not found for interface IOpen, forgot register?")
	}
	return localOpen
}

func RegisterOpen(i IOpen) {
	localOpen = i
}
