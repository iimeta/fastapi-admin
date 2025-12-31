package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
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
