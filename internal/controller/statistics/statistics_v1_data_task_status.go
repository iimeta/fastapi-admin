package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataTaskStatus(ctx context.Context, req *v1.DataTaskStatusReq) (res *v1.DataTaskStatusRes, err error) {

	data, err := service.Statistics().DataTaskStatus(ctx, req.StatisticsTaskStatusReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataTaskStatusRes{
		StatisticsTaskStatusRes: data,
	}

	return
}
