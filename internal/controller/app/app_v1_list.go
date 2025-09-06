package app

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/app/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.App().List(ctx, req.AppListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		AppListRes: &model.AppListRes{
			Items: items,
		},
	}

	return
}
