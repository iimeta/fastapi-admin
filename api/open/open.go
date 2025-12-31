// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package open

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/open/v1"
)

type IOpenV1 interface {
	SysConfig(ctx context.Context, req *v1.SysConfigReq) (res *v1.SysConfigRes, err error)
	SiteConfig(ctx context.Context, req *v1.SiteConfigReq) (res *v1.SiteConfigRes, err error)
	Video(ctx context.Context, req *v1.VideoReq) (res *v1.VideoRes, err error)
	File(ctx context.Context, req *v1.FileReq) (res *v1.FileRes, err error)
}
