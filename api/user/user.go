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
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	ChangeStatus(ctx context.Context, req *v1.ChangeStatusReq) (res *v1.ChangeStatusRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	GrantQuota(ctx context.Context, req *v1.GrantQuotaReq) (res *v1.GrantQuotaRes, err error)
}
