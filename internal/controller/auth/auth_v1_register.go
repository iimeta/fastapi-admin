package auth

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/api/auth/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {

	if req.Domain == "" {
		req.Domain = g.RequestFromCtx(ctx).GetHost()
	}

	err = service.Auth().Register(ctx, req.RegisterReq)

	return
}
