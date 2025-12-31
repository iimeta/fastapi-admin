// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package finance

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/finance/v1"
)

type IFinanceV1 interface {
	BillDetail(ctx context.Context, req *v1.BillDetailReq) (res *v1.BillDetailRes, err error)
	BillPage(ctx context.Context, req *v1.BillPageReq) (res *v1.BillPageRes, err error)
	BillExport(ctx context.Context, req *v1.BillExportReq) (res *v1.BillExportRes, err error)
	DealRecordPage(ctx context.Context, req *v1.DealRecordPageReq) (res *v1.DealRecordPageRes, err error)
}
