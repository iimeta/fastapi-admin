package model

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/model/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	m, err := service.Model().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		ModelDetailRes: &model.ModelDetailRes{
			Model: m,
		},
	}

	return
}
