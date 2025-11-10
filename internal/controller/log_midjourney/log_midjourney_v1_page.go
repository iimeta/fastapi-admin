package log_midjourney

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_midjourney/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	midjourneyPageRes, err := service.LogMidjourney().Page(ctx, req.LogMidjourneyPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		LogMidjourneyPageRes: midjourneyPageRes,
	}

	return
}
