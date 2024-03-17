// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/user/v1"
)

type IUserV1 interface {
	Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error)
	ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error)
	ChangeEmail(ctx context.Context, req *v1.ChangeEmailReq) (res *v1.ChangeEmailRes, err error)
	UpdateInfo(ctx context.Context, req *v1.UpdateInfoReq) (res *v1.UpdateInfoRes, err error)
}
