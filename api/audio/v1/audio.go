package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 音频详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" role:"user,admin" tags:"log/audio" method:"get" summary:"音频详情接口"`
	Id     string `json:"id"`
}

// 音频详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AudioDetailRes
}

// 音频分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" role:"user,admin" tags:"log/audio" method:"post" summary:"音频分页列表接口"`
	model.AudioPageReq
}

// 音频分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AudioPageRes
}

// 音频日志详情复制字段值接口请求参数
type CopyFieldReq struct {
	g.Meta `path:"/copy/field" role:"user,admin" tags:"log/audio" method:"post" summary:"音频日志详情复制字段值详情接口"`
	model.AudioCopyFieldReq
}

// 音频日志详情复制字段值接口响应参数
type CopyFieldRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AudioCopyFieldRes
}
