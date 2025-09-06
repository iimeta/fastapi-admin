package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/model_agent/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {

	err = service.ModelAgent().Delete(ctx, req.Id)

	return
}
