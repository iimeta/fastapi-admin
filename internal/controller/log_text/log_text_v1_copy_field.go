package log_text

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/log_text/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.LogText().CopyField(ctx, req.LogTextCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		LogTextCopyFieldRes: &model.LogTextCopyFieldRes{
			Value: value,
		},
	}

	return
}
