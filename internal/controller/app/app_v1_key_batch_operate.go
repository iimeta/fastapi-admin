package app

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/app/v1"
)

func (c *ControllerV1) KeyBatchOperate(ctx context.Context, req *v1.KeyBatchOperateReq) (res *v1.KeyBatchOperateRes, err error) {

	keys, err := service.App().KeyBatchOperate(ctx, req.AppKeyBatchOperateReq)
	if err != nil {
		return nil, err
	}

	res = &v1.KeyBatchOperateRes{
		Keys: keys,
	}

	return
}
