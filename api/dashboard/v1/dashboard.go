package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 仪表盘基础数据接口请求参数
type BaseDataReq struct {
	g.Meta `path:"/base/data" tags:"dashboard" method:"get" summary:"仪表盘基础数据接口"`
}

// 仪表盘基础数据接口响应参数
type BaseDataRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardBaseDataRes
}

// 仪表盘调用数据接口请求参数
type CallDataReq struct {
	g.Meta `path:"/call/data/:days" tags:"dashboard" method:"get" summary:"仪表盘调用数据接口"`
	model.DashboardCallDataReq
}

// 仪表盘调用数据接口响应参数
type CallDataRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardCallDataRes
}

// 仪表盘费用接口请求参数
type ExpenseReq struct {
	g.Meta `path:"/expense" tags:"dashboard" method:"get" summary:"仪表盘费用接口"`
}

// 仪表盘费用接口响应参数
type ExpenseRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardExpenseRes
}

// 仪表盘数据TOP5接口请求参数
type DataTop5Req struct {
	g.Meta `path:"/data/top5/:days" tags:"dashboard" method:"get" summary:"仪表盘数据TOP5接口"`
	model.DashboardDataTop5Req
}

// 仪表盘数据TOP5接口响应参数
type DataTop5Res struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardDataTop5Res
}

// 仪表盘模型占比接口请求参数
type ModelPercentReq struct {
	g.Meta `path:"/model/percent/:days" tags:"dashboard" method:"get" summary:"仪表盘模型占比接口"`
	model.DashboardModelPercentReq
}

// 仪表盘模型占比接口响应参数
type ModelPercentRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardModelPercentRes
}
