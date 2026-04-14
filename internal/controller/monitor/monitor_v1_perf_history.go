package monitor

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/monitor/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) PerfHistory(ctx context.Context, req *v1.PerfHistoryReq) (res *v1.PerfHistoryRes, err error) {

	data, err := service.Monitor().PerfHistory(ctx, req.MonitorPerfHistoryReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PerfHistoryRes{
		MonitorPerfHistoryRes: data,
	}

	return
}
