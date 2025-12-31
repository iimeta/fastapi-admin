package app_key

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/app_key/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	key, err := service.AppKey().Create(ctx, req.AppKeyCreateReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CreateRes{
		AppKeyCreateRes: &model.AppKeyCreateRes{
			AppId: req.AppId,
			Key:   key,
		},
	}

	return
}
