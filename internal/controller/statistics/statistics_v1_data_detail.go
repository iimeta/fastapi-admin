package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataDetail(ctx context.Context, req *v1.DataDetailReq) (res *v1.DataDetailRes, err error) {

	data, err := service.Statistics().DataDetail(ctx, req.StatisticsDetailReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataDetailRes{
		StatisticsDetailRes: data,
	}

	return
}
