package log_general

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_general/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.LogGeneral().CopyField(ctx, req.LogGeneralCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		LogGeneralCopyFieldRes: &model.LogGeneralCopyFieldRes{
			Value: value,
		},
	}

	return
}
