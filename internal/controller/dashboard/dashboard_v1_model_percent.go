package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) ModelPercent(ctx context.Context, req *v1.ModelPercentReq) (res *v1.ModelPercentRes, err error) {

	models, items, err := service.Dashboard().ModelPercent(ctx, req.DashboardModelPercentReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ModelPercentRes{
		DashboardModelPercentRes: &model.DashboardModelPercentRes{
			Models: models,
			Items:  items,
		},
	}

	return
}
