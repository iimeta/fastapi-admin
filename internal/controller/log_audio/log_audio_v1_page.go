package log_audio

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_audio/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	audioPageRes, err := service.LogAudio().Page(ctx, req.LogAudioPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		LogAudioPageRes: audioPageRes,
	}

	return
}
