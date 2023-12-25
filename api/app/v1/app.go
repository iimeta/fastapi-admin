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
	g.Meta `path:"/delete" tags:"app" method:"post" summary:"删除应用接口"`
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
	g.Meta `path:"/page" tags:"app" method:"post" summary:"应用分页列表接口"`
	model.AppPageReq
}

// 应用分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppPageRes
}

// 应用列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" tags:"app" method:"get" summary:"应用列表接口"`
	model.AppListReq
}

// 应用列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppListRes
}

// 新建应用密钥接口请求参数
type CreateKeyReq struct {
	g.Meta `path:"/create/key" tags:"app" method:"post" summary:"新建应用密钥接口"`
	model.AppCreateKeyReq
}

// 新建应用密钥接口响应参数
type CreateKeyRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppCreateKeyRes
}

// 应用密钥配置接口请求参数
type KeyConfigReq struct {
	g.Meta `path:"/key/config" tags:"app" method:"post" summary:"应用密钥配置接口"`
	model.AppKeyConfigReq
}

// 应用密钥配置接口响应参数
type KeyConfigRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
