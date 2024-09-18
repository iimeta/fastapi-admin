package app

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/app/v1"
)

func (c *ControllerV1) CreateKey(ctx context.Context, req *v1.CreateKeyReq) (res *v1.CreateKeyRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

	key, err := service.App().CreateKey(ctx, req.AppCreateKeyReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CreateKeyRes{
		AppCreateKeyRes: &model.AppCreateKeyRes{
			AppId: req.AppId,
			Key:   key,
		},
	}

	return
}
