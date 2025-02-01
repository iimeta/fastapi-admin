package common

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/common/v1"
)

func (c *ControllerV1) SendEmail(ctx context.Context, req *v1.SendEmailReq) (res *v1.SendEmailRes, err error) {

	err = service.Common().EmailCode(ctx, req.SendEmailReq)

	return
}
