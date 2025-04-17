package admin_user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/admin_user/v1"
)

func (c *ControllerV1) Permissions(ctx context.Context, req *v1.PermissionsReq) (res *v1.PermissionsRes, err error) {

	err = service.AdminUser().Permissions(ctx, req.UserPermissionsReq)

	return
}
