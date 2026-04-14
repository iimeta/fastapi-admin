package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataAppKey(ctx context.Context, req *v1.DataAppKeyReq) (res *v1.DataAppKeyRes, err error) {

	data, err := service.Statistics().DataAppKey(ctx, req.StatisticsDataReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataAppKeyRes{
		StatisticsDataRes: data,
	}

	return
}
