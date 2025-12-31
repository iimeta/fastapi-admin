// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package log_text

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/log_text/v1"
)

type ILogTextV1 interface {
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	Export(ctx context.Context, req *v1.ExportReq) (res *v1.ExportRes, err error)
	BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error)
	CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error)
}
