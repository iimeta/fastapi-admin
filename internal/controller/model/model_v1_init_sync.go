package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model/v1"
)

func (c *ControllerV1) InitSync(ctx context.Context, req *v1.InitSyncReq) (res *v1.InitSyncRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.Model().InitSync(ctx, req.ModelInitSyncReq)

	return
}
