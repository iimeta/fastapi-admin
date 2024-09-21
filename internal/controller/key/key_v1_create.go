package key

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/key/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.Key().Create(ctx, req.KeyCreateReq, false)

	return
}
