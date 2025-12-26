package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 批处理任务详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"user,reseller,admin" tags:"task_audio" summary:"批处理任务详情接口"`
	Id     string `json:"id"`
}

// 批处理任务详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TaskBatchDetailRes
}

// 批处理任务分页列表接口请求参数
type BatchPageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"task_audio" summary:"批处理任务分页列表接口"`
	model.TaskBatchPageReq
}

// 批处理任务分页列表接口响应参数
type BatchPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TaskBatchPageRes
}

// 批处理任务详情复制字段值接口请求参数
type CopyFieldReq struct {
	g.Meta `path:"/copy/field" method:"post" auth:"true" role:"user,reseller,admin" tags:"task_audio" summary:"批处理任务详情复制字段值详情接口"`
	model.TaskBatchCopyFieldReq
}

// 批处理任务详情复制字段值接口响应参数
type CopyFieldRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TaskBatchCopyFieldRes
}
