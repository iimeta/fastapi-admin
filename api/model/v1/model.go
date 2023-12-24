package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建模型接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" tags:"model" method:"post" summary:"新建模型接口"`
	model.ModelCreateReq
}

// 新建模型接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新模型接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" tags:"model" method:"post" summary:"更新模型接口"`
	model.ModelUpdateReq
}

// 更新模型接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除模型接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" tags:"model" method:"post" summary:"删除模型接口"`
	Id     string `json:"id"`
}

// 删除模型接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 模型详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"model" method:"get" summary:"模型详情接口"`
	Id     string `json:"id"`
}

// 模型详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelDetailRes
}

// 模型分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"model" method:"post" summary:"模型分页列表接口"`
	model.ModelPageReq
}

// 模型分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelPageRes
}

// 模型列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" tags:"model" method:"get" summary:"模型列表接口"`
	model.ModelListReq
}

// 模型列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelListRes
}
