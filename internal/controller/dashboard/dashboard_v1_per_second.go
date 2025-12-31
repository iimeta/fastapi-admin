package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) PerSecond(ctx context.Context, req *v1.PerSecondReq) (res *v1.PerSecondRes, err error) {

	rps, tps, err := service.Dashboard().PerSecond(ctx, req.DashboardPerSecondReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PerSecondRes{
		DashboardPerSecondRes: &model.DashboardPerSecondRes{
			RPS: rps,
			TPS: tps,
		},
	}

	return
}
