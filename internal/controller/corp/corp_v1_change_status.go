package corp

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/corp/v1"
)

func (c *ControllerV1) ChangeStatus(ctx context.Context, req *v1.ChangeStatusReq) (res *v1.ChangeStatusRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	err = service.Corp().ChangeStatus(ctx, req.CorpChangeStatusReq)

	return
}
