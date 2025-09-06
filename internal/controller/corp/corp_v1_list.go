package corp

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/corp/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.Corp().List(ctx, req.CorpListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		CorpListRes: &model.CorpListRes{
			Items: items,
		},
	}

	return
}
