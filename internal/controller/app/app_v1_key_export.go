package app

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/api/app/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) KeyExport(ctx context.Context, req *v1.KeyExportReq) (res *v1.KeyExportRes, err error) {

	filePath, err := service.App().KeyExport(ctx, req.AppKeyExportReq)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.ServeFile(filePath)

	return
}
