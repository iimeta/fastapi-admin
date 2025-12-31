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
	ILogBatch interface {
		// 批处理日志详情
		Detail(ctx context.Context, id string) (*model.LogBatch, error)
		// 批处理日志分页列表
		Page(ctx context.Context, params model.LogBatchPageReq) (*model.LogBatchPageRes, error)
		// 批处理日志详情复制字段值
		CopyField(ctx context.Context, params model.LogBatchCopyFieldReq) (string, error)
	}
)

var (
	localLogBatch ILogBatch
)

func LogBatch() ILogBatch {
	if localLogBatch == nil {
		panic("implement not found for interface ILogBatch, forgot register?")
	}
	return localLogBatch
}

func RegisterLogBatch(i ILogBatch) {
	localLogBatch = i
}
