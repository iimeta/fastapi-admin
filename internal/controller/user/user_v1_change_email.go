package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) ChangeEmail(ctx context.Context, req *v1.ChangeEmailReq) (res *v1.ChangeEmailRes, err error) {

	if service.Session().IsUserRole(ctx) {
		err = service.User().ChangeEmail(ctx, req.UserChangeEmailReq)
	}

	if service.Session().IsAdminRole(ctx) {
		err = service.SysAdmin().ChangeEmail(ctx, req.UserChangeEmailReq)
	}

	return
}
