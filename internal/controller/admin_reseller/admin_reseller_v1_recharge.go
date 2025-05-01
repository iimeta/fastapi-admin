package admin_reseller

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/admin_reseller/v1"
)

func (c *ControllerV1) Recharge(ctx context.Context, req *v1.RechargeReq) (res *v1.RechargeRes, err error) {

	err = service.AdminReseller().Recharge(ctx, req.ResellerRechargeReq)

	return
}
