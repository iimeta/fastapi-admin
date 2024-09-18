package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建管理员接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" role:"admin" tags:"admin" method:"post" summary:"新建管理员接口"`
	model.SysAdminCreateReq
}

// 新建管理员接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新管理员接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" role:"admin" tags:"admin" method:"post" summary:"更新管理员接口"`
	model.SysAdminUpdateReq
}

// 更新管理员接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除管理员接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" role:"admin" tags:"admin" method:"get" summary:"删除管理员接口"`
	Id     string `json:"id"`
}

// 删除管理员接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 管理员详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" role:"admin" tags:"admin" method:"get" summary:"管理员详情接口"`
	Id     string `json:"id"`
}

// 管理员详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SysAdminDetailRes
}

// 管理员分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" role:"admin" tags:"admin" method:"get" summary:"管理员分页列表接口"`
	model.SysAdminPageReq
}

// 管理员分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SysAdminPageRes
}
