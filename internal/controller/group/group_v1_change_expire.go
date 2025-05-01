package group

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/group/v1"
)

func (c *ControllerV1) ChangeExpire(ctx context.Context, req *v1.ChangeExpireReq) (res *v1.ChangeExpireRes, err error) {

	err = service.Group().ChangeExpire(ctx, req.GroupChangeExpireReq)

	return
}
