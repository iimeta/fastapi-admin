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
	ITaskVideo interface {
		// 视频任务分页列表
		Page(ctx context.Context, params model.TaskVideoPageReq) (*model.TaskVideoPageRes, error)
	}
)

var (
	localTaskVideo ITaskVideo
)

func TaskVideo() ITaskVideo {
	if localTaskVideo == nil {
		panic("implement not found for interface ITaskVideo, forgot register?")
	}
	return localTaskVideo
}

func RegisterTaskVideo(i ITaskVideo) {
	localTaskVideo = i
}
