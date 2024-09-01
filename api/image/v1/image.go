package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 绘图详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"log/image" method:"get" summary:"绘图详情接口"`
	Id     string `json:"id"`
}

// 绘图详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ImageDetailRes
}

// 绘图分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"log/image" method:"post" summary:"绘图分页列表接口"`
	model.ImagePageReq
}

// 绘图分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ImagePageRes
}
