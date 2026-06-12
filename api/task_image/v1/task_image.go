package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

// 绘图任务详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"user,reseller,admin" tags:"task_image" summary:"绘图任务详情接口"`
	Id     string `json:"id"`
}

// 绘图任务详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TaskImageDetailRes
}

// 绘图任务分页列表接口请求参数
type ImagePageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"task_image" summary:"绘图任务分页列表接口"`
	model.TaskImagePageReq
}

// 绘图任务分页列表接口响应参数
type ImagePageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TaskImagePageRes
}

// 绘图任务详情复制字段值接口请求参数
type CopyFieldReq struct {
	g.Meta `path:"/copy/field" method:"post" auth:"true" role:"user,reseller,admin" tags:"task_image" summary:"绘图任务详情复制字段值详情接口"`
	model.TaskImageCopyFieldReq
}

// 绘图任务详情复制字段值接口响应参数
type CopyFieldRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TaskImageCopyFieldRes
}

// 绘图任务重新生成接口请求参数
type RegenerateReq struct {
	g.Meta `path:"/regenerate" method:"post" auth:"true" role:"admin" tags:"task_image" summary:"绘图任务重新生成接口"`
	Id     string `json:"id"`
}

// 绘图任务重新生成接口响应参数
type RegenerateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 绘图任务批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"admin" tags:"task_image" summary:"绘图任务批量操作接口"`
	model.TaskImageBatchOperateReq
}

// 绘图任务批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
