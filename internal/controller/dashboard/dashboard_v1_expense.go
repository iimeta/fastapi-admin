package dashboard

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/dashboard/v1"
)

func (c *ControllerV1) Expense(ctx context.Context, req *v1.ExpenseReq) (res *v1.ExpenseRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

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
