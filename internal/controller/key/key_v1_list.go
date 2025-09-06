package key

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/key/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.Key().List(ctx, req.KeyListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		KeyListRes: &model.KeyListRes{
			Items: items,
		},
	}

	return
}
