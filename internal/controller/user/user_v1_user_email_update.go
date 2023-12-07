package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) UserEmailUpdate(ctx context.Context, req *v1.UserEmailUpdateReq) (res *v1.UserEmailUpdateRes, err error) {

	err = service.User().ChangeEmail(ctx, req.UserEmailUpdateReq)

	return
}
