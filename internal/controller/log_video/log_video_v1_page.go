package log_video

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_video/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	videoPageRes, err := service.LogVideo().Page(ctx, req.LogVideoPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		LogVideoPageRes: videoPageRes,
	}

	return
}
