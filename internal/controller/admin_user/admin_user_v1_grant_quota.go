package admin_user

import (
	"context"
	"github.com/iimeta/fastapi-admin/api/admin_user/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) GrantQuota(ctx context.Context, req *v1.GrantQuotaReq) (res *v1.GrantQuotaRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.AdminUser().GrantQuota(ctx, req.UserGrantQuotaReq)

	return
}
