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

func (c *ControllerV1) DataTop(ctx context.Context, req *v1.DataTopReq) (res *v1.DataTopRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

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
