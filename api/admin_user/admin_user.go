// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin_user

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/admin_user/v1"
)

type IAdminUserV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	ChangeStatus(ctx context.Context, req *v1.ChangeStatusReq) (res *v1.ChangeStatusRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	GrantQuota(ctx context.Context, req *v1.GrantQuotaReq) (res *v1.GrantQuotaRes, err error)
	Models(ctx context.Context, req *v1.ModelsReq) (res *v1.ModelsRes, err error)
}
