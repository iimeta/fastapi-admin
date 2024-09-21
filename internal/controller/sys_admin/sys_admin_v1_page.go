package sys_admin

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/iimeta/fastapi-admin/api/sys_admin/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
