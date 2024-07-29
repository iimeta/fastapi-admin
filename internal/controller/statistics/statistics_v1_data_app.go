package statistics

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/statistics/v1"
)

func (c *ControllerV1) DataApp(ctx context.Context, req *v1.DataAppReq) (res *v1.DataAppRes, err error) {

	_, err = service.Statistics().DataApp(ctx, req.StatisticsDataReq)
	if err != nil {
		return nil, err
	}

	return
}
