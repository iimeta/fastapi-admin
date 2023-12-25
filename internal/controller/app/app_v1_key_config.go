package app

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/app/v1"
)

func (c *ControllerV1) KeyConfig(ctx context.Context, req *v1.KeyConfigReq) (res *v1.KeyConfigRes, err error) {

	err = service.App().KeyConfig(ctx, req.AppKeyConfigReq)

	return
}
