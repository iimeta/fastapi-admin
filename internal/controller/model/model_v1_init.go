package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model/v1"
)

func (c *ControllerV1) Init(ctx context.Context, req *v1.InitReq) (res *v1.InitRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.Model().Init(ctx, req.ModelInitReq)

	return
}
