package log_audio

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_audio/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	audio, err := service.LogAudio().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		LogAudioDetailRes: &model.LogAudioDetailRes{
			LogAudio: audio,
		},
	}

	return
}
