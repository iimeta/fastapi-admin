package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) UpdatePrivacy(ctx context.Context, req *v1.UpdatePrivacyReq) (res *v1.UpdatePrivacyRes, err error) {

	if err = service.User().UpdatePrivacy(ctx, req.UserPrivacyReq); err != nil {
		return nil, err
	}

	res = &v1.UpdatePrivacyRes{}

	return
}
