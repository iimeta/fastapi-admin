package dashboard

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
)

func (c *ControllerV1) ModelPercent(ctx context.Context, req *v1.ModelPercentReq) (res *v1.ModelPercentRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

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
