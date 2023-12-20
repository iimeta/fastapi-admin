package app

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/app/v1"
)

func (c *ControllerV1) CreateKey(ctx context.Context, req *v1.CreateKeyReq) (res *v1.CreateKeyRes, err error) {

	err = service.App().CreateKey(ctx, req.AppCreateKeyReq)

	return
}
