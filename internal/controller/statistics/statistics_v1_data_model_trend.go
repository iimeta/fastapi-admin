package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataModelTrend(ctx context.Context, req *v1.DataModelTrendReq) (res *v1.DataModelTrendRes, err error) {

	data, err := service.Statistics().DataModelTrend(ctx, req.StatisticsModelTrendReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataModelTrendRes{
		StatisticsModelTrendRes: data,
	}

	return
}
