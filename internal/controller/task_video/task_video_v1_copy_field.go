package task_video

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/task_video/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.TaskVideo().CopyField(ctx, req.TaskVideoCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		TaskVideoCopyFieldRes: &model.TaskVideoCopyFieldRes{
			Value: value,
		},
	}

	return
}
