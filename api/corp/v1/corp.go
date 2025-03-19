package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建公司接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" auth:"true" role:"admin" tags:"corp" method:"post" summary:"新建公司接口"`
	model.CorpCreateReq
}

// 新建公司接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新公司接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" auth:"true" role:"admin" tags:"corp" method:"post" summary:"更新公司接口"`
	model.CorpUpdateReq
}

// 更新公司接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改公司状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" auth:"true" role:"admin" tags:"corp" method:"post" summary:"更改公司状态接口"`
	model.CorpChangeStatusReq
}

// 更改公司状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改公司公开状态接口请求参数
type ChangePublicReq struct {
	g.Meta `path:"/change/public" auth:"true" role:"admin" tags:"corp" method:"post" summary:"更改公司公开状态接口"`
	model.CorpChangePublicReq
}

// 更改公司公开状态接口响应参数
type ChangePublicRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除公司接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" auth:"true" role:"admin" tags:"corp" method:"post" summary:"删除公司接口"`
	Id     string `json:"id"`
}

// 删除公司接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 公司详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" auth:"true" role:"admin" tags:"corp" method:"get" summary:"公司详情接口"`
	Id     string `json:"id"`
}

// 公司详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.CorpDetailRes
}

// 公司分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" auth:"true" role:"admin" tags:"corp" method:"post" summary:"公司分页列表接口"`
	model.CorpPageReq
}

// 公司分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.CorpPageRes
}

// 公司列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" auth:"true" role:"user,admin" tags:"corp" method:"get" summary:"公司列表接口"`
	model.CorpListReq
}

// 公司列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.CorpListRes
}

// 公司批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" auth:"true" role:"admin" tags:"corp" method:"post" summary:"公司批量操作接口"`
	model.CorpBatchOperateReq
}

// 公司批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
