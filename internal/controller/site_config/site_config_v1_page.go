package site_config

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/site_config/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	siteConfigPageRes, err := service.SiteConfig().Page(ctx, req.SiteConfigPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		SiteConfigPageRes: siteConfigPageRes,
	}

	return
}
