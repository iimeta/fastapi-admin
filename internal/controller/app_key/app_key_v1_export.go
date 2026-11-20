package app_key

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/iimeta/fastapi-admin/v2/api/app_key/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
)

func (c *ControllerV1) Export(ctx context.Context, req *v1.ExportReq) (res *v1.ExportRes, err error) {

	filePath, err := service.AppKey().Export(ctx, req.AppKeyExportReq)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.ServeFile(filePath)

	if rerr := gfile.RemoveFile(filePath); rerr != nil {
		logger.Error(ctx, rerr)
	}

	return
}
