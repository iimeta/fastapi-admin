package notice

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/notice/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	noticePageRes, err := service.Notice().Page(ctx, req.NoticePageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		NoticePageRes: noticePageRes,
	}

	return
}
