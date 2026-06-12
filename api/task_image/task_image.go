// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_image

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_image/v1"
)

type ITaskImageV1 interface {
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	ImagePage(ctx context.Context, req *v1.ImagePageReq) (res *v1.ImagePageRes, err error)
	CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error)
	Regenerate(ctx context.Context, req *v1.RegenerateReq) (res *v1.RegenerateRes, err error)
	BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error)
}
