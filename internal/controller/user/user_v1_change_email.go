package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ChangeEmail(ctx context.Context, req *v1.ChangeEmailReq) (res *v1.ChangeEmailRes, err error) {

	if service.Session().IsResellerRole(ctx) {
		err = service.Reseller().ChangeEmail(ctx, req.UserChangeEmailReq)
	}

	if service.Session().IsUserRole(ctx) {
		err = service.User().ChangeEmail(ctx, req.UserChangeEmailReq)
	}

	if service.Session().IsAdminRole(ctx) {
		err = service.SysAdmin().ChangeEmail(ctx, req.UserChangeEmailReq)
	}

	return
}
