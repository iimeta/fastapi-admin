// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_video

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_video/v1"
)

type ITaskVideoV1 interface {
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	VideoPage(ctx context.Context, req *v1.VideoPageReq) (res *v1.VideoPageRes, err error)
	CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error)
}
