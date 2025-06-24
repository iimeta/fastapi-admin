package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建通知公告接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"reseller,admin" tags:"notice" summary:"新建通知公告接口"`
	model.NoticeCreateReq
}

// 新建通知公告接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新通知公告接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"reseller,admin" tags:"notice" summary:"更新通知公告接口"`
	model.NoticeUpdateReq
}

// 更新通知公告接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除通知公告接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"reseller,admin" tags:"notice" summary:"删除通知公告接口"`
	Id     string `json:"id"`
}

// 删除通知公告接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 通知公告详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"reseller,admin" tags:"notice" summary:"通知公告详情接口"`
	Id     string `json:"id"`
}

// 通知公告详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticeDetailRes
}

// 通知公告分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"reseller,admin" tags:"notice" summary:"通知公告分页列表接口"`
	model.NoticePageReq
}

// 通知公告分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticePageRes
}

// 通知公告列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" method:"get" auth:"true" role:"reseller,user,admin" tags:"notice" summary:"通知公告列表接口"`
	model.NoticeListReq
}

// 通知公告列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticeListRes
}

// 通知公告批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"reseller,admin" tags:"notice" summary:"通知公告批量操作接口"`
	model.NoticeBatchOperateReq
}

// 通知公告批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
