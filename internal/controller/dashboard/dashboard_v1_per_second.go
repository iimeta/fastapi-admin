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

func (c *ControllerV1) PerSecond(ctx context.Context, req *v1.PerSecondReq) (res *v1.PerSecondRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

	rps, tps, err := service.Dashboard().PerSecond(ctx, req.DashboardPerSecondReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PerSecondRes{
		DashboardPerSecondRes: &model.DashboardPerSecondRes{
			RPS: rps,
			TPS: tps,
		},
	}

	return
}
