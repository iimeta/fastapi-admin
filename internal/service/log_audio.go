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
	ILogAudio interface {
		// 音频日志详情
		Detail(ctx context.Context, id string) (*model.LogAudio, error)
		// 音频日志分页列表
		Page(ctx context.Context, params model.LogAudioPageReq) (*model.LogAudioPageRes, error)
		// 音频日志详情复制字段值
		CopyField(ctx context.Context, params model.LogAudioCopyFieldReq) (string, error)
	}
)

var (
	localLogAudio ILogAudio
)

func LogAudio() ILogAudio {
	if localLogAudio == nil {
		panic("implement not found for interface ILogAudio, forgot register?")
	}
	return localLogAudio
}

func RegisterLogAudio(i ILogAudio) {
	localLogAudio = i
}
