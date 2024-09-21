package statistics

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/statistics/v1"
)

func (c *ControllerV1) DataAppKey(ctx context.Context, req *v1.DataAppKeyReq) (res *v1.DataAppKeyRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	return
}
