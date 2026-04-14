package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

// 监控中心全局实时指标接口请求参数
type GlobalReq struct {
	g.Meta `path:"/global" method:"post" auth:"true" role:"user,reseller,admin" tags:"monitor" summary:"监控中心全局实时指标接口"`
	model.MonitorGlobalReq
}

// 监控中心全局实时指标接口响应参数
type GlobalRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.MonitorGlobalRes
}

// 监控中心实时性能维度分析接口请求参数
type PerfBreakdownReq struct {
	g.Meta `path:"/perf/breakdown" method:"post" auth:"true" role:"user,reseller,admin" tags:"monitor" summary:"监控中心实时性能维度分析接口"`
	model.MonitorPerfBreakdownReq
}

// 监控中心实时性能维度分析接口响应参数
type PerfBreakdownRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.MonitorPerfBreakdownRes
}

// 监控中心历史性能数据接口请求参数
type PerfHistoryReq struct {
	g.Meta `path:"/perf/history" method:"post" auth:"true" role:"user,reseller,admin" tags:"monitor" summary:"监控中心历史性能数据接口"`
	model.MonitorPerfHistoryReq
}

// 监控中心历史性能数据接口响应参数
type PerfHistoryRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.MonitorPerfHistoryRes
}
