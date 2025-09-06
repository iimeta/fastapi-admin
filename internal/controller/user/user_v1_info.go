package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/user/v1"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/util"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {

	if service.Session().IsResellerRole(ctx) {

		reseller := service.Session().GetReseller(ctx)
		if reseller == nil {
			return nil, errors.New("Unauthorized")
		}

		res = &v1.InfoRes{
			UserInfoRes: &model.UserInfoRes{
				UserId:    reseller.UserId,
				Name:      reseller.Name,
				Avatar:    reseller.Avatar,
				Email:     reseller.Email,
				Phone:     reseller.Phone,
				Account:   reseller.Account,
				Role:      consts.RESELLER_CHANNEL,
				CreatedAt: reseller.CreatedAt,
			},
		}

	} else if service.Session().IsUserRole(ctx) {

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
				CreatedAt: user.CreatedAt,
				Rid:       user.Rid,
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
