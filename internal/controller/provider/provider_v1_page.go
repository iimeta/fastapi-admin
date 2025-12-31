package provider

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/provider/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	providerPageRes, err := service.Provider().Page(ctx, req.ProviderPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		ProviderPageRes: providerPageRes,
	}

	return
}
