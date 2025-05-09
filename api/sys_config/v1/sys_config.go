package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 更新配置接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"admin" tags:"sys_config" summary:"更新配置接口"`
	model.SysConfigUpdateReq
}

// 更新配置接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改配置状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"admin" tags:"sys_config" summary:"更改配置状态接口"`
	model.SysConfigChangeStatusReq
}

// 更改配置状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 配置详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"admin" tags:"sys_config" summary:"配置详情接口"`
}

// 配置详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SysConfigDetailRes
}

// 重置配置接口请求参数
type ResetReq struct {
	g.Meta `path:"/reset" method:"post" auth:"true" role:"admin" tags:"sys_config" summary:"重置配置接口"`
	model.SysConfigResetReq
}

// 重置配置接口响应参数
type ResetRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 刷新配置接口请求参数
type RefreshReq struct {
	g.Meta `path:"/refresh" method:"post" auth:"true" role:"admin" tags:"sys_config" summary:"刷新配置接口"`
	model.SysConfigRefreshReq
}

// 刷新配置接口响应参数
type RefreshRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
