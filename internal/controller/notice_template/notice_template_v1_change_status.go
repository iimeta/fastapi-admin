package notice_template

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/notice_template/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ChangeStatus(ctx context.Context, req *v1.ChangeStatusReq) (res *v1.ChangeStatusRes, err error) {

	err = service.NoticeTemplate().ChangeStatus(ctx, req.NoticeTemplateChangeStatusReq)

	return
}
