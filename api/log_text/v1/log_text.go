package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 文本日志详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"user,reseller,admin" tags:"log_text" summary:"文本日志详情接口"`
	Id     string `json:"id"`
}

// 文本日志详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.LogTextDetailRes
}

// 文本日志分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"log_text" summary:"文本日志分页列表接口"`
	model.LogTextPageReq
}

// 文本日志分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.LogTextPageRes
}

// 文本日志导出接口请求参数
type ExportReq struct {
	g.Meta `path:"/export" method:"post" auth:"true" role:"user,reseller,admin" tags:"log_text" summary:"文本日志导出接口"`
	model.LogTextExportReq
}

// 文本日志导出接口响应参数
type ExportRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文本日志批量操作接口请求参数
type BatchOperateReq struct {
	g.Meta `path:"/batch/operate" method:"post" auth:"true" role:"admin" tags:"log_text" summary:"文本日志批量操作接口"`
	model.LogTextBatchOperateReq
}

// 文本日志批量操作接口响应参数
type BatchOperateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文本日志详情复制字段值接口请求参数
type CopyFieldReq struct {
	g.Meta `path:"/copy/field" method:"post" auth:"true" role:"user,reseller,admin" tags:"log_text" summary:"文本日志详情复制字段值详情接口"`
	model.LogTextCopyFieldReq
}

// 文本日志详情复制字段值接口响应参数
type CopyFieldRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.LogTextCopyFieldRes
}
