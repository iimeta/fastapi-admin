// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
)

type IDashboardV1 interface {
	BaseData(ctx context.Context, req *v1.BaseDataReq) (res *v1.BaseDataRes, err error)
	CallData(ctx context.Context, req *v1.CallDataReq) (res *v1.CallDataRes, err error)
	Expense(ctx context.Context, req *v1.ExpenseReq) (res *v1.ExpenseRes, err error)
	DataTop(ctx context.Context, req *v1.DataTopReq) (res *v1.DataTopRes, err error)
	ModelPercent(ctx context.Context, req *v1.ModelPercentReq) (res *v1.ModelPercentRes, err error)
	PerSecond(ctx context.Context, req *v1.PerSecondReq) (res *v1.PerSecondRes, err error)
	PerMinute(ctx context.Context, req *v1.PerMinuteReq) (res *v1.PerMinuteRes, err error)
	WarningConfig(ctx context.Context, req *v1.WarningConfigReq) (res *v1.WarningConfigRes, err error)
}
