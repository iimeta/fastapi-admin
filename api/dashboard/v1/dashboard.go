package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 仪表盘基础数据接口请求参数
type BaseDataReq struct {
	g.Meta `path:"/base_data" tags:"dashboard" method:"get" summary:"仪表盘基础数据接口"`
}

// 仪表盘基础数据接口响应参数
type BaseDataRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.DashboardBaseDataRes
}
