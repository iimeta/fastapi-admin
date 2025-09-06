package app

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/app/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) KeyConfig(ctx context.Context, req *v1.KeyConfigReq) (res *v1.KeyConfigRes, err error) {

	key, err := service.App().KeyConfig(ctx, req.AppKeyConfigReq)
	if err != nil {
		return nil, err
	}

	res = &v1.KeyConfigRes{
		Key: key,
	}

	return
}
