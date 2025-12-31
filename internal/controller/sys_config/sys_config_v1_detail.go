package sys_config

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/sys_config/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

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
