package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建用户接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"admin" tags:"admin_user" summary:"新建用户接口"`
	model.UserCreateReq
}

// 新建用户接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新用户接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"admin" tags:"admin_user" summary:"更新用户接口"`
	model.UserUpdateReq
}

// 更新用户接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改用户额度过期时间接口请求参数
type ChangeQuotaExpireReq struct {
	g.Meta `path:"/change/quota/expire" method:"post" auth:"true" role:"admin" tags:"admin_user" summary:"更改用户额度过期时间接口"`
	model.UserChangeQuotaExpireReq
}

// 更改用户额度过期时间接口响应参数
type ChangeQuotaExpireRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改用户状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"admin" tags:"admin_user" summary:"更改用户状态接口"`
	model.UserChangeStatusReq
}

// 更改用户状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除用户接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"admin" tags:"admin_user" summary:"删除用户接口"`
	Id     string `json:"id"`
}

// 删除用户接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 用户详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"admin" tags:"admin_user" summary:"用户详情接口"`
	Id     string `json:"id"`
}

// 用户详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UserDetailRes
}

// 用户分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"admin" tags:"admin_user" summary:"用户分页列表接口"`
	model.UserPageReq
}

// 用户分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UserPageRes
}

// 用户列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" method:"get" auth:"true" role:"admin" tags:"admin_user" summary:"用户列表接口"`
	model.UserListReq
}

// 用户列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UserListRes
}

// 授予用户额度接口请求参数
type GrantQuotaReq struct {
	g.Meta `path:"/grant/quota" method:"post" auth:"true" role:"admin" tags:"admin_user" summary:"授予用户额度接口"`
	model.UserGrantQuotaReq
}

// 授予用户额度接口响应参数
type GrantQuotaRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 用户模型权限接口请求参数
type ModelsReq struct {
	g.Meta `path:"/models" method:"post" auth:"true" role:"admin" tags:"admin_user" summary:"用户模型权限接口"`
	model.UserModelsReq
}

// 用户模型权限接口响应参数
type ModelsRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
