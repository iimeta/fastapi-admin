package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model/v1"
)

func (c *ControllerV1) Permissions(ctx context.Context, req *v1.PermissionsReq) (res *v1.PermissionsRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	items, err := service.Model().Permissions(ctx, req.ModelPermissionsReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PermissionsRes{
		ModelPermissionsRes: &model.ModelPermissionsRes{
			Items: items,
		},
	}

	return
}
