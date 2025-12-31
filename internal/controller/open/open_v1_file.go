package open

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/service"

	"github.com/iimeta/fastapi-admin/v2/api/open/v1"
)

func (c *ControllerV1) File(ctx context.Context, req *v1.FileReq) (res *v1.FileRes, err error) {

	filePath, err := service.TaskFile().File(ctx, req.FileName)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.ServeFile(filePath)

	return
}
