package corp

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/corp/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	_, err = service.Corp().Create(ctx, req.CorpCreateReq)

	return
}
