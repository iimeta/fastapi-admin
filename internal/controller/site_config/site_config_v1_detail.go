package site_config

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/site_config/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	siteConfig, err := service.SiteConfig().Detail(ctx, req.SiteConfigDetailReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		SiteConfigDetailRes: &model.SiteConfigDetailRes{
			SiteConfig: siteConfig,
		},
	}

	return
}
