package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) DataTop(ctx context.Context, req *v1.DataTopReq) (res *v1.DataTopRes, err error) {

	items, err := service.Dashboard().DataTop(ctx, req.DashboardDataTopReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DataTopRes{
		DashboardDataTopRes: &model.DashboardDataTopRes{
			Items: items,
		},
	}

	return
}
