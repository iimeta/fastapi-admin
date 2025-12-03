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
	ILogVideo interface {
		// 视频日志详情
		Detail(ctx context.Context, id string) (*model.LogVideo, error)
		// 视频日志分页列表
		Page(ctx context.Context, params model.LogVideoPageReq) (*model.LogVideoPageRes, error)
		// 视频日志详情复制字段值
		CopyField(ctx context.Context, params model.LogVideoCopyFieldReq) (string, error)
	}
)

var (
	localLogVideo ILogVideo
)

func LogVideo() ILogVideo {
	if localLogVideo == nil {
		panic("implement not found for interface ILogVideo, forgot register?")
	}
	return localLogVideo
}

func RegisterLogVideo(i ILogVideo) {
	localLogVideo = i
}
