package admin_user

import (
	"context"
	"github.com/iimeta/fastapi-admin/api/admin_user/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Models(ctx context.Context, req *v1.ModelsReq) (res *v1.ModelsRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.AdminUser().Models(ctx, req.UserModelsReq)

	return
}
