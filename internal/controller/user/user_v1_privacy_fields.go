package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/user/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) PrivacyFields(ctx context.Context, req *v1.PrivacyFieldsReq) (res *v1.PrivacyFieldsRes, err error) {

	res = &v1.PrivacyFieldsRes{
		Items: service.User().PrivacyFields(ctx),
	}

	return
}
