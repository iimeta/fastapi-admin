package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 视频任务分页列表接口请求参数
type VideoPageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"task_audio" summary:"视频任务分页列表接口"`
	model.TaskVideoPageReq
}

// 视频任务分页列表接口响应参数
type VideoPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TaskVideoPageRes
}
