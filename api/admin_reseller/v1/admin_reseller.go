package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建代理商接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"admin" tags:"admin_reseller" summary:"新建代理商接口"`
	model.ResellerCreateReq
}

// 新建代理商接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新代理商接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"admin" tags:"admin_reseller" summary:"更新代理商接口"`
	model.ResellerUpdateReq
}

// 更新代理商接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改代理商额度过期时间接口请求参数
type ChangeQuotaExpireReq struct {
	g.Meta `path:"/change/quota/expire" method:"post" auth:"true" role:"admin" tags:"admin_reseller" summary:"更改代理商额度过期时间接口"`
	model.ResellerChangeQuotaExpireReq
}

// 更改代理商额度过期时间接口响应参数
type ChangeQuotaExpireRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更改代理商状态接口请求参数
type ChangeStatusReq struct {
	g.Meta `path:"/change/status" method:"post" auth:"true" role:"admin" tags:"admin_reseller" summary:"更改代理商状态接口"`
	model.ResellerChangeStatusReq
}

// 更改代理商状态接口响应参数
type ChangeStatusRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除代理商接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"admin" tags:"admin_reseller" summary:"删除代理商接口"`
	model.ResellerDeleteReq
}

// 删除代理商接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 代理商详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"admin" tags:"admin_reseller" summary:"代理商详情接口"`
	Id     string `json:"id"`
}

// 代理商详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ResellerDetailRes
}

// 代理商分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"admin" tags:"admin_reseller" summary:"代理商分页列表接口"`
	model.ResellerPageReq
}

// 代理商分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ResellerPageRes
}

// 代理商列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" method:"get" auth:"true" role:"admin" tags:"admin_reseller" summary:"代理商列表接口"`
	model.ResellerListReq
}

// 代理商列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ResellerListRes
}

// 代理商批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"admin" tags:"admin_reseller" summary:"代理商批量操作接口"`
	model.ResellerBatchOperateReq
}

// 代理商批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 代理商充值接口请求参数
type RechargeReq struct {
	g.Meta `path:"/recharge" method:"post" auth:"true" role:"admin" tags:"admin_reseller" summary:"代理商充值接口"`
	model.ResellerRechargeReq
}

// 代理商充值接口响应参数
type RechargeRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
