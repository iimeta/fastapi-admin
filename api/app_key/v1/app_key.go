package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建应用密钥接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"reseller,user,admin" tags:"app_key" summary:"新建应用密钥接口"`
	model.AppKeyCreateReq
}

// 新建应用密钥接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppKeyCreateRes
}

// 应用密钥配置接口请求参数
type ConfigReq struct {
	g.Meta `path:"/config" method:"post" auth:"true" role:"reseller,user,admin" tags:"app_key" summary:"应用密钥配置接口"`
	model.AppKeyConfigReq
}

// 应用密钥配置接口响应参数
type ConfigRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Key    string `json:"key,omitempty"`
}

// 更改应用密钥状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"reseller,user,admin" tags:"app_key" summary:"更改应用密钥状态接口"`
	model.AppKeyChangeStatusReq
}

// 更改应用密钥状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除应用密钥接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"reseller,user,admin" tags:"app_key" summary:"删除应用密钥接口"`
	Id     string `json:"id"`
}

// 删除应用密钥接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 应用密钥详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"reseller,user,admin" tags:"app_key" summary:"应用密钥详情接口"`
	Id     string `json:"id"`
}

// 应用密钥详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppKeyDetailRes
}

// 应用密钥分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"reseller,user,admin" tags:"app_key" summary:"应用密钥分页列表接口"`
	model.AppKeyPageReq
}

// 应用密钥分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppKeyPageRes
}

// 应用密钥批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"reseller,user,admin" tags:"app_key" summary:"应用密钥批量操作接口"`
	model.AppKeyBatchOperateReq
}

// 应用密钥批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Keys   string `json:"keys,omitempty"`
}

// 应用密钥导出接口请求参数
type ExportReq struct {
	g.Meta `path:"/export" method:"post" auth:"true" role:"reseller,user,admin" tags:"app_key" summary:"应用密钥导出接口"`
	model.AppKeyExportReq
}

// 应用密钥导出接口响应参数
type ExportRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
