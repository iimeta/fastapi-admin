package model

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

	modelPageRes, err := service.Model().Page(ctx, req.ModelPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		ModelPageRes: modelPageRes,
	}

	return
}
