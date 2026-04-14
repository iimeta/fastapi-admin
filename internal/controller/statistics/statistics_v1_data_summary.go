package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataSummary(ctx context.Context, req *v1.DataSummaryReq) (res *v1.DataSummaryRes, err error) {

	data, err := service.Statistics().DataSummary(ctx, req.StatisticsSummaryReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataSummaryRes{
		StatisticsSummaryRes: data,
	}

	return
}
