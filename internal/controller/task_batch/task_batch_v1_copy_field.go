package task_batch

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_batch/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.TaskBatch().CopyField(ctx, req.TaskBatchCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		TaskBatchCopyFieldRes: &model.TaskBatchCopyFieldRes{
			Value: value,
		},
	}

	return
}
