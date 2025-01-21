package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 更新配置接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" role:"admin" tags:"sys_config" method:"post" summary:"更新配置接口"`
	model.SysConfigUpdateReq
}

// 更新配置接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改配置状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" role:"admin" tags:"sys_config" method:"post" summary:"更改配置状态接口"`
	model.SysConfigChangeStatusReq
}

// 更改配置状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 配置详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" role:"admin" tags:"sys_config" method:"get" summary:"配置详情接口"`
}

// 配置详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SysConfigDetailRes
}

// 重置配置接口请求参数
type ResetReq struct {
	g.Meta `path:"/reset" role:"admin" tags:"sys_config" method:"get" summary:"重置配置接口"`
}

// 重置配置接口响应参数
type ResetRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
