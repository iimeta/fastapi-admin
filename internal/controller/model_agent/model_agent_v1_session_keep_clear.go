package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model_agent/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) SessionKeepClear(ctx context.Context, req *v1.SessionKeepClearReq) (res *v1.SessionKeepClearRes, err error) {

	_, err = service.ModelAgent().SessionKeepClear(ctx, req.Id)

	return
}
