// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IQuota interface {
		// 额度通知任务
		NoticeTask(ctx context.Context)
		// 额度清零任务
		ClearTask(ctx context.Context)
	}
)

var (
	localQuota IQuota
)

func Quota() IQuota {
	if localQuota == nil {
		panic("implement not found for interface IQuota, forgot register?")
	}
	return localQuota
}

func RegisterQuota(i IQuota) {
	localQuota = i
}
