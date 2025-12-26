package task_batch

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/task_batch/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	taskBatch, err := service.TaskBatch().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		TaskBatchDetailRes: &model.TaskBatchDetailRes{
			TaskBatch: taskBatch,
		},
	}

	return
}
