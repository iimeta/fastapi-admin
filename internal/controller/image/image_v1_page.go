package image

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/image/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	imagePageRes, err := service.Image().Page(ctx, req.ImagePageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		ImagePageRes: imagePageRes,
	}

	return
}
