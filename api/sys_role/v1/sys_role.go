package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建角色接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" tags:"role" method:"post" summary:"新建角色接口"`
	model.RoleCreateReq
}

// 新建角色接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新角色接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" tags:"role" method:"post" summary:"更新角色接口"`
	model.RoleUpdateReq
}

// 更新角色接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除角色接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" tags:"role" method:"get" summary:"删除角色接口"`
	Id     string `json:"id"`
}

// 删除角色接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 角色详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"role" method:"get" summary:"角色详情接口"`
	Id     string `json:"id"`
}

// 角色详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.RoleDetailRes
}

// 角色分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"role" method:"get" summary:"角色分页列表接口"`
	model.RolePageReq
}

// 角色分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.RolePageRes
}
