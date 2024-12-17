// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IStatistics interface {
		// 统计任务
		StatisticsTask(ctx context.Context)
		// 统计数据
		StatisticsData(ctx context.Context, collection string, index string, lastTimeKey string, lastIdKey string)
	}
)

var (
	localStatistics IStatistics
)

func Statistics() IStatistics {
	if localStatistics == nil {
		panic("implement not found for interface IStatistics, forgot register?")
	}
	return localStatistics
}

func RegisterStatistics(i IStatistics) {
	localStatistics = i
}
