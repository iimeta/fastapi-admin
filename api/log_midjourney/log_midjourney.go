// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package log_midjourney

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/log_midjourney/v1"
)

type ILogMidjourneyV1 interface {
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
}
