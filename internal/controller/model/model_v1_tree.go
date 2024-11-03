package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model/v1"
)

func (c *ControllerV1) Tree(ctx context.Context, req *v1.TreeReq) (res *v1.TreeRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

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
