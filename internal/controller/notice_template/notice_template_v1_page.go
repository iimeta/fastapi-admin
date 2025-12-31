package notice_template

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/notice_template/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	noticeTemplatePageRes, err := service.NoticeTemplate().Page(ctx, req.NoticeTemplatePageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		NoticeTemplatePageRes: noticeTemplatePageRes,
	}

	return
}
