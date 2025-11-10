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
	ILogMidjourney interface {
		// Midjourney日志详情
		Detail(ctx context.Context, id string) (*model.LogMidjourney, error)
		// Midjourney日志分页列表
		Page(ctx context.Context, params model.LogMidjourneyPageReq) (*model.LogMidjourneyPageRes, error)
	}
)

var (
	localLogMidjourney ILogMidjourney
)

func LogMidjourney() ILogMidjourney {
	if localLogMidjourney == nil {
		panic("implement not found for interface ILogMidjourney, forgot register?")
	}
	return localLogMidjourney
}

func RegisterLogMidjourney(i ILogMidjourney) {
	localLogMidjourney = i
}
