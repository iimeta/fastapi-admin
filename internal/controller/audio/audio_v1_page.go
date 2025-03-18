package audio

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/audio/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	audioPageRes, err := service.Audio().Page(ctx, req.AudioPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		AudioPageRes: audioPageRes,
	}

	return
}
