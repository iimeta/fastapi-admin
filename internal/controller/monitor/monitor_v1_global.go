package monitor

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/monitor/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Global(ctx context.Context, req *v1.GlobalReq) (res *v1.GlobalRes, err error) {

	data, err := service.Monitor().Global(ctx, req.MonitorGlobalReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GlobalRes{
		MonitorGlobalRes: data,
	}

	return
}
