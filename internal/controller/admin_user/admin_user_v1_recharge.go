package admin_user

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/admin_user/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Recharge(ctx context.Context, req *v1.RechargeReq) (res *v1.RechargeRes, err error) {

	err = service.AdminUser().Recharge(ctx, req.UserRechargeReq)

	return
}
