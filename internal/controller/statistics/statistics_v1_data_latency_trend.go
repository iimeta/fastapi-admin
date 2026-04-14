package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataLatencyTrend(ctx context.Context, req *v1.DataLatencyTrendReq) (res *v1.DataLatencyTrendRes, err error) {

	data, err := service.Statistics().DataLatencyTrend(ctx, req.StatisticsLatencyTrendReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataLatencyTrendRes{
		StatisticsLatencyTrendRes: data,
	}

	return
}
