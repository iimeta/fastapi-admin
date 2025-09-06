package dashboard

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Expense(ctx context.Context, req *v1.ExpenseReq) (res *v1.ExpenseRes, err error) {

	expense, err := service.Dashboard().Expense(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ExpenseRes{
		DashboardExpenseRes: &model.DashboardExpenseRes{
			Expense: expense,
		},
	}

	return
}
