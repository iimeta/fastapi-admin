package dashboard

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
)

func (c *ControllerV1) WarningConfig(ctx context.Context, req *v1.WarningConfigReq) (res *v1.WarningConfigRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.Dashboard().WarningConfig(ctx, req.DashboardWarningConfigReq)

	return
}
