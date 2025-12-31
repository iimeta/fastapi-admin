// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_file

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_file/v1"
)

type ITaskFileV1 interface {
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	FilePage(ctx context.Context, req *v1.FilePageReq) (res *v1.FilePageRes, err error)
	CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error)
}
