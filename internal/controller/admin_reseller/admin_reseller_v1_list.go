package admin_reseller

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/admin_reseller/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.AdminReseller().List(ctx, req.ResellerListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		ResellerListRes: &model.ResellerListRes{
			Items: items,
		},
	}

	return
}
