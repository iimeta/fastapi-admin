package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 系统配置接口请求参数
type SysConfigReq struct {
	g.Meta `path:"/sys/config" method:"get" tags:"open" summary:"系统配置接口"`
}

// 系统配置接口响应参数
type SysConfigRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SysConfigRes
}

// 站点配置接口请求参数
type SiteConfigReq struct {
	g.Meta `path:"/site/config" method:"post" tags:"open" summary:"站点配置接口"`
	model.SiteConfigDetailReq
}

// 站点配置接口响应参数
type SiteConfigRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SiteConfigDetailRes
}
