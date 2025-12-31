package admin_reseller

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/admin_reseller/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ChangeQuotaExpire(ctx context.Context, req *v1.ChangeQuotaExpireReq) (res *v1.ChangeQuotaExpireRes, err error) {

	err = service.AdminReseller().ChangeQuotaExpire(ctx, req.ResellerChangeQuotaExpireReq)

	return
}
