package app

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/app/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	app, err := service.App().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		AppDetailRes: &model.AppDetailRes{
			App: app,
		},
	}

	return
}
