// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package finance

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/finance/v1"
)

type IFinanceV1 interface {
	BillPage(ctx context.Context, req *v1.BillPageReq) (res *v1.BillPageRes, err error)
	DealRecordPage(ctx context.Context, req *v1.DealRecordPageReq) (res *v1.DealRecordPageRes, err error)
	BillExport(ctx context.Context, req *v1.BillExportReq) (res *v1.BillExportRes, err error)
}
