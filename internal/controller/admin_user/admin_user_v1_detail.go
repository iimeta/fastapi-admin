package admin_user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/admin_user/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

	user, err := service.AdminUser().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		UserDetailRes: &model.UserDetailRes{
			User: user,
		},
	}

	return
}
