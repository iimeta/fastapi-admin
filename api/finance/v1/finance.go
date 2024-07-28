package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 交易记录分页列表接口请求参数
type DealRecordPageReq struct {
	g.Meta `path:"/deal/record" tags:"chat" method:"post" summary:"交易记录分页列表接口"`
	model.FinanceDealRecordPageReq
}

// 交易记录分页列表接口响应参数
type DealRecordPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.FinanceDealRecordPageRes
}
