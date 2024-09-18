package finance

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/finance/v1"
)

func (c *ControllerV1) DealRecordPage(ctx context.Context, req *v1.DealRecordPageReq) (res *v1.DealRecordPageRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

	dealRecordPageRes, err := service.Finance().DealRecordPage(ctx, req.FinanceDealRecordPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DealRecordPageRes{
		FinanceDealRecordPageRes: dealRecordPageRes,
	}

	return
}
