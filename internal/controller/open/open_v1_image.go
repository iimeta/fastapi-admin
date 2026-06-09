package open

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/service"

	"github.com/iimeta/fastapi-admin/v2/api/open/v1"
)

func (c *ControllerV1) Image(ctx context.Context, req *v1.ImageReq) (res *v1.ImageRes, err error) {

	filePath, err := service.Open().Image(ctx, req.FileName)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.ServeFile(filePath)

	return
}
