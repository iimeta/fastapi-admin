package dashboard

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
)

func (c *ControllerV1) PerSecond(ctx context.Context, req *v1.PerSecondReq) (res *v1.PerSecondRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

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
