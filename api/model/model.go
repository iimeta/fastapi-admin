// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model/v1"
)

type IModelV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	ChangeStatus(ctx context.Context, req *v1.ChangeStatusReq) (res *v1.ChangeStatusRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error)
	Permissions(ctx context.Context, req *v1.PermissionsReq) (res *v1.PermissionsRes, err error)
	InitSync(ctx context.Context, req *v1.InitSyncReq) (res *v1.InitSyncRes, err error)
	Tree(ctx context.Context, req *v1.TreeReq) (res *v1.TreeRes, err error)
}
