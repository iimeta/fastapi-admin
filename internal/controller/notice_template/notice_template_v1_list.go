package notice_template

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/notice_template/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.NoticeTemplate().List(ctx, req.NoticeTemplateListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		NoticeTemplateListRes: &model.NoticeTemplateListRes{
			Items: items,
		},
	}

	return
}
