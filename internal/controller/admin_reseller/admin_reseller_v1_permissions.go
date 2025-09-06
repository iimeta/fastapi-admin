package admin_reseller

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/admin_reseller/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Permissions(ctx context.Context, req *v1.PermissionsReq) (res *v1.PermissionsRes, err error) {

	err = service.AdminReseller().Permissions(ctx, req.ResellerPermissionsReq)

	return
}
