package text

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/text/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	textPageRes, err := service.Text().Page(ctx, req.TextPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		TextPageRes: textPageRes,
	}

	return
}
