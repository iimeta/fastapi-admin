package app

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/app/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) ChangeStatus(ctx context.Context, req *v1.ChangeStatusReq) (res *v1.ChangeStatusRes, err error) {

	err = service.App().ChangeStatus(ctx, req.AppChangeStatusReq)

	return
}
