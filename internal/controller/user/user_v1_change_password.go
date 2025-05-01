package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error) {

	if service.Session().IsResellerRole(ctx) {
		err = service.Reseller().ChangePassword(ctx, req.UserChangePasswordReq)
	}

	if service.Session().IsUserRole(ctx) {
		err = service.User().ChangePassword(ctx, req.UserChangePasswordReq)
	}

	if service.Session().IsAdminRole(ctx) {
		err = service.SysAdmin().ChangePassword(ctx, req.UserChangePasswordReq)
	}

	return
}
