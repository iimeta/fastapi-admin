package log_image

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_image/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	imagePageRes, err := service.LogImage().Page(ctx, req.LogImagePageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		LogImagePageRes: imagePageRes,
	}

	return
}
