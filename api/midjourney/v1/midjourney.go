package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// Midjourney详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"reseller,user,admin" tags:"chat" summary:"Midjourney详情接口"`
	Id     string `json:"id"`
}

// Midjourney详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.MidjourneyDetailRes
}

// Midjourney分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"reseller,user,admin" tags:"chat" summary:"Midjourney分页列表接口"`
	model.MidjourneyPageReq
}

// Midjourney分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.MidjourneyPageRes
}
