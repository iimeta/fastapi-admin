package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model_agent/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ChangeWeight(ctx context.Context, req *v1.ChangeWeightReq) (res *v1.ChangeWeightRes, err error) {

	err = service.ModelAgent().ChangeWeight(ctx, req.ModelAgentChangeWeightReq)

	return
}
