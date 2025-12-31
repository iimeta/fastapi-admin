package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

// 新建应用接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"user,reseller,admin" tags:"app" summary:"新建应用接口"`
	model.AppCreateReq
}

// 新建应用接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Key    string `json:"key,omitempty"` // 密钥
}

// 更新应用接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"user,reseller,admin" tags:"app" summary:"更新应用接口"`
	model.AppUpdateReq
}

// 更新应用接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改应用状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"user,reseller,admin" tags:"app" summary:"更改应用状态接口"`
	model.AppChangeStatusReq
}

// 更改应用状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除应用接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"user,reseller,admin" tags:"app" summary:"删除应用接口"`
	Id     string `json:"id"`
}

// 删除应用接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 应用详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"user,reseller,admin" tags:"app" summary:"应用详情接口"`
	Id     string `json:"id"`
}

// 应用详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppDetailRes
}

// 应用分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"app" summary:"应用分页列表接口"`
	model.AppPageReq
}

// 应用分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppPageRes
}

// 应用列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" method:"get" auth:"true" role:"user,reseller,admin" tags:"app" summary:"应用列表接口"`
	model.AppListReq
}

// 应用列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AppListRes
}

// 应用批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"user,reseller,admin" tags:"app" summary:"应用批量操作接口"`
	model.AppBatchOperateReq
}

// 应用批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
