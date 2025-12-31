// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
)

type IStatisticsV1 interface {
	DataUser(ctx context.Context, req *v1.DataUserReq) (res *v1.DataUserRes, err error)
	DataApp(ctx context.Context, req *v1.DataAppReq) (res *v1.DataAppRes, err error)
	DataAppKey(ctx context.Context, req *v1.DataAppKeyReq) (res *v1.DataAppKeyRes, err error)
}
