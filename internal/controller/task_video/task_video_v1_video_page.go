package task_video

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_video/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) VideoPage(ctx context.Context, req *v1.VideoPageReq) (res *v1.VideoPageRes, err error) {

	pageRes, err := service.TaskVideo().Page(ctx, req.TaskVideoPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.VideoPageRes{
		TaskVideoPageRes: pageRes,
	}

	return
}
