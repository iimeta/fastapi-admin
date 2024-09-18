package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) ChangeEmail(ctx context.Context, req *v1.ChangeEmailReq) (res *v1.ChangeEmailRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

	if service.Session().IsUserRole(ctx) {
		err = service.User().ChangeEmail(ctx, req.UserChangeEmailReq)
	}

	if service.Session().IsAdminRole(ctx) {
		err = service.SysAdmin().ChangeEmail(ctx, req.UserChangeEmailReq)
	}

	return
}
