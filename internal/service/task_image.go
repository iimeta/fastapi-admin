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
	ITaskImage interface {
		// 绘图任务详情
		Detail(ctx context.Context, id string) (*model.TaskImage, error)
		// 绘图任务分页列表
		Page(ctx context.Context, params model.TaskImagePageReq) (*model.TaskImagePageRes, error)
		// 绘图任务详情复制字段值
		CopyField(ctx context.Context, params model.TaskImageCopyFieldReq) (string, error)
		// 绘图任务重新生成
		Regenerate(ctx context.Context, id string) error
		// 绘图任务批量操作
		BatchOperate(ctx context.Context, params model.TaskImageBatchOperateReq) error
		// 绘图任务
		Task(ctx context.Context)
	}
)

var (
	localTaskImage ITaskImage
)

func TaskImage() ITaskImage {
	if localTaskImage == nil {
		panic("implement not found for interface ITaskImage, forgot register?")
	}
	return localTaskImage
}

func RegisterTaskImage(i ITaskImage) {
	localTaskImage = i
}
