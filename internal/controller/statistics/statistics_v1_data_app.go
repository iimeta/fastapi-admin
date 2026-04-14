package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataApp(ctx context.Context, req *v1.DataAppReq) (res *v1.DataAppRes, err error) {

	data, err := service.Statistics().DataApp(ctx, req.StatisticsDataReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataAppRes{
		StatisticsDataRes: data,
	}

	return
}
