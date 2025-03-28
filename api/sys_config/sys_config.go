// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_config

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/sys_config/v1"
)

type ISysConfigV1 interface {
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	ChangeStatus(ctx context.Context, req *v1.ChangeStatusReq) (res *v1.ChangeStatusRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Reset(ctx context.Context, req *v1.ResetReq) (res *v1.ResetRes, err error)
	Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error)
}
