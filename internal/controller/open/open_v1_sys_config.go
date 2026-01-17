package open

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/api/open/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) SysConfig(ctx context.Context, req *v1.SysConfigReq) (res *v1.SysConfigRes, err error) {

	if req.Domain == "" {
		req.Domain = g.RequestFromCtx(ctx).GetHost()
	}

	sysConfig, err := service.SysConfig().Config(ctx, req.SysConfigReq)
	if err != nil {
		return nil, err
	}

	res = &v1.SysConfigRes{
		SysConfigRes: &model.SysConfigRes{
			SysConfig: sysConfig,
		},
	}

	return
}
