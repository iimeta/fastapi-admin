// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_menu

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/sys_menu/v1"
)

type ISysMenuV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error)
	Forget(ctx context.Context, req *v1.ForgetReq) (res *v1.ForgetRes, err error)
}
