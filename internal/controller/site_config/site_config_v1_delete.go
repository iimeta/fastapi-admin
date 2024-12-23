package site_config

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/site_config/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.SiteConfig().Delete(ctx, req.Id)

	return
}
