package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataKeyStatus(ctx context.Context, req *v1.DataKeyStatusReq) (res *v1.DataKeyStatusRes, err error) {

	data, err := service.Statistics().DataKeyStatus(ctx, req.StatisticsKeyStatusReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataKeyStatusRes{
		StatisticsKeyStatusRes: data,
	}

	return
}
