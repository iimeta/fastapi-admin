package log_file

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_file/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	file, err := service.LogFile().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		LogFileDetailRes: &model.LogFileDetailRes{
			LogFile: file,
		},
	}

	return
}
