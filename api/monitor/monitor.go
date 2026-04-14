// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package monitor

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/monitor/v1"
)

type IMonitorV1 interface {
	Global(ctx context.Context, req *v1.GlobalReq) (res *v1.GlobalRes, err error)
	PerfBreakdown(ctx context.Context, req *v1.PerfBreakdownReq) (res *v1.PerfBreakdownRes, err error)
	PerfHistory(ctx context.Context, req *v1.PerfHistoryReq) (res *v1.PerfHistoryRes, err error)
}
