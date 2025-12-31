package log_general

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/log_general/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	generalPageRes, err := service.LogGeneral().Page(ctx, req.LogGeneralPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		LogGeneralPageRes: generalPageRes,
	}

	return
}
