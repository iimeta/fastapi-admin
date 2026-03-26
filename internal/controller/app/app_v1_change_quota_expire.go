package app

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/internal/service"

	"github.com/iimeta/fastapi-admin/v2/api/app/v1"
)

func (c *ControllerV1) ChangeQuotaExpire(ctx context.Context, req *v1.ChangeQuotaExpireReq) (res *v1.ChangeQuotaExpireRes, err error) {

	err = service.App().ChangeQuotaExpire(ctx, req.AppChangeQuotaExpireReq)

	return
}
