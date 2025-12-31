package log_batch

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/log_batch/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	batchPageRes, err := service.LogBatch().Page(ctx, req.LogBatchPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		LogBatchPageRes: batchPageRes,
	}

	return
}
