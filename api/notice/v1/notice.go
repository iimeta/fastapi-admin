package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 新建消息通知接口请求参数
type CreateReq struct {
	g.Meta `path:"/create" method:"post" auth:"true" role:"admin" tags:"notice" summary:"新建消息通知接口"`
	model.NoticeCreateReq
}

// 新建消息通知接口响应参数
type CreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 更新消息通知接口请求参数
type UpdateReq struct {
	g.Meta `path:"/update" method:"post" auth:"true" role:"admin" tags:"notice" summary:"更新消息通知接口"`
	model.NoticeUpdateReq
}

// 更新消息通知接口响应参数
type UpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除消息通知接口请求参数
type DeleteReq struct {
	g.Meta `path:"/delete" method:"post" auth:"true" role:"admin" tags:"notice" summary:"删除消息通知接口"`
	Id     string `json:"id"`
}

// 删除消息通知接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 消息通知详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"admin" tags:"notice" summary:"消息通知详情接口"`
	Id     string `json:"id"`
}

// 消息通知详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticeDetailRes
}

// 消息通知分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"admin" tags:"notice" summary:"消息通知分页列表接口"`
	model.NoticePageReq
}

// 消息通知分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticePageRes
}

// 消息通知列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" method:"get" auth:"true" role:"admin" tags:"notice" summary:"消息通知列表接口"`
	model.NoticeListReq
}

// 消息通知列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticeListRes
}

// 消息通知批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"admin" tags:"notice" summary:"消息通知批量操作接口"`
	model.NoticeBatchOperateReq
}

// 消息通知批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
