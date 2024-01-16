package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	items, err := service.User().List(ctx, req.UserListReq)
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
