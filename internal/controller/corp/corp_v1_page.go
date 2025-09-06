package corp

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/corp/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	corpPageRes, err := service.Corp().Page(ctx, req.CorpPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		CorpPageRes: corpPageRes,
	}

	return
}
