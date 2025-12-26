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
	ILogFile interface {
		// 文件日志详情
		Detail(ctx context.Context, id string) (*model.LogFile, error)
		// 文件日志分页列表
		Page(ctx context.Context, params model.LogFilePageReq) (*model.LogFilePageRes, error)
		// 文件日志详情复制字段值
		CopyField(ctx context.Context, params model.LogFileCopyFieldReq) (string, error)
	}
)

var (
	localLogFile ILogFile
)

func LogFile() ILogFile {
	if localLogFile == nil {
		panic("implement not found for interface ILogFile, forgot register?")
	}
	return localLogFile
}

func RegisterLogFile(i ILogFile) {
	localLogFile = i
}
