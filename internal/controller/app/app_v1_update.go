package app

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/app/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	err = service.App().Update(ctx, req.AppUpdateReq)

	return
}
