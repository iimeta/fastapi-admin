package admin_user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/admin_user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.AdminUser().List(ctx, req.UserListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		UserListRes: &model.UserListRes{
			Items: items,
		},
	}

	return
}
