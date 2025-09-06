package model

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/model/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) InitSync(ctx context.Context, req *v1.InitSyncReq) (res *v1.InitSyncRes, err error) {

	err = service.Model().InitSync(ctx, req.ModelInitSyncReq)

	return
}
