package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

// 用户数据接口请求参数
type DataUserReq struct {
	g.Meta `path:"/data/user" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"用户数据接口"`
	model.StatisticsDataReq
}

// 用户数据接口响应参数
type DataUserRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsDataRes
}

// 应用数据接口请求参数
type DataAppReq struct {
	g.Meta `path:"/data/app" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"应用数据接口"`
	model.StatisticsDataReq
}

// 应用数据接口响应参数
type DataAppRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsDataRes
}

// 应用密钥数据接口请求参数
type DataAppKeyReq struct {
	g.Meta `path:"/data/app/key" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"应用密钥数据接口"`
	model.StatisticsDataReq
}

// 应用密钥数据接口响应参数
type DataAppKeyRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsDataRes
}

// 数据看板汇总接口请求参数
type DataSummaryReq struct {
	g.Meta `path:"/data/summary" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"数据看板汇总接口"`
	model.StatisticsSummaryReq
}

// 数据看板汇总接口响应参数
type DataSummaryRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsSummaryRes
}

// 数据看板趋势接口请求参数
type DataTrendReq struct {
	g.Meta `path:"/data/trend" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"数据看板趋势接口"`
	model.StatisticsTrendReq
}

// 数据看板趋势接口响应参数
type DataTrendRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsTrendRes
}

// 数据看板模型分布接口请求参数
type DataModelPercentReq struct {
	g.Meta `path:"/data/model/percent" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"数据看板模型分布接口"`
	model.StatisticsModelPercentReq
}

// 数据看板模型分布接口响应参数
type DataModelPercentRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsModelPercentRes
}

// 数据看板排行接口请求参数
type DataTopReq struct {
	g.Meta `path:"/data/top" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"数据看板排行接口"`
	model.StatisticsTopReq
}

// 数据看板排行接口响应参数
type DataTopRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsTopRes
}

// 数据看板明细接口请求参数
type DataDetailReq struct {
	g.Meta `path:"/data/detail" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"数据看板明细接口"`
	model.StatisticsDetailReq
}

// 数据看板明细接口响应参数
type DataDetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsDetailRes
}

// 数据看板全局总览接口请求参数
type DataOverviewReq struct {
	g.Meta `path:"/data/overview" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"数据看板全局总览接口"`
	model.StatisticsOverviewReq
}

// 数据看板全局总览接口响应参数
type DataOverviewRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsOverviewRes
}

// 数据看板模型趋势接口请求参数
type DataModelTrendReq struct {
	g.Meta `path:"/data/model/trend" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"数据看板模型趋势接口"`
	model.StatisticsModelTrendReq
}

// 数据看板模型趋势接口响应参数
type DataModelTrendRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsModelTrendRes
}

// 数据看板响应耗时趋势接口请求参数
type DataLatencyTrendReq struct {
	g.Meta `path:"/data/latency/trend" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"数据看板响应耗时趋势接口"`
	model.StatisticsLatencyTrendReq
}

// 数据看板响应耗时趋势接口响应参数
type DataLatencyTrendRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsLatencyTrendRes
}

// 数据看板任务状态分布接口请求参数
type DataTaskStatusReq struct {
	g.Meta `path:"/data/task/status" method:"post" auth:"true" role:"user,reseller,admin" tags:"statistics" summary:"数据看板任务状态分布接口"`
	model.StatisticsTaskStatusReq
}

// 数据看板任务状态分布接口响应参数
type DataTaskStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsTaskStatusRes
}

// 数据看板代理状态接口请求参数
type DataAgentStatusReq struct {
	g.Meta `path:"/data/agent/status" method:"post" auth:"true" role:"admin" tags:"statistics" summary:"数据看板代理状态接口"`
	model.StatisticsAgentStatusReq
}

// 数据看板代理状态接口响应参数
type DataAgentStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsAgentStatusRes
}

// 数据看板密钥状态接口请求参数
type DataKeyStatusReq struct {
	g.Meta `path:"/data/key/status" method:"post" auth:"true" role:"admin" tags:"statistics" summary:"数据看板密钥状态接口"`
	model.StatisticsKeyStatusReq
}

// 数据看板密钥状态接口响应参数
type DataKeyStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsKeyStatusRes
}
