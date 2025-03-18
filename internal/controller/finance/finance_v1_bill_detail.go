package finance

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/finance/v1"
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
