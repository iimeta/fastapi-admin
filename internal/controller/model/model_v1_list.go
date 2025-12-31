package model

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

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
