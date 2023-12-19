package key

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/key/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	keyPageRes, err := service.Key().Page(ctx, req.KeyPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		KeyPageRes: keyPageRes,
	}

	return
}
