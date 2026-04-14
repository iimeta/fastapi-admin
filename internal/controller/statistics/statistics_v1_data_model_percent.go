package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataModelPercent(ctx context.Context, req *v1.DataModelPercentReq) (res *v1.DataModelPercentRes, err error) {

	data, err := service.Statistics().DataModelPercent(ctx, req.StatisticsModelPercentReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataModelPercentRes{
		StatisticsModelPercentRes: data,
	}

	return
}
