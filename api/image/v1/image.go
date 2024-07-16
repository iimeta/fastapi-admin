package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 图像详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"chat" method:"get" summary:"图像详情接口"`
	Id     string `json:"id"`
}

// 图像详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ImageDetailRes
}

// 图像分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"chat" method:"post" summary:"图像分页列表接口"`
	model.ImagePageReq
}

// 图像分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ImagePageRes
}
