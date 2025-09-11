package app_key

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/app_key/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Config(ctx context.Context, req *v1.ConfigReq) (res *v1.ConfigRes, err error) {

	key, err := service.AppKey().Config(ctx, req.AppKeyConfigReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ConfigRes{
		Key: key,
	}

	return
}
