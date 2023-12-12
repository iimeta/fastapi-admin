package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建应用接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" tags:"app" method:"post" summary:"新建应用接口"`
	model.AppCreateReq
}

// 新建应用接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新应用接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" tags:"app" method:"post" summary:"更新应用接口"`
	model.AppUpdateReq
}

// 更新应用接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除应用接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" tags:"app" method:"get" summary:"删除应用接口"`
	Id     string `json:"id"`
}

// 删除应用接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 应用详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"app" method:"get" summary:"应用详情接口"`
	Id     string `json:"id"`
}

// 应用详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppDetailRes
}

// 应用分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"app" method:"get" summary:"应用分页列表接口"`
	model.AppPageReq
}

// 应用分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppPageRes
}
