package sys_config

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/sys_config/v1"
)

func (c *ControllerV1) Reset(ctx context.Context, req *v1.ResetReq) (res *v1.ResetRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	_, err = service.SysConfig().Reset(ctx, req.SysConfigResetReq)

	return
}
