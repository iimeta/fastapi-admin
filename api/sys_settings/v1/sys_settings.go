package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建设置接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" tags:"settings" method:"post" summary:"新建设置接口"`
	model.SettingsCreateReq
}

// 新建设置接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新设置接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" tags:"settings" method:"post" summary:"更新设置接口"`
	model.SettingsUpdateReq
}

// 更新设置接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除设置接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" tags:"settings" method:"get" summary:"删除设置接口"`
	Id     string `json:"id"`
}

// 删除设置接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 设置详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"settings" method:"get" summary:"设置详情接口"`
	Id     string `json:"id"`
}

// 设置详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SettingsDetailRes
}

// 设置分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"settings" method:"get" summary:"设置分页列表接口"`
	model.SettingsPageReq
}

// 设置分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SettingsPageRes
}
