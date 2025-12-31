package admin_user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/admin_user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ChangeQuotaExpire(ctx context.Context, req *v1.ChangeQuotaExpireReq) (res *v1.ChangeQuotaExpireRes, err error) {

	err = service.AdminUser().ChangeQuotaExpire(ctx, req.UserChangeQuotaExpireReq)

	return
}
