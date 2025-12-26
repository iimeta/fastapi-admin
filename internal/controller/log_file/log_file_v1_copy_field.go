package log_file

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_file/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.LogFile().CopyField(ctx, req.LogFileCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		LogFileCopyFieldRes: &model.LogFileCopyFieldRes{
			Value: value,
		},
	}

	return
}
