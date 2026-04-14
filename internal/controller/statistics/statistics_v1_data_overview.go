package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataOverview(ctx context.Context, req *v1.DataOverviewReq) (res *v1.DataOverviewRes, err error) {

	data, err := service.Statistics().DataOverview(ctx, req.StatisticsOverviewReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataOverviewRes{
		StatisticsOverviewRes: data,
	}

	return
}
