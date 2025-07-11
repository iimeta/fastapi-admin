package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建通知模板接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"admin" tags:"notice" summary:"新建通知模板接口"`
	model.NoticeTemplateCreateReq
}

// 新建通知模板接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新通知模板接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"admin" tags:"notice" summary:"更新通知模板接口"`
	model.NoticeTemplateUpdateReq
}

// 更新通知模板接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改通知模板公开状态接口请求参数
type ChangePublicReq struct {
	g.Meta `path:"/change/public" method:"post" auth:"true" role:"admin" tags:"group" summary:"更改通知模板公开状态接口"`
	model.NoticeTemplateChangePublicReq
}

// 更改通知模板公开状态接口响应参数
type ChangePublicRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改通知模板状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"admin" tags:"group" summary:"更改通知模板状态接口"`
	model.NoticeTemplateChangeStatusReq
}

// 更改通知模板状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除通知模板接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"admin" tags:"notice" summary:"删除通知模板接口"`
	Id     string `json:"id"`
}

// 删除通知模板接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 通知模板详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"admin" tags:"notice" summary:"通知模板详情接口"`
	Id     string `json:"id"`
}

// 通知模板详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticeTemplateDetailRes
}

// 通知模板分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"admin" tags:"notice" summary:"通知模板分页列表接口"`
	model.NoticeTemplatePageReq
}

// 通知模板分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticeTemplatePageRes
}

// 通知模板列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" method:"get" auth:"true" role:"admin" tags:"notice" summary:"通知模板列表接口"`
	model.NoticeTemplateListReq
}

// 通知模板列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticeTemplateListRes
}

// 通知模板批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"admin" tags:"notice" summary:"通知模板批量操作接口"`
	model.NoticeTemplateBatchOperateReq
}

// 通知模板批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
