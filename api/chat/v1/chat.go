package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 聊天详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"chat" method:"get" summary:"聊天详情接口"`
	Id     string `json:"id"`
}

// 聊天详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ChatDetailRes
}

// 聊天分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"chat" method:"post" summary:"聊天分页列表接口"`
	model.ChatPageReq
}

// 聊天分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ChatPageRes
}
