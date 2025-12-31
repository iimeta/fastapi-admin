package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

// 文件日志详情接口请求参数
type DetailReq struct {
	g.Meta `path:"/detail" method:"get" auth:"true" role:"user,reseller,admin" tags:"log_file" summary:"文件日志详情接口"`
	Id     string `json:"id"`
}

// 文件日志详情接口响应参数
type DetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.LogFileDetailRes
}

// 文件日志分页列表接口请求参数
type PageReq struct {
	g.Meta `path:"/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"log_file" summary:"文件日志分页列表接口"`
	model.LogFilePageReq
}

// 文件日志分页列表接口响应参数
type PageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.LogFilePageRes
}

// 文件日志详情复制字段值接口请求参数
type CopyFieldReq struct {
	g.Meta `path:"/copy/field" method:"post" auth:"true" role:"user,reseller,admin" tags:"log_file" summary:"文件日志详情复制字段值详情接口"`
	model.LogFileCopyFieldReq
}

// 文件日志详情复制字段值接口响应参数
type CopyFieldRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.LogFileCopyFieldRes
}
