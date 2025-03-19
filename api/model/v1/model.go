package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建模型接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"admin" tags:"model" summary:"新建模型接口"`
	model.ModelCreateReq
}

// 新建模型接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新模型接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"admin" tags:"model" summary:"更新模型接口"`
	model.ModelUpdateReq
}

// 更新模型接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改模型状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"admin" tags:"model" summary:"更改模型状态接口"`
	model.ModelChangeStatusReq
}

// 更改模型状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除模型接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"admin" tags:"model" summary:"删除模型接口"`
	Id     string `json:"id"`
}

// 删除模型接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 模型详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"admin" tags:"model" summary:"模型详情接口"`
	Id     string `json:"id"`
}

// 模型详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelDetailRes
}

// 模型分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"user,admin" tags:"model" summary:"模型分页列表接口"`
	model.ModelPageReq
}

// 模型分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelPageRes
}

// 模型列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" method:"get" auth:"true" role:"user,admin" tags:"model" summary:"模型列表接口"`
	model.ModelListReq
}

// 模型列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelListRes
}

// 模型批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"admin" tags:"model" summary:"模型批量操作接口"`
	model.ModelBatchOperateReq
}

// 模型批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 模型权限列表接口请求参数
type PermissionsReq struct {
	g.Meta `path:"/permissions" method:"post" auth:"true" role:"user,admin" tags:"model" summary:"模型权限列表接口"`
	model.ModelPermissionsReq
}

// 模型权限列表接口响应参数
type PermissionsRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelPermissionsRes
}

// 模型初始化同步接口请求参数
type InitSyncReq struct {
	g.Meta `path:"/init" method:"post" auth:"true" role:"admin" tags:"model" summary:"模型初始化同步接口"`
	model.ModelInitSyncReq
}

// 模型初始化同步接口响应参数
type InitSyncRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 模型树接口请求参数
type TreeReq struct {
	g.Meta `path:"/tree" method:"get" auth:"true" role:"user,admin" tags:"model" summary:"模型树接口"`
	model.ModelTreeReq
}

// 模型树接口响应参数
type TreeRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ModelTreeRes
}
