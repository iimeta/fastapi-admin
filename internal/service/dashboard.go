// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/internal/model"
)

type (
	IDashboard interface {
		// 基础数据
		BaseData(ctx context.Context) (dashboard model.Dashboard, err error)
		// 调用数据
		CallData(ctx context.Context, params model.DashboardCallDataReq) ([]*model.CallData, error)
		// 费用
		Expense(ctx context.Context) (*model.Expense, error)
		// 数据TOP
		DataTop(ctx context.Context, params model.DashboardDataTopReq) ([]*model.DataTop, error)
		// 模型占比
		ModelPercent(ctx context.Context, params model.DashboardModelPercentReq) ([]string, []*model.ModelPercent, error)
		// 每秒钟数据
		PerSecond(ctx context.Context, params model.DashboardPerSecondReq) (int, int, error)
		// 每分钟数据
		PerMinute(ctx context.Context, params model.DashboardPerMinuteReq) (int, int, error)
		// 额度预警
		QuotaWarning(ctx context.Context, params model.DashboardQuotaWarningReq) error
	}
)

var (
	localDashboard IDashboard
)

func Dashboard() IDashboard {
	if localDashboard == nil {
		panic("implement not found for interface IDashboard, forgot register?")
	}
	return localDashboard
}

func RegisterDashboard(i IDashboard) {
	localDashboard = i
}
