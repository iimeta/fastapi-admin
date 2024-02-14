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

// 新建用户接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" tags:"user" method:"post" summary:"新建用户接口"`
	model.UserCreateReq
}

// 新建用户接口响应参数
type CreateRes struct {
	g.Meta `mime:"userlication/json" example:"json"`
}

// 更新用户接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" tags:"user" method:"post" summary:"更新用户接口"`
	model.UserUpdateReq
}

// 更新用户接口响应参数
type UpdateRes struct {
	g.Meta `mime:"userlication/json" example:"json"`
}

// 更改用户状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" tags:"user" method:"post" summary:"更改用户状态接口"`
	model.UserChangeStatusReq
}

// 更改用户状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除用户接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" tags:"user" method:"post" summary:"删除用户接口"`
	Id     string `json:"id"`
}

// 删除用户接口响应参数
type DeleteRes struct {
	g.Meta `mime:"userlication/json" example:"json"`
}

// 用户详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" tags:"user" method:"get" summary:"用户详情接口"`
	Id     string `json:"id"`
}

// 用户详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"userlication/json" example:"json"`
	*model.UserDetailRes
}

// 用户分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" tags:"user" method:"post" summary:"用户分页列表接口"`
	model.UserPageReq
}

// 用户分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"userlication/json" example:"json"`
	*model.UserPageRes
}

// 用户列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" tags:"user" method:"get" summary:"用户列表接口"`
	model.UserListReq
}

// 用户列表接口响应参数
type ListRes struct {
	g.Meta `mime:"userlication/json" example:"json"`
	*model.UserListRes
}

// 用户授予额度接口请求参数
type GrantQuotaReq struct {
	g.Meta `path:"/grant/quota" tags:"user" method:"post" summary:"用户授予额度列表接口"`
	model.UserGrantQuotaReq
}

// 用户授予额度接口响应参数
type GrantQuotaRes struct {
	g.Meta `mime:"userlication/json" example:"json"`
}
