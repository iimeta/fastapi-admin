package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 登录接口请求参数
type LoginReq struct {
	g.Meta `path:"/login" tags:"auth" method:"post" summary:"登录接口"`
	model.LoginReq
}

// 登录接口响应参数
type LoginRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.LoginRes
}

// 登出接口请求参数
type LogoutReq struct {
	g.Meta `path:"/logout" tags:"auth" method:"post" summary:"登出接口"`
}

// 登出接口响应参数
type LogoutRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 注册接口请求参数
type RegisterReq struct {
	g.Meta `path:"/register" tags:"auth" method:"post" summary:"注册接口"`
	model.RegisterReq
}

// 注册接口响应参数
type RegisterRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// Token 刷新接口请求参数
type RefreshReq struct {
	g.Meta `path:"/refresh" tags:"auth" method:"post" summary:"刷新Token接口"`
}

// Token 刷新接口响应参数
type RefreshRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.RefreshRes
}

// 找回密码接口请求参数
type ForgetReq struct {
	g.Meta `path:"/forget" tags:"auth" method:"post" summary:"找回密码接口"`
	model.ForgetReq
}

// 找回密码接口响应参数
type ForgetRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
