package finance

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/finance/v1"
)

func (c *ControllerV1) DealRecordPage(ctx context.Context, req *v1.DealRecordPageReq) (res *v1.DealRecordPageRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	dealRecordPageRes, err := service.Finance().DealRecordPage(ctx, req.FinanceDealRecordPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DealRecordPageRes{
		FinanceDealRecordPageRes: dealRecordPageRes,
	}

	return
}
