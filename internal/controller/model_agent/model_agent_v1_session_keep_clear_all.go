package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model_agent/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) SessionKeepClearAll(ctx context.Context, req *v1.SessionKeepClearAllReq) (res *v1.SessionKeepClearAllRes, err error) {

	_, err = service.ModelAgent().SessionKeepClearAll(ctx)

	return
}
