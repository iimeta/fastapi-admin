package model_agent

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/iimeta/fastapi-admin/api/model_agent/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
