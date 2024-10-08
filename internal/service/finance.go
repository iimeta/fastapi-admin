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
	IFinance interface {
		// 账单明细详情
		BillDetail(ctx context.Context, id string) (*model.StatisticsUser, error)
		// 账单明细分页列表
		BillPage(ctx context.Context, params model.FinanceBillPageReq) (*model.FinanceBillPageRes, error)
		// 账单明细导出
		BillExport(ctx context.Context, params model.FinanceBillExportReq) (string, error)
		// 交易记录分页列表
		DealRecordPage(ctx context.Context, params model.FinanceDealRecordPageReq) (*model.FinanceDealRecordPageRes, error)
	}
)

var (
	localFinance IFinance
)

func Finance() IFinance {
	if localFinance == nil {
		panic("implement not found for interface IFinance, forgot register?")
	}
	return localFinance
}

func RegisterFinance(i IFinance) {
	localFinance = i
}
