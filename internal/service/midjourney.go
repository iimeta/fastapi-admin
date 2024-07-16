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
	IMidjourney interface {
		// Midjourney详情
		Detail(ctx context.Context, id string) (*model.Midjourney, error)
		// Midjourney分页列表
		Page(ctx context.Context, params model.MidjourneyPageReq) (*model.MidjourneyPageRes, error)
	}
)

var (
	localMidjourney IMidjourney
)

func Midjourney() IMidjourney {
	if localMidjourney == nil {
		panic("implement not found for interface IMidjourney, forgot register?")
	}
	return localMidjourney
}

func RegisterMidjourney(i IMidjourney) {
	localMidjourney = i
}
