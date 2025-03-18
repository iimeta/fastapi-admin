package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) ChangeAvatar(ctx context.Context, req *v1.ChangeAvatarReq) (res *v1.ChangeAvatarRes, err error) {

	if service.Session().IsUserRole(ctx) {
		err = service.User().ChangeAvatar(ctx, req.File)
	}

	if service.Session().IsAdminRole(ctx) {
		err = service.SysAdmin().ChangeAvatar(ctx, req.File)
	}

	return
}
