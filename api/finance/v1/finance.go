package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 账单明细分页列表接口请求参数
type BillPageReq struct {
	g.Meta `path:"/bill" tags:"finance" method:"post" summary:"账单明细分页列表接口"`
	model.FinanceBillPageReq
}

// 账单明细分页列表接口响应参数
type BillPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.FinanceBillPageRes
}

// 交易记录分页列表接口请求参数
type DealRecordPageReq struct {
	g.Meta `path:"/deal/record" tags:"finance" method:"post" summary:"交易记录分页列表接口"`
	model.FinanceDealRecordPageReq
}

// 交易记录分页列表接口响应参数
type DealRecordPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.FinanceDealRecordPageRes
}
