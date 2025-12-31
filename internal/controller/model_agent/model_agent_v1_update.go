package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model_agent/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	err = service.ModelAgent().Update(ctx, req.ModelAgentUpdateReq)

	return
}
