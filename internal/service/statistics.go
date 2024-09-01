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
		// 统计聊天数据
		StatisticsChat(ctx context.Context)
		// 统计绘图数据
		StatisticsImage(ctx context.Context)
		// 统计音频数据
		StatisticsAudio(ctx context.Context)
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
