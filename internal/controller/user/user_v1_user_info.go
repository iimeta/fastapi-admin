package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {

	userInfoRes, err := service.User().Info(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.UserInfoRes{
		UserInfoRes: userInfoRes,
	}

	return
}
