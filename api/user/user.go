// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

type IUserV1 interface {
	UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error)
	UserSetting(ctx context.Context, req *v1.UserSettingReq) (res *v1.UserSettingRes, err error)
	UserDetailUpdate(ctx context.Context, req *v1.UserDetailUpdateReq) (res *v1.UserDetailUpdateRes, err error)
	UserPasswordUpdate(ctx context.Context, req *v1.UserPasswordUpdateReq) (res *v1.UserPasswordUpdateRes, err error)
	UserPhoneUpdate(ctx context.Context, req *v1.UserPhoneUpdateReq) (res *v1.UserPhoneUpdateRes, err error)
	UserEmailUpdate(ctx context.Context, req *v1.UserEmailUpdateReq) (res *v1.UserEmailUpdateRes, err error)
}
