package admin_user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/admin_user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	err = service.AdminUser().Update(ctx, req.UserUpdateReq)

	return
}
