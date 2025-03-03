// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	INotice interface {
		// 额度预警任务
		QuotaWarningTask(ctx context.Context)
	}
)

var (
	localNotice INotice
)

func Notice() INotice {
	if localNotice == nil {
		panic("implement not found for interface INotice, forgot register?")
	}
	return localNotice
}

func RegisterNotice(i INotice) {
	localNotice = i
}
