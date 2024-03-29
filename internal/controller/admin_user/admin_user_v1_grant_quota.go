package admin_user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/admin_user/v1"
)

func (c *ControllerV1) GrantQuota(ctx context.Context, req *v1.GrantQuotaReq) (res *v1.GrantQuotaRes, err error) {

	err = service.AdminUser().GrantQuota(ctx, req.UserGrantQuotaReq)

	return
}
