package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) QuotaWarning(ctx context.Context, req *v1.QuotaWarningReq) (res *v1.QuotaWarningRes, err error) {

	err = service.Dashboard().QuotaWarning(ctx, req.DashboardQuotaWarningReq)

	return
}
