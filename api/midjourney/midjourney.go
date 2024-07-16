// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package midjourney

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/midjourney/v1"
)

type IMidjourneyV1 interface {
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
}
