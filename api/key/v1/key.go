package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建密钥接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" role:"admin" tags:"key" method:"post" summary:"新建密钥接口"`
	model.KeyCreateReq
}

// 新建密钥接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新密钥接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" role:"admin" tags:"key" method:"post" summary:"更新密钥接口"`
	model.KeyUpdateReq
}

// 更新密钥接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改密钥状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" role:"user,admin" tags:"key" method:"post" summary:"更改密钥状态接口"`
	model.KeyChangeStatusReq
}

// 更改密钥状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除密钥接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" role:"user,admin" tags:"key" method:"post" summary:"删除密钥接口"`
	Id     string `json:"id"`
}

// 删除密钥接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 密钥详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" role:"user,admin" tags:"key" method:"get" summary:"密钥详情接口"`
	Id     string `json:"id"`
}

// 密钥详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.KeyDetailRes
}

// 密钥分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" role:"user,admin" tags:"key" method:"post" summary:"密钥分页列表接口"`
	model.KeyPageReq
}

// 密钥分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.KeyPageRes
}

// 密钥列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" role:"admin" tags:"key" method:"get" summary:"密钥列表接口"`
	model.KeyListReq
}

// 密钥列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.KeyListRes
}

// 密钥批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" role:"user,admin" tags:"key" method:"post" summary:"密钥批量操作接口"`
	model.KeyBatchOperateReq
}

// 密钥批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
