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
	IText interface {
		// 文本日志详情
		Detail(ctx context.Context, id string) (*model.Text, error)
		// 文本日志分页列表
		Page(ctx context.Context, params model.TextPageReq) (*model.TextPageRes, error)
		// 文本导出
		Export(ctx context.Context, params model.TextExportReq) (string, error)
		// 文本批量操作
		BatchOperate(ctx context.Context, params model.TextBatchOperateReq) error
		// 文本日志详情复制字段值
		CopyField(ctx context.Context, params model.TextCopyFieldReq) (string, error)
	}
)

var (
	localText IText
)

func Text() IText {
	if localText == nil {
		panic("implement not found for interface IText, forgot register?")
	}
	return localText
}

func RegisterText(i IText) {
	localText = i
}
