package corp

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/corp/v1"
)

func (c *ControllerV1) ChangePublic(ctx context.Context, req *v1.ChangePublicReq) (res *v1.ChangePublicRes, err error) {

	err = service.Corp().ChangePublic(ctx, req.CorpChangePublicReq)

	return
}
