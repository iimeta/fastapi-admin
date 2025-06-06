package app

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/app/v1"
)

func (c *ControllerV1) BatchCreateKey(ctx context.Context, req *v1.BatchCreateKeyReq) (res *v1.BatchCreateKeyRes, err error) {

	keys, err := service.App().BatchCreateKey(ctx, req.AppBatchCreateKeyReq)
	if err != nil {
		return nil, err
	}

	res = &v1.BatchCreateKeyRes{
		Keys: keys,
	}

	return
}
