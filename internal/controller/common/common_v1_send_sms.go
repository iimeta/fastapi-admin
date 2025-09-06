package common

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/api/common/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) SendSms(ctx context.Context, req *v1.SendSmsReq) (res *v1.SendSmsRes, err error) {

	if req.Domain == "" {
		req.Domain = g.RequestFromCtx(ctx).GetHost()
	}

	err = service.Common().SmsCode(ctx, req.SendSmsReq)

	return
}
