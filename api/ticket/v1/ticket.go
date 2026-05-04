package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

// 创建工单接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"user,reseller" tags:"ticket" summary:"创建工单接口"`
	model.TicketCreateReq
}

// 创建工单接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 回复工单接口请求参数
type ReplyReq struct {
	g.Meta `path:"/reply" method:"post" auth:"true" role:"user,reseller,admin" tags:"ticket" summary:"回复工单接口"`
	model.TicketReplyReq
}

// 回复工单接口响应参数
type ReplyRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 工单详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"user,reseller,admin" tags:"ticket" summary:"工单详情接口"`
	Id     string `json:"id"`
}

// 工单详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TicketDetailRes
}

// 工单分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"ticket" summary:"工单分页列表接口"`
	model.TicketPageReq
}

// 工单分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TicketPageRes
}

// 更新工单状态接口请求参数
// 状态: 1-待回复, 2-待处理, 3-处理中, 4-已回复, 5-已解决, 6-已关闭
type UpdateStatusReq struct {
	g.Meta `path:"/status/update" method:"post" auth:"true" role:"user,reseller,admin" tags:"ticket" summary:"更新工单状态接口"`
	model.TicketUpdateStatusReq
}

// 更新工单状态接口响应参数
type UpdateStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 关闭工单接口请求参数
type CloseReq struct {
	g.Meta `path:"/close" method:"post" auth:"true" role:"user,reseller,admin" tags:"ticket" summary:"关闭工单接口"`
	model.TicketCloseReq
}

// 关闭工单接口响应参数
type CloseRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 分配工单接口请求参数
type AssignReq struct {
	g.Meta `path:"/assign" method:"post" auth:"true" role:"admin" tags:"ticket" summary:"分配工单接口"`
	model.TicketAssignReq
}

// 分配工单接口响应参数
type AssignRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 批量删除工单接口请求参数
type BatchDeleteReq struct {
	g.Meta `path:"/batch/delete" method:"post" auth:"true" role:"admin" tags:"ticket" summary:"批量删除工单接口"`
	model.TicketBatchDeleteReq
}

// 批量删除工单接口响应参数
type BatchDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
