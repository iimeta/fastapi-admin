package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 用户信息接口请求参数
type InfoReq struct {
	g.Meta `path:"/info" tags:"user" method:"get" summary:"用户信息接口"`
}

// 用户信息接口响应参数
type InfoRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UserInfoRes
}

// 用户修改密码接口请求参数
type ChangePasswordReq struct {
	g.Meta `path:"/change/password" tags:"user" method:"post" summary:"用户修改密码接口"`
	model.UserChangePasswordReq
}

// 用户修改密码接口响应参数
type ChangePasswordRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 用户修改邮箱接口请求参数
type ChangeEmailReq struct {
	g.Meta `path:"/change/email" tags:"user" method:"post" summary:"用户修改邮箱接口"`
	model.UserChangeEmailReq
}

// 用户修改邮箱接口响应参数
type ChangeEmailRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 用户更新信息接口请求参数
type UpdateInfoReq struct {
	g.Meta `path:"/update/info" tags:"user" method:"post" summary:"用户更新信息接口"`
	model.UserUpdateInfoReq
}

// 用户更新信息接口响应参数
type UpdateInfoRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 用户更改头像接口请求参数
type ChangeAvatarReq struct {
	g.Meta `path:"/change/avatar" tags:"user" method:"post" summary:"用户更改头像接口"`
	File   *ghttp.UploadFile `json:"file" type:"file"`
}

// 用户更改头像接口响应参数
type ChangeAvatarRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
