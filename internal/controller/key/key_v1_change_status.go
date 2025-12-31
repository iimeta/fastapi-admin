package key

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/key/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ChangeStatus(ctx context.Context, req *v1.ChangeStatusReq) (res *v1.ChangeStatusRes, err error) {

	err = service.Key().ChangeStatus(ctx, req.KeyChangeStatusReq)

	return
}
