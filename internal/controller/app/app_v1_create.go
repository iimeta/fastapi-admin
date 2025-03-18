package app

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/app/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	key, err := service.App().Create(ctx, req.AppCreateReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CreateRes{
		Key: key,
	}

	return
}
