package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	userPageRes, err := service.User().Page(ctx, req.UserPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		UserPageRes: userPageRes,
	}

	return
}
