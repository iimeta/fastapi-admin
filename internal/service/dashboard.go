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
		BaseData(ctx context.Context) (dashboard *model.Dashboard, err error)
		// 调用数据
		CallData(ctx context.Context) ([]*model.CallData, error)
		// 费用
		Expense(ctx context.Context) (*model.Expense, error)
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
