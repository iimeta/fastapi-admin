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
	Expense(ctx context.Context, req *v1.ExpenseReq) (res *v1.ExpenseRes, err error)
}
