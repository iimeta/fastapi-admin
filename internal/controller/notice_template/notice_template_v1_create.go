package notice_template

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/notice_template/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	_, err = service.NoticeTemplate().Create(ctx, req.NoticeTemplateCreateReq)

	return
}
