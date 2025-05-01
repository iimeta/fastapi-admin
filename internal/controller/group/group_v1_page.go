package group

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/group/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	groupPageRes, err := service.Group().Page(ctx, req.GroupPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		GroupPageRes: groupPageRes,
	}

	return
}
