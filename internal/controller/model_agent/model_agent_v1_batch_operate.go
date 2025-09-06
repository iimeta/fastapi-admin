package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/model_agent/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error) {

	err = service.ModelAgent().BatchOperate(ctx, req.ModelAgentBatchOperateReq)

	return
}
