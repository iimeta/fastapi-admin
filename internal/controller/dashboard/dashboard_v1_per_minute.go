package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) PerMinute(ctx context.Context, req *v1.PerMinuteReq) (res *v1.PerMinuteRes, err error) {

	rpm, tpm, err := service.Dashboard().PerMinute(ctx, req.DashboardPerMinuteReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PerMinuteRes{
		DashboardPerMinuteRes: &model.DashboardPerMinuteRes{
			RPM: rpm,
			TPM: tpm,
		},
	}

	return
}
