package text

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/text/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	value, err := service.Text().CopyField(ctx, req.TextCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		TextCopyFieldRes: &model.TextCopyFieldRes{
			Value: value,
		},
	}

	return
}
