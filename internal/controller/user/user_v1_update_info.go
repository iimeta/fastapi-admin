package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) UpdateInfo(ctx context.Context, req *v1.UpdateInfoReq) (res *v1.UpdateInfoRes, err error) {

	if service.Session().IsResellerRole(ctx) {
		err = service.Reseller().UpdateInfo(ctx, req.UserUpdateInfoReq)
	}

	if service.Session().IsUserRole(ctx) {
		err = service.User().UpdateInfo(ctx, req.UserUpdateInfoReq)
	}

	if service.Session().IsAdminRole(ctx) {
		err = service.SysAdmin().UpdateInfo(ctx, req.UserUpdateInfoReq)
	}

	return
}
