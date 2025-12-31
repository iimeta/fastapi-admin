package model

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	modelPageRes, err := service.Model().Page(ctx, req.ModelPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		ModelPageRes: modelPageRes,
	}

	return
}
