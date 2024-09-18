package finance

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/finance/v1"
)

func (c *ControllerV1) BillPage(ctx context.Context, req *v1.BillPageReq) (res *v1.BillPageRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

	billPageRes, err := service.Finance().BillPage(ctx, req.FinanceBillPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.BillPageRes{
		FinanceBillPageRes: billPageRes,
	}

	return
}
