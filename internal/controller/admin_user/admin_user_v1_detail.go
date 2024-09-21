package admin_user

import (
	"context"
	"github.com/iimeta/fastapi-admin/api/admin_user/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	user, err := service.AdminUser().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		UserDetailRes: &model.UserDetailRes{
			User: user,
		},
	}

	return
}
