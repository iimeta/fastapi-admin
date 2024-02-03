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
	g.Meta `path:"/call/data" tags:"dashboard" method:"get" summary:"仪表盘调用数据接口"`
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
