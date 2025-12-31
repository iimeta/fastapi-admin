package admin_user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/admin_user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	userPageRes, err := service.AdminUser().Page(ctx, req.UserPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		UserPageRes: userPageRes,
	}

	return
}
