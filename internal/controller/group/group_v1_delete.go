package group

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/group/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {

	err = service.Group().Delete(ctx, req.Id)

	return
}
