package log_image

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_image/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	image, err := service.LogImage().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		LogImageDetailRes: &model.LogImageDetailRes{
			LogImage: image,
		},
	}

	return
}
