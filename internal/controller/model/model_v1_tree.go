package model

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Tree(ctx context.Context, req *v1.TreeReq) (res *v1.TreeRes, err error) {

	items, err := service.Model().Tree(ctx, req.ModelTreeReq)
	if err != nil {
		return nil, err
	}

	res = &v1.TreeRes{
		ModelTreeRes: &model.ModelTreeRes{
			Items: items,
		},
	}

	return
}
