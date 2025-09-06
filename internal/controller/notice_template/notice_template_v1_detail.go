package notice_template

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/notice_template/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	noticeTemplate, err := service.NoticeTemplate().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		NoticeTemplateDetailRes: &model.NoticeTemplateDetailRes{
			NoticeTemplate: noticeTemplate,
		},
	}

	return
}
