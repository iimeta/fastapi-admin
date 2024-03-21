package dashboard

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
)

func (c *ControllerV1) DataTop5(ctx context.Context, req *v1.DataTop5Req) (res *v1.DataTop5Res, err error) {

	items, err := service.Dashboard().DataTop5(ctx, req.DashboardDataTop5Req)
	if err != nil {
		return nil, err
	}

	res = &v1.DataTop5Res{
		DashboardDataTop5Res: &model.DashboardDataTop5Res{
			Items: items,
		},
	}

	return
}
