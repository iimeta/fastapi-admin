package group

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/group/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.Group().List(ctx, req.GroupListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		GroupListRes: &model.GroupListRes{
			Items: items,
		},
	}

	return
}
