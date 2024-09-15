package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/util"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {

	if service.Session().IsUserRole(ctx) {

		user := service.Session().GetUser(ctx)
		if user == nil {
			return nil, errors.New("Unauthorized")
		}

		res = &v1.InfoRes{
			UserInfoRes: &model.UserInfoRes{
				UserId:    user.UserId,
				Name:      user.Name,
				Avatar:    user.Avatar,
				Email:     user.Email,
				Phone:     user.Phone,
				Account:   user.Account,
				Role:      consts.USER_CHANNEL,
				RPS:       user.RPS,
				RPM:       user.RPM,
				RPD:       user.RPD,
				TPS:       user.TPS,
				TPM:       user.TPM,
				TPD:       user.TPD,
				IPS:       user.IPS,
				IPM:       user.IPM,
				IPD:       user.IPD,
				CreatedAt: user.CreatedAt,
			},
		}

	} else {

		admin := service.Session().GetAdmin(ctx)
		if admin == nil {
			return nil, errors.New("Unauthorized")
		}

		res = &v1.InfoRes{
			UserInfoRes: &model.UserInfoRes{
				UserId:    admin.UserId,
				Name:      admin.Name,
				Avatar:    admin.Avatar,
				Email:     admin.Email,
				Phone:     admin.Phone,
				Account:   admin.Account,
				Role:      consts.ADMIN_CHANNEL,
				CreatedAt: util.FormatDateTime(admin.CreatedAt),
			},
		}
	}

	return
}
