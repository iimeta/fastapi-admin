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
	IStatistics interface {
		// 用户数据
		DataUser(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error)
		// 应用数据
		DataApp(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error)
		// 应用密钥数据
		DataAppKey(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error)
		// 统计任务
		StatisticsTask(ctx context.Context)
		// 统计聊天数据
		StatisticsChat(ctx context.Context)
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
