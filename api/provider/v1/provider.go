package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建提供商接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"admin" tags:"provider" summary:"新建提供商接口"`
	model.ProviderCreateReq
}

// 新建提供商接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新提供商接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"admin" tags:"provider" summary:"更新提供商接口"`
	model.ProviderUpdateReq
}

// 更新提供商接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改提供商状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"admin" tags:"provider" summary:"更改提供商状态接口"`
	model.ProviderChangeStatusReq
}

// 更改提供商状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改提供商公开状态接口请求参数
type ChangePublicReq struct {
	g.Meta `path:"/change/public" method:"post" auth:"true" role:"admin" tags:"provider" summary:"更改提供商公开状态接口"`
	model.ProviderChangePublicReq
}

// 更改提供商公开状态接口响应参数
type ChangePublicRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除提供商接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"admin" tags:"provider" summary:"删除提供商接口"`
	Id     string `json:"id"`
}

// 删除提供商接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 提供商详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"admin" tags:"provider" summary:"提供商详情接口"`
	Id     string `json:"id"`
}

// 提供商详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ProviderDetailRes
}

// 提供商分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"admin" tags:"provider" summary:"提供商分页列表接口"`
	model.ProviderPageReq
}

// 提供商分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ProviderPageRes
}

// 提供商列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" method:"get" auth:"true" role:"user,reseller,admin" tags:"provider" summary:"提供商列表接口"`
	model.ProviderListReq
}

// 提供商列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ProviderListRes
}

// 提供商批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"admin" tags:"provider" summary:"提供商批量操作接口"`
	model.ProviderBatchOperateReq
}

// 提供商批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
