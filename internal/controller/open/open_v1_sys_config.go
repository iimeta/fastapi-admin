package open

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/open/v1"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) SysConfig(ctx context.Context, req *v1.SysConfigReq) (res *v1.SysConfigRes, err error) {

	sysConfig, err := service.SysConfig().Config(ctx)
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
