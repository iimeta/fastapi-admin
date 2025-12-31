package open

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/api/open/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) SiteConfig(ctx context.Context, req *v1.SiteConfigReq) (res *v1.SiteConfigRes, err error) {

	if req.Domain == "" {
		req.Domain = g.RequestFromCtx(ctx).GetHost()
	}

	siteConfig := service.SiteConfig().Site(ctx, req.SiteConfigDetailReq)

	res = &v1.SiteConfigRes{
		SiteConfigDetailRes: &model.SiteConfigDetailRes{
			SiteConfig: siteConfig,
		},
	}

	return
}
