// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_batch

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_batch/v1"
)

type ITaskBatchV1 interface {
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	BatchPage(ctx context.Context, req *v1.BatchPageReq) (res *v1.BatchPageRes, err error)
	CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error)
}
