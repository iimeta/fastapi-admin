package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	items, err := service.Model().List(ctx, req.ModelListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		ModelListRes: &model.ModelListRes{
			Items: items,
		},
	}

	return
}
