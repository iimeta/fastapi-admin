package midjourney

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/midjourney/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	midjourneyPageRes, err := service.Midjourney().Page(ctx, req.MidjourneyPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		MidjourneyPageRes: midjourneyPageRes,
	}

	return
}
