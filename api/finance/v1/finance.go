package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 账单明细详情接口请求参数
type BillDetailReq struct {
	g.Meta `path:"/bill/detail" role:"user,admin" tags:"finance" method:"get" summary:"账单明细详情接口"`
	Id     string `json:"id"`
}

// 账单明细详情接口响应参数
type BillDetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.FinanceBillDetailRes
}

// 账单明细分页列表接口请求参数
type BillPageReq struct {
	g.Meta `path:"/bill/page" role:"user,admin" tags:"finance" method:"post" summary:"账单明细分页列表接口"`
	model.FinanceBillPageReq
}

// 账单明细分页列表接口响应参数
type BillPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.FinanceBillPageRes
}

// 账单明细导出接口请求参数
type BillExportReq struct {
	g.Meta `path:"/bill/export" role:"user,admin" tags:"finance" method:"post" summary:"账单明细导出接口"`
	model.FinanceBillExportReq
}

// 账单明细导出接口响应参数
type BillExportRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 交易记录分页列表接口请求参数
type DealRecordPageReq struct {
	g.Meta `path:"/deal/record/page" role:"user,admin" tags:"finance" method:"post" summary:"交易记录分页列表接口"`
	model.FinanceDealRecordPageReq
}

// 交易记录分页列表接口响应参数
type DealRecordPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.FinanceDealRecordPageRes
}
