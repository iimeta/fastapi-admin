package auth

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/auth/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {

	loginRes, err := service.Auth().Login(ctx, req.LoginReq)
	if err != nil {
		return nil, err
	}

	res = &v1.LoginRes{
		LoginRes: loginRes,
	}

	return
}
