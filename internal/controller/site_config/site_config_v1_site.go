package site_config

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/site_config/v1"
)

func (c *ControllerV1) Site(ctx context.Context, req *v1.SiteReq) (res *v1.SiteRes, err error) {

	if req.Domain == "" {
		req.Domain = g.RequestFromCtx(ctx).GetHost()
	}

	siteConfig := service.SiteConfig().Site(ctx, req.SiteConfigDetailReq)

	res = &v1.SiteRes{
		SiteConfigDetailRes: &model.SiteConfigDetailRes{
			SiteConfig: siteConfig,
		},
	}

	return
}
