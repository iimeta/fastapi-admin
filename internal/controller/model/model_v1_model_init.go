package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model/v1"
)

func (c *ControllerV1) ModelInit(ctx context.Context, req *v1.ModelInitReq) (res *v1.ModelInitRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.Model().Init(ctx, req.ModelInitReq)

	return
}
