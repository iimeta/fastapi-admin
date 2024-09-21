package app

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/app/v1"
)

func (c *ControllerV1) CreateKey(ctx context.Context, req *v1.CreateKeyReq) (res *v1.CreateKeyRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	key, err := service.App().CreateKey(ctx, req.AppCreateKeyReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CreateKeyRes{
		AppCreateKeyRes: &model.AppCreateKeyRes{
			AppId: req.AppId,
			Key:   key,
		},
	}

	return
}
