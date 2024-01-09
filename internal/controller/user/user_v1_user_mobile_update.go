package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) UserPhoneUpdate(ctx context.Context, req *v1.UserPhoneUpdateReq) (res *v1.UserPhoneUpdateRes, err error) {

	err = service.User().ChangePhone(ctx, req.UserPhoneUpdateReq)

	return
}
