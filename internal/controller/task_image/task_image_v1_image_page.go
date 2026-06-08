package task_image

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_image/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ImagePage(ctx context.Context, req *v1.ImagePageReq) (res *v1.ImagePageRes, err error) {

	pageRes, err := service.TaskImage().Page(ctx, req.TaskImagePageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ImagePageRes{
		TaskImagePageRes: pageRes,
	}

	return
}
