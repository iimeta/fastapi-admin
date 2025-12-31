package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ChangeAvatar(ctx context.Context, req *v1.ChangeAvatarReq) (res *v1.ChangeAvatarRes, err error) {

	if service.Session().IsResellerRole(ctx) {
		err = service.Reseller().ChangeAvatar(ctx, req.File)
	}

	if service.Session().IsUserRole(ctx) {
		err = service.User().ChangeAvatar(ctx, req.File)
	}

	if service.Session().IsAdminRole(ctx) {
		err = service.SysAdmin().ChangeAvatar(ctx, req.File)
	}

	return
}
