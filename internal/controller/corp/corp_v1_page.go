package corp

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/corp/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	if service.Session().IsUserRole(ctx) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

	corpPageRes, err := service.Corp().Page(ctx, req.CorpPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		CorpPageRes: corpPageRes,
	}

	return
}
