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
	IAudio interface {
		// 音频日志详情
		Detail(ctx context.Context, id string) (*model.Audio, error)
		// 音频日志分页列表
		Page(ctx context.Context, params model.AudioPageReq) (*model.AudioPageRes, error)
		// 音频日志详情复制字段值
		CopyField(ctx context.Context, params model.AudioCopyFieldReq) (string, error)
	}
)

var (
	localAudio IAudio
)

func Audio() IAudio {
	if localAudio == nil {
		panic("implement not found for interface IAudio, forgot register?")
	}
	return localAudio
}

func RegisterAudio(i IAudio) {
	localAudio = i
}
