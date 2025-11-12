package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// Midjourney日志详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"user,reseller,admin" tags:"log_midjourney" summary:"Midjourney日志详情接口"`
	Id     string `json:"id"`
}

// Midjourney日志详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.LogMidjourneyDetailRes
}

// Midjourney日志分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"log_midjourney" summary:"Midjourney日志分页列表接口"`
	model.LogMidjourneyPageReq
}

// Midjourney日志分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.LogMidjourneyPageRes
}
