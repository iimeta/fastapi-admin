// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package finance

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/finance/v1"
)

type IFinanceV1 interface {
	DealRecordPage(ctx context.Context, req *v1.DealRecordPageReq) (res *v1.DealRecordPageRes, err error)
}
