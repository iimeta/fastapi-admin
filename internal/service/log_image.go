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
	ILogImage interface {
		// 绘图日志详情
		Detail(ctx context.Context, id string) (*model.LogImage, error)
		// 绘图日志分页列表
		Page(ctx context.Context, params model.LogImagePageReq) (*model.LogImagePageRes, error)
		// 绘图日志详情复制字段值
		CopyField(ctx context.Context, params model.LogImageCopyFieldReq) (string, error)
	}
)

var (
	localLogImage ILogImage
)

func LogImage() ILogImage {
	if localLogImage == nil {
		panic("implement not found for interface ILogImage, forgot register?")
	}
	return localLogImage
}

func RegisterLogImage(i ILogImage) {
	localLogImage = i
}
