package key

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/key/v1"
)

func (c *ControllerV1) BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.Key().BatchOperate(ctx, req.KeyBatchOperateReq)

	return
}
