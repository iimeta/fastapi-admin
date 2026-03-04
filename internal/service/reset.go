// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IReset interface {
		// 重置任务
		Task(ctx context.Context)
	}
)

var (
	localReset IReset
)

func Reset() IReset {
	if localReset == nil {
		panic("implement not found for interface IReset, forgot register?")
	}
	return localReset
}

func RegisterReset(i IReset) {
	localReset = i
}
