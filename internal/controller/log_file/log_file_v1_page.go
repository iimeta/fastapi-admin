package log_file

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_file/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	filePageRes, err := service.LogFile().Page(ctx, req.LogFilePageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		LogFilePageRes: filePageRes,
	}

	return
}
