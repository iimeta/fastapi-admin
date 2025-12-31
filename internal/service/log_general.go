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
	ILogGeneral interface {
		// 通用日志详情
		Detail(ctx context.Context, id string) (*model.LogGeneral, error)
		// 通用日志分页列表
		Page(ctx context.Context, params model.LogGeneralPageReq) (*model.LogGeneralPageRes, error)
		// 通用日志详情复制字段值
		CopyField(ctx context.Context, params model.LogGeneralCopyFieldReq) (string, error)
	}
)

var (
	localLogGeneral ILogGeneral
)

func LogGeneral() ILogGeneral {
	if localLogGeneral == nil {
		panic("implement not found for interface ILogGeneral, forgot register?")
	}
	return localLogGeneral
}

func RegisterLogGeneral(i ILogGeneral) {
	localLogGeneral = i
}
