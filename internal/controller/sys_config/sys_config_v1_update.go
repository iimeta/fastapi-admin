package sys_config

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/sys_config/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	_, err = service.SysConfig().Update(ctx, req.SysConfigUpdateReq)

	return
}
