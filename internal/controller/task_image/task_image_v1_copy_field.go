package task_image

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_image/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.TaskImage().CopyField(ctx, req.TaskImageCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		TaskImageCopyFieldRes: &model.TaskImageCopyFieldRes{
			Value: value,
		},
	}

	return
}
