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
		// 视频任务详情
		Detail(ctx context.Context, id string) (*model.TaskVideo, error)
		// 视频任务分页列表
		Page(ctx context.Context, params model.TaskVideoPageReq) (*model.TaskVideoPageRes, error)
		// 视频任务详情复制字段值
		CopyField(ctx context.Context, params model.TaskVideoCopyFieldReq) (string, error)
		// 视频文件
		Video(ctx context.Context, fileName string) (string, error)
		// 视频任务
		Task(ctx context.Context)
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
