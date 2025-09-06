package common

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/api/common/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) SendEmail(ctx context.Context, req *v1.SendEmailReq) (res *v1.SendEmailRes, err error) {

	if req.Domain == "" {
		req.Domain = g.RequestFromCtx(ctx).GetHost()
	}

	err = service.Common().EmailCode(ctx, req.SendEmailReq)

	return
}
