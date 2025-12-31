package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

// 新建站点配置接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"reseller,admin" tags:"site_config" summary:"新建站点配置接口"`
	model.SiteConfigCreateReq
}

// 新建站点配置接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新站点配置接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"reseller,admin" tags:"site_config" summary:"更新站点配置接口"`
	model.SiteConfigUpdateReq
}

// 更新站点配置接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改站点配置状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"reseller,admin" tags:"site_config" summary:"更改站点配置状态接口"`
	model.SiteConfigChangeStatusReq
}

// 更改站点配置状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除站点配置接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"reseller,admin" tags:"site_config" summary:"删除站点配置接口"`
	Id     string `json:"id"`
}

// 删除站点配置接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 站点配置详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"reseller,admin" tags:"site_config" summary:"站点配置详情接口"`
	model.SiteConfigDetailReq
}

// 站点配置详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SiteConfigDetailRes
}

// 站点配置分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"reseller,admin" tags:"site_config" summary:"站点配置分页列表接口"`
	model.SiteConfigPageReq
}

// 站点配置分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SiteConfigPageRes
}

// 站点配置批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"reseller,admin" tags:"site_config" summary:"站点配置批量操作接口"`
	model.SiteConfigBatchOperateReq
}

// 站点配置批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
