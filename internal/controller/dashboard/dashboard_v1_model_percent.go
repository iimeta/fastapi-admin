package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
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
