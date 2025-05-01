package admin_reseller

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/admin_reseller/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	userPageRes, err := service.AdminReseller().Page(ctx, req.ResellerPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		ResellerPageRes: userPageRes,
	}

	return
}
