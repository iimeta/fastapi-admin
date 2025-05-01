package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建分组接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"admin" tags:"group" summary:"新建分组接口"`
	model.GroupCreateReq
}

// 新建分组接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新分组接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"admin" tags:"group" summary:"更新分组接口"`
	model.GroupUpdateReq
}

// 更新分组接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改过期时间接口请求参数
type ChangeExpireReq struct {
	g.Meta `path:"/change/expire" method:"post" auth:"true" role:"admin" tags:"group" summary:"更改过期时间接口"`
	model.GroupChangeExpireReq
}

// 更改过期时间接口响应参数
type ChangeExpireRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改分组公开状态接口请求参数
type ChangePublicReq struct {
	g.Meta `path:"/change/public" method:"post" auth:"true" role:"admin" tags:"group" summary:"更改分组公开状态接口"`
	model.GroupChangePublicReq
}

// 更改分组公开状态接口响应参数
type ChangePublicRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改分组状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"admin" tags:"group" summary:"更改分组状态接口"`
	model.GroupChangeStatusReq
}

// 更改分组状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除分组接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"admin" tags:"group" summary:"删除分组接口"`
	Id     string `json:"id"`
}

// 删除分组接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 分组详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"admin" tags:"group" summary:"分组详情接口"`
	Id     string `json:"id"`
}

// 分组详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupDetailRes
}

// 分组分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"reseller,user,admin" tags:"group" summary:"分组分页列表接口"`
	model.GroupPageReq
}

// 分组分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupPageRes
}

// 分组列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" method:"get" auth:"true" role:"reseller,user,admin" tags:"group" summary:"分组列表接口"`
	model.GroupListReq
}

// 分组列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupListRes
}

// 分组批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"admin" tags:"group" summary:"分组批量操作接口"`
	model.GroupBatchOperateReq
}

// 分组批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
