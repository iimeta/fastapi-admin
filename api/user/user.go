// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

type IUserV1 interface {
	UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error)
	UserSetting(ctx context.Context, req *v1.UserSettingReq) (res *v1.UserSettingRes, err error)
	UserDetailUpdate(ctx context.Context, req *v1.UserDetailUpdateReq) (res *v1.UserDetailUpdateRes, err error)
	UserPasswordUpdate(ctx context.Context, req *v1.UserPasswordUpdateReq) (res *v1.UserPasswordUpdateRes, err error)
	UserMobileUpdate(ctx context.Context, req *v1.UserMobileUpdateReq) (res *v1.UserMobileUpdateRes, err error)
	UserEmailUpdate(ctx context.Context, req *v1.UserEmailUpdateReq) (res *v1.UserEmailUpdateRes, err error)
}
