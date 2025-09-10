package provider

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/provider/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error) {

	err = service.Provider().BatchOperate(ctx, req.ProviderBatchOperateReq)

	return
}
