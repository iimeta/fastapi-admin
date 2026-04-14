package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataTrend(ctx context.Context, req *v1.DataTrendReq) (res *v1.DataTrendRes, err error) {

	data, err := service.Statistics().DataTrend(ctx, req.StatisticsTrendReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataTrendRes{
		StatisticsTrendRes: data,
	}

	return
}
