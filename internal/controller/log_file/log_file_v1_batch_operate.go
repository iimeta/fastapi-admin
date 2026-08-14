package log_file

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/log_file/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error) {

	err = service.LogFile().BatchOperate(ctx, req.LogFileBatchOperateReq)

	return
}
