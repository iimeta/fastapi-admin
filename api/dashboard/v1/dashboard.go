package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 仪表盘基础数据接口请求参数
type BaseDataReq struct {
	g.Meta `path:"/base/data" method:"get" auth:"true" role:"user,reseller,admin" tags:"dashboard" summary:"仪表盘基础数据接口"`
}

// 仪表盘基础数据接口响应参数
type BaseDataRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardBaseDataRes
}

// 仪表盘调用数据接口请求参数
type CallDataReq struct {
	g.Meta `path:"/call/data/:days" method:"get" auth:"true" role:"user,reseller,admin" tags:"dashboard" summary:"仪表盘调用数据接口"`
	model.DashboardCallDataReq
}

// 仪表盘调用数据接口响应参数
type CallDataRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardCallDataRes
}

// 仪表盘费用接口请求参数
type ExpenseReq struct {
	g.Meta `path:"/expense" method:"get" auth:"true" role:"user,reseller,admin" tags:"dashboard" summary:"仪表盘费用接口"`
}

// 仪表盘费用接口响应参数
type ExpenseRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardExpenseRes
}

// 仪表盘数据TOP接口请求参数
type DataTopReq struct {
	g.Meta `path:"/data/top" method:"get" auth:"true" role:"user,reseller,admin" tags:"dashboard" summary:"仪表盘数据TOP接口"`
	model.DashboardDataTopReq
}

// 仪表盘数据TOP接口响应参数
type DataTopRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardDataTopRes
}

// 仪表盘模型占比接口请求参数
type ModelPercentReq struct {
	g.Meta `path:"/model/percent/:days" method:"get" auth:"true" role:"user,reseller,admin" tags:"dashboard" summary:"仪表盘模型占比接口"`
	model.DashboardModelPercentReq
}

// 仪表盘模型占比接口响应参数
type ModelPercentRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardModelPercentRes
}

// 每秒钟数据接口请求参数
type PerSecondReq struct {
	g.Meta `path:"/per/second" method:"post" auth:"true" role:"user,reseller,admin" tags:"dashboard" summary:"每秒钟数据接口"`
	model.DashboardPerSecondReq
}

// 每秒钟数据接口响应参数
type PerSecondRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardPerSecondRes
}

// 每分钟数据接口请求参数
type PerMinuteReq struct {
	g.Meta `path:"/per/minute" method:"post" auth:"true" role:"user,reseller,admin" tags:"dashboard" summary:"每分钟数据接口"`
	model.DashboardPerMinuteReq
}

// 每分钟数据接口响应参数
type PerMinuteRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardPerMinuteRes
}

// 额度预警接口请求参数
type QuotaWarningReq struct {
	g.Meta `path:"/quota/warning" method:"post" auth:"true" role:"reseller,user" tags:"dashboard" summary:"额度预警接口"`
	model.DashboardQuotaWarningReq
}

// 额度预警接口响应参数
type QuotaWarningRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
