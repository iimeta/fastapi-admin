package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) UserPasswordUpdate(ctx context.Context, req *v1.UserPasswordUpdateReq) (res *v1.UserPasswordUpdateRes, err error) {

	err = service.User().ChangePassword(ctx, req.UserPasswordUpdateReq)

	return
}
