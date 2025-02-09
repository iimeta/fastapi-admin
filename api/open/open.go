// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package open

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/open/v1"
)

type IOpenV1 interface {
	SysConfig(ctx context.Context, req *v1.SysConfigReq) (res *v1.SysConfigRes, err error)
	SiteConfig(ctx context.Context, req *v1.SiteConfigReq) (res *v1.SiteConfigRes, err error)
}
