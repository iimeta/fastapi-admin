package notice

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/notice/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	notice, err := service.Notice().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		NoticeDetailRes: &model.NoticeDetailRes{
			Notice: notice,
		},
	}

	return
}
