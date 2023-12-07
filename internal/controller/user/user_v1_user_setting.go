package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) UserSetting(ctx context.Context, req *v1.UserSettingReq) (res *v1.UserSettingRes, err error) {

	userSettingRes, err := service.User().Setting(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.UserSettingRes{
		UserSettingRes: userSettingRes,
	}

	return
}
