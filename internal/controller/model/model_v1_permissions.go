package model

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Permissions(ctx context.Context, req *v1.PermissionsReq) (res *v1.PermissionsRes, err error) {

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
