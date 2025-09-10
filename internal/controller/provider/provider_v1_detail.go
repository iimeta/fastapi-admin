package provider

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/provider/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	provider, err := service.Provider().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		ProviderDetailRes: &model.ProviderDetailRes{
			Provider: provider,
		},
	}

	return
}
