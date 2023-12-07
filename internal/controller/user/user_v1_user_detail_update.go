package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) UserDetailUpdate(ctx context.Context, req *v1.UserDetailUpdateReq) (res *v1.UserDetailUpdateRes, err error) {

	err = service.User().ChangeDetail(ctx, req.UserDetailUpdateReq)

	return
}
