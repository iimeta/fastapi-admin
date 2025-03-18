package statistics

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/iimeta/fastapi-admin/api/statistics/v1"
)

func (c *ControllerV1) DataApp(ctx context.Context, req *v1.DataAppReq) (res *v1.DataAppRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
