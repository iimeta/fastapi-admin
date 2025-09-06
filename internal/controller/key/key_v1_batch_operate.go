package key

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/key/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error) {

	err = service.Key().BatchOperate(ctx, req.KeyBatchOperateReq)

	return
}
