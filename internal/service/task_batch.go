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
	ITaskBatch interface {
		// 批处理任务详情
		Detail(ctx context.Context, id string) (*model.TaskBatch, error)
		// 批处理任务分页列表
		Page(ctx context.Context, params model.TaskBatchPageReq) (*model.TaskBatchPageRes, error)
		// 批处理任务详情复制字段值
		CopyField(ctx context.Context, params model.TaskBatchCopyFieldReq) (string, error)
		// 批处理任务
		Task(ctx context.Context)
	}
)

var (
	localTaskBatch ITaskBatch
)

func TaskBatch() ITaskBatch {
	if localTaskBatch == nil {
		panic("implement not found for interface ITaskBatch, forgot register?")
	}
	return localTaskBatch
}

func RegisterTaskBatch(i ITaskBatch) {
	localTaskBatch = i
}
