// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app_key

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/app_key/v1"
)

type IAppKeyV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Config(ctx context.Context, req *v1.ConfigReq) (res *v1.ConfigRes, err error)
	ChangeStatus(ctx context.Context, req *v1.ChangeStatusReq) (res *v1.ChangeStatusRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error)
	Export(ctx context.Context, req *v1.ExportReq) (res *v1.ExportRes, err error)
}
