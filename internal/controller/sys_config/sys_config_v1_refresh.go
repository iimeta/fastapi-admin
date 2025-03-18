package sys_config

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/sys_config/v1"
)

func (c *ControllerV1) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {

	err = service.SysConfig().Refresh(ctx, req.SysConfigRefreshReq)

	return
}
