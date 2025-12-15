package open

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/api/open/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Video(ctx context.Context, req *v1.VideoReq) (res *v1.VideoRes, err error) {

	filePath, err := service.TaskVideo().Video(ctx, req.FileName)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.ServeFile(filePath)

	return
}
