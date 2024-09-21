package audio

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/audio/v1"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	value, err := service.Audio().CopyField(ctx, req.AudioCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		AudioCopyFieldRes: &model.AudioCopyFieldRes{
			Value: value,
		},
	}

	return
}
