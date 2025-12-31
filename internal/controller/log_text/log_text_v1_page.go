package log_text

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/log_text/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	textPageRes, err := service.LogText().Page(ctx, req.LogTextPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		LogTextPageRes: textPageRes,
	}

	return
}
