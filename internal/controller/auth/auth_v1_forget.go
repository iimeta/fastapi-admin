package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/auth/v1"
)

func (c *ControllerV1) Forget(ctx context.Context, req *v1.ForgetReq) (res *v1.ForgetRes, err error) {

	if req.Domain == "" {
		req.Domain = g.RequestFromCtx(ctx).GetHost()
	}

	err = service.Auth().Forget(ctx, req.ForgetReq)

	return
}
