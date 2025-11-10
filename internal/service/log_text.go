// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/internal/model"
)

type (
	ILogText interface {
		// 文本日志详情
		Detail(ctx context.Context, id string) (*model.LogText, error)
		// 文本日志分页列表
		Page(ctx context.Context, params model.LogTextPageReq) (*model.LogTextPageRes, error)
		// 文本日志导出
		Export(ctx context.Context, params model.LogTextExportReq) (string, error)
		// 文本日志批量操作
		BatchOperate(ctx context.Context, params model.LogTextBatchOperateReq) error
		// 文本日志详情复制字段值
		CopyField(ctx context.Context, params model.LogTextCopyFieldReq) (string, error)
	}
)

var (
	localLogText ILogText
)

func LogText() ILogText {
	if localLogText == nil {
		panic("implement not found for interface ILogText, forgot register?")
	}
	return localLogText
}

func RegisterLogText(i ILogText) {
	localLogText = i
}
