package log_text

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_text/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	text, err := service.LogText().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		LogTextDetailRes: &model.LogTextDetailRes{
			LogText: text,
		},
	}

	return
}
