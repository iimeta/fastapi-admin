package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 聊天详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"log/chat" method:"get" summary:"聊天详情接口"`
	Id     string `json:"id"`
}

// 聊天详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ChatDetailRes
}

// 聊天分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"log/chat" method:"post" summary:"聊天分页列表接口"`
	model.ChatPageReq
}

// 聊天分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ChatPageRes
}

// 聊天导出接口请求参数
type ExportReq struct {
	g.Meta `path:"/export" tags:"log/chat" method:"post" summary:"聊天导出接口"`
	model.ChatExportReq
}

// 聊天导出接口响应参数
type ExportRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 聊天批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" tags:"log/chat" method:"post" summary:"聊天批量操作接口"`
	model.ChatBatchOperateReq
}

// 聊天批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
