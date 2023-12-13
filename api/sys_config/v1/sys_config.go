package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建配置接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" tags:"config" method:"post" summary:"新建配置接口"`
	model.ConfigCreateReq
}

// 新建配置接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新配置接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" tags:"config" method:"post" summary:"更新配置接口"`
	model.ConfigUpdateReq
}

// 更新配置接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除配置接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" tags:"config" method:"get" summary:"删除配置接口"`
	Id     string `json:"id"`
}

// 删除配置接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 配置详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"config" method:"get" summary:"配置详情接口"`
	Id     string `json:"id"`
}

// 配置详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ConfigDetailRes
}

// 配置分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"config" method:"get" summary:"配置分页列表接口"`
	model.ConfigPageReq
}

// 配置分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ConfigPageRes
}
