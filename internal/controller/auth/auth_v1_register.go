package auth

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/auth/v1"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {

	err = service.Auth().Register(ctx, req.RegisterReq)

	return
}
