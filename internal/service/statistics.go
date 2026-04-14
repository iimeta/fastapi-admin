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
	IStatistics interface {
		// 用户数据
		DataUser(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error)
		// 应用数据
		DataApp(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error)
		// 应用密钥数据
		DataAppKey(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error)
		// 数据看板汇总
		DataSummary(ctx context.Context, params model.StatisticsSummaryReq) (*model.StatisticsSummaryRes, error)
		// 数据看板趋势
		DataTrend(ctx context.Context, params model.StatisticsTrendReq) (*model.StatisticsTrendRes, error)
		// 数据看板模型分布
		DataModelPercent(ctx context.Context, params model.StatisticsModelPercentReq) (*model.StatisticsModelPercentRes, error)
		// 数据看板排行
		DataTop(ctx context.Context, params model.StatisticsTopReq) (*model.StatisticsTopRes, error)
		// 数据看板明细
		DataDetail(ctx context.Context, params model.StatisticsDetailReq) (*model.StatisticsDetailRes, error)
		// 数据看板全局总览
		DataOverview(ctx context.Context, params model.StatisticsOverviewReq) (*model.StatisticsOverviewRes, error)
		// 数据看板模型趋势
		DataModelTrend(ctx context.Context, params model.StatisticsModelTrendReq) (*model.StatisticsModelTrendRes, error)
		// 数据看板响应耗时趋势
		DataLatencyTrend(ctx context.Context, params model.StatisticsLatencyTrendReq) (*model.StatisticsLatencyTrendRes, error)
		// 数据看板任务状态分布
		DataTaskStatus(ctx context.Context, params model.StatisticsTaskStatusReq) (*model.StatisticsTaskStatusRes, error)
		// 数据看板代理状态
		DataAgentStatus(ctx context.Context, params model.StatisticsAgentStatusReq) (*model.StatisticsAgentStatusRes, error)
		// 数据看板密钥状态
		DataKeyStatus(ctx context.Context, params model.StatisticsKeyStatusReq) (*model.StatisticsKeyStatusRes, error)
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
