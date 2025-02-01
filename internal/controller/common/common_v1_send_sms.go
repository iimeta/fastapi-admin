package common

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/common/v1"
)

func (c *ControllerV1) SendSms(ctx context.Context, req *v1.SendSmsReq) (res *v1.SendSmsRes, err error) {

	err = service.Common().SmsCode(ctx, req.SendSmsReq)

	return
}
