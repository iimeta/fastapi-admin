package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model/v1"
)

func (c *ControllerV1) SyncInit(ctx context.Context, req *v1.SyncInitReq) (res *v1.SyncInitRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.Model().SyncInit(ctx, req.ModelSyncInitReq)

	return
}
