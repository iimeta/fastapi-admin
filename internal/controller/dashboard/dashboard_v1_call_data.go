package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) CallData(ctx context.Context, req *v1.CallDataReq) (res *v1.CallDataRes, err error) {

	items, err := service.Dashboard().CallData(ctx, req.DashboardCallDataReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CallDataRes{
		DashboardCallDataRes: &model.DashboardCallDataRes{
			Items: items,
		},
	}

	return
}
