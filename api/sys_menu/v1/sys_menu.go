package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建菜单接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" tags:"menu" method:"post" summary:"新建菜单接口"`
	model.SysMenuCreateReq
}

// 新建菜单接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新菜单接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" tags:"menu" method:"post" summary:"更新菜单接口"`
	model.SysMenuUpdateReq
}

// 更新菜单接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除菜单接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" tags:"menu" method:"get" summary:"删除菜单接口"`
	Id     string `json:"id"`
}

// 删除菜单接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 菜单详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"menu" method:"get" summary:"菜单详情接口"`
	Id     string `json:"id"`
}

// 菜单详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SysMenuDetailRes
}

// 菜单分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"menu" method:"get" summary:"菜单分页列表接口"`
	model.SysMenuPageReq
}

// 菜单分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SysMenuPageRes
}
