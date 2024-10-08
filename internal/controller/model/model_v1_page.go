package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	modelPageRes, err := service.Model().Page(ctx, req.ModelPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		ModelPageRes: modelPageRes,
	}

	return
}
