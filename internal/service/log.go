// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ILog interface {
		// 删除任务
		DelTask(ctx context.Context)
	}
)

var (
	localLog ILog
)

func Log() ILog {
	if localLog == nil {
		panic("implement not found for interface ILog, forgot register?")
	}
	return localLog
}

func RegisterLog(i ILog) {
	localLog = i
}
