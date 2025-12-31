package site_config

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/site_config/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	err = service.SiteConfig().Create(ctx, req.SiteConfigCreateReq)

	return
}
