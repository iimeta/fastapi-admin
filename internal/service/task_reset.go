// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ITaskReset interface {
		// 重置任务
		Task(ctx context.Context)
	}
)

var (
	localTaskReset ITaskReset
)

func TaskReset() ITaskReset {
	if localTaskReset == nil {
		panic("implement not found for interface ITaskReset, forgot register?")
	}
	return localTaskReset
}

func RegisterTaskReset(i ITaskReset) {
	localTaskReset = i
}
