package log_video

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_video/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.LogVideo().CopyField(ctx, req.LogVideoCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		LogVideoCopyFieldRes: &model.LogVideoCopyFieldRes{
			Value: value,
		},
	}

	return
}
