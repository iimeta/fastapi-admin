package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建模型代理接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" auth:"true" role:"admin" tags:"model_agent" method:"post" summary:"新建模型代理接口"`
	model.ModelAgentCreateReq
}

// 新建模型代理接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新模型代理接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" auth:"true" role:"admin" tags:"model_agent" method:"post" summary:"更新模型代理接口"`
	model.ModelAgentUpdateReq
}

// 更新模型代理接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改模型代理状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" auth:"true" role:"admin" tags:"model_agent" method:"post" summary:"更改模型代理状态接口"`
	model.ModelAgentChangeStatusReq
}

// 更改模型代理状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除模型代理接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" auth:"true" role:"admin" tags:"model_agent" method:"post" summary:"删除模型代理接口"`
	Id     string `json:"id"`
}

// 删除模型代理接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 模型代理详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" auth:"true" role:"admin" tags:"model_agent" method:"get" summary:"模型代理详情接口"`
	Id     string `json:"id"`
}

// 模型代理详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelAgentDetailRes
}

// 模型代理分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" auth:"true" role:"admin" tags:"model_agent" method:"post" summary:"模型代理分页列表接口"`
	model.ModelAgentPageReq
}

// 模型代理分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelAgentPageRes
}

// 模型代理列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" auth:"true" role:"admin" tags:"model_agent" method:"get" summary:"模型代理列表接口"`
	model.ModelAgentListReq
}

// 模型代理列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelAgentListRes
}

// 模型代理批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" auth:"true" role:"admin" tags:"model_agent" method:"post" summary:"模型代理批量操作接口"`
	model.ModelAgentBatchOperateReq
}

// 模型代理批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
