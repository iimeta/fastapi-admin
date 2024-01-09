package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 登录用户信息接口请求参数
type UserInfoReq struct {
	g.Meta `path:"/info" tags:"user" method:"get" summary:"登录用户信息接口"`
}

// 登录用户信息接口响应参数
type UserInfoRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UserInfoRes
}

// 用户配置信息请求参数
type UserSettingReq struct {
	g.Meta `path:"/setting" tags:"user" method:"get" summary:"用户配置信息接口"`
}

// 用户配置信息响应参数
type UserSettingRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UserSettingRes
}

// 用户信息更新接口请求参数
type UserDetailUpdateReq struct {
	g.Meta `path:"/change/detail" tags:"user" method:"post" summary:"用户信息更新接口"`
	model.UserDetailUpdateReq
}

// 用户信息更新接口响应参数
type UserDetailUpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 用户密码更新接口请求参数
type UserPasswordUpdateReq struct {
	g.Meta `path:"/change/password" tags:"user" method:"post" summary:"用户密码更新接口"`
	model.UserPasswordUpdateReq
}

// 用户密码更新接口响应参数
type UserPasswordUpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 用户手机号更新接口请求参数
type UserPhoneUpdateReq struct {
	g.Meta `path:"/change/phone" tags:"user" method:"post" summary:"用户手机号更新接口"`
	model.UserPhoneUpdateReq
}

// 用户手机号更新接口响应参数
type UserPhoneUpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 用户邮箱更新接口请求参数
type UserEmailUpdateReq struct {
	g.Meta `path:"/change/email" tags:"user" method:"post" summary:"用户邮箱更新接口"`
	model.UserEmailUpdateReq
}

// 用户邮箱更新接口响应参数
type UserEmailUpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
