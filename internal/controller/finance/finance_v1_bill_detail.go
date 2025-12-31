package finance

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/finance/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) BillDetail(ctx context.Context, req *v1.BillDetailReq) (res *v1.BillDetailRes, err error) {

	statisticsUser, err := service.Finance().BillDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.BillDetailRes{
		FinanceBillDetailRes: &model.FinanceBillDetailRes{
			StatisticsUser: statisticsUser,
		},
	}

	return
}
