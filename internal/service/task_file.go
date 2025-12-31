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
	ITaskFile interface {
		// 文件任务详情
		Detail(ctx context.Context, id string) (*model.TaskFile, error)
		// 文件任务分页列表
		Page(ctx context.Context, params model.TaskFilePageReq) (*model.TaskFilePageRes, error)
		// 文件任务详情复制字段值
		CopyField(ctx context.Context, params model.TaskFileCopyFieldReq) (string, error)
		// 文件
		File(ctx context.Context, fileName string) (string, error)
		// 文件任务
		Task(ctx context.Context)
	}
)

var (
	localTaskFile ITaskFile
)

func TaskFile() ITaskFile {
	if localTaskFile == nil {
		panic("implement not found for interface ITaskFile, forgot register?")
	}
	return localTaskFile
}

func RegisterTaskFile(i ITaskFile) {
	localTaskFile = i
}
