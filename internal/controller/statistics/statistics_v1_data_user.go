package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataUser(ctx context.Context, req *v1.DataUserReq) (res *v1.DataUserRes, err error) {

	data, err := service.Statistics().DataUser(ctx, req.StatisticsDataReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataUserRes{
		StatisticsDataRes: data,
	}

	return
}
