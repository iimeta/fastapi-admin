package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Privacy(ctx context.Context, req *v1.PrivacyReq) (res *v1.PrivacyRes, err error) {

	privacy, err := service.User().Privacy(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.PrivacyRes{
		UserPrivacyRes: privacy,
	}

	return
}
