package dashboard

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
)

func (c *ControllerV1) DataTop(ctx context.Context, req *v1.DataTopReq) (res *v1.DataTopRes, err error) {

	items, err := service.Dashboard().DataTopNew(ctx, req.DashboardDataTopReq)
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
