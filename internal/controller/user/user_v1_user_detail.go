package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error) {

	userDetailRes, err := service.User().Detail(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.UserDetailRes{
		UserDetailRes: userDetailRes,
	}

	return
}
