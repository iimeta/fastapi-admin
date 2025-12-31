package site_config

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/site_config/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {

	err = service.SiteConfig().Delete(ctx, req.Id)

	return
}
