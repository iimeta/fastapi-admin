package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataAgentStatus(ctx context.Context, req *v1.DataAgentStatusReq) (res *v1.DataAgentStatusRes, err error) {

	data, err := service.Statistics().DataAgentStatus(ctx, req.StatisticsAgentStatusReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataAgentStatusRes{
		StatisticsAgentStatusRes: data,
	}

	return
}
