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
	IImage interface {
		// 绘图日志详情
		Detail(ctx context.Context, id string) (*model.Image, error)
		// 绘图日志分页列表
		Page(ctx context.Context, params model.ImagePageReq) (*model.ImagePageRes, error)
	}
)

var (
	localImage IImage
)

func Image() IImage {
	if localImage == nil {
		panic("implement not found for interface IImage, forgot register?")
	}
	return localImage
}

func RegisterImage(i IImage) {
	localImage = i
}
