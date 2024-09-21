package statistics

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/statistics/v1"
)

func (c *ControllerV1) DataUser(ctx context.Context, req *v1.DataUserReq) (res *v1.DataUserRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	return
}
