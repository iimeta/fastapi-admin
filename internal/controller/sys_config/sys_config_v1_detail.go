package sys_config

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/sys_config/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	sysConfig, err := service.SysConfig().Detail(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		SysConfigDetailRes: &model.SysConfigDetailRes{
			SysConfig: sysConfig,
		},
	}

	return
}
