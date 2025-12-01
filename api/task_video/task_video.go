// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_video

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/task_video/v1"
)

type ITaskVideoV1 interface {
	VideoPage(ctx context.Context, req *v1.VideoPageReq) (res *v1.VideoPageRes, err error)
}
