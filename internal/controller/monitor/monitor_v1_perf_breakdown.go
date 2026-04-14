package monitor

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/monitor/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) PerfBreakdown(ctx context.Context, req *v1.PerfBreakdownReq) (res *v1.PerfBreakdownRes, err error) {

	data, err := service.Monitor().PerfBreakdown(ctx, req.MonitorPerfBreakdownReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PerfBreakdownRes{
		MonitorPerfBreakdownRes: data,
	}

	return
}
