package auth

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/auth/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {

	err = service.Auth().Logout(ctx)

	return
}
