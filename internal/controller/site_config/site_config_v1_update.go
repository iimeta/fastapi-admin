package site_config

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/site_config/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	err = service.SiteConfig().Update(ctx, req.SiteConfigUpdateReq)

	return
}
