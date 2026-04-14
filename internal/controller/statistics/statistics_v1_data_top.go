package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataTop(ctx context.Context, req *v1.DataTopReq) (res *v1.DataTopRes, err error) {

	data, err := service.Statistics().DataTop(ctx, req.StatisticsTopReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataTopRes{
		StatisticsTopRes: data,
	}

	return
}
