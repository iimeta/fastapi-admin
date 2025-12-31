package app

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/app/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	appPageRes, err := service.App().Page(ctx, req.AppPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		AppPageRes: appPageRes,
	}

	return
}
