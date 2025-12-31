package log_audio

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/log_audio/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.LogAudio().CopyField(ctx, req.LogAudioCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		LogAudioCopyFieldRes: &model.LogAudioCopyFieldRes{
			Value: value,
		},
	}

	return
}
