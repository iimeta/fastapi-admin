package user

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {

	if gstr.HasPrefix(service.Session().GetToken(ctx), consts.USER_TOKEN_PREFIX) {

		user := service.Session().GetUser(ctx)
		if user == nil {
			return nil, errors.New("Unauthorized")
		}

		res = &v1.UserInfoRes{
			UserInfoRes: &model.UserInfoRes{
				Id:     gconv.String(user.UserId),
				Phone:  user.Phone,
				Email:  user.Email,
				Name:   user.Name,
				Avatar: user.Avatar,
				Gender: user.Gender,
				Role:   consts.USER_CHANNEL,
			},
		}
	} else {

		admin := service.Session().GetAdmin(ctx)
		if admin == nil {
			return nil, errors.New("Unauthorized")
		}

		res = &v1.UserInfoRes{
			UserInfoRes: &model.UserInfoRes{
				Id:     admin.Id,
				Phone:  admin.Phone,
				Email:  admin.Email,
				Name:   admin.Name,
				Avatar: admin.Avatar,
				Gender: admin.Gender,
				Role:   consts.ADMIN_CHANNEL,
			},
		}
	}

	return
}
