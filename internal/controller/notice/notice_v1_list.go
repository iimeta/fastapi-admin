package notice

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/notice/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.Notice().List(ctx, req.NoticeListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		NoticeListRes: &model.NoticeListRes{
			Items: items,
		},
	}

	return
}
