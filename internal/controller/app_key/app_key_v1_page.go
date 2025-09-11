package app_key

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/app_key/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	appKeyPageRes, err := service.AppKey().Page(ctx, req.AppKeyPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		AppKeyPageRes: appKeyPageRes,
	}

	return
}
