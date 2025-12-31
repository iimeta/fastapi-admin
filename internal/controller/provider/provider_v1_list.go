package provider

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/provider/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.Provider().List(ctx, req.ProviderListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		ProviderListRes: &model.ProviderListRes{
			Items: items,
		},
	}

	return
}
