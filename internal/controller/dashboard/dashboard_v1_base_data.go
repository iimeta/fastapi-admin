package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) BaseData(ctx context.Context, req *v1.BaseDataReq) (res *v1.BaseDataRes, err error) {

	dashboard, err := service.Dashboard().BaseData(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.BaseDataRes{
		DashboardBaseDataRes: &model.DashboardBaseDataRes{
			Dashboard: dashboard,
		},
	}

	return
}
