package app_key

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/app_key/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error) {

	keys, err := service.AppKey().BatchOperate(ctx, req.AppKeyBatchOperateReq)
	if err != nil {
		return nil, err
	}

	res = &v1.BatchOperateRes{
		Keys: keys,
	}

	return
}
