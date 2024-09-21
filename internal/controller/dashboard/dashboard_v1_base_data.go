package dashboard

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
)

func (c *ControllerV1) BaseData(ctx context.Context, req *v1.BaseDataReq) (res *v1.BaseDataRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

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
