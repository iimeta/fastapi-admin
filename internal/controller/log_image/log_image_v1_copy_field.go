package log_image

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_image/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.LogImage().CopyField(ctx, req.LogImageCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		LogImageCopyFieldRes: &model.LogImageCopyFieldRes{
			Value: value,
		},
	}

	return
}
