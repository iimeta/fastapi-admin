package task_batch

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_batch/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) BatchPage(ctx context.Context, req *v1.BatchPageReq) (res *v1.BatchPageRes, err error) {

	pageRes, err := service.TaskBatch().Page(ctx, req.TaskBatchPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.BatchPageRes{
		TaskBatchPageRes: pageRes,
	}

	return
}
