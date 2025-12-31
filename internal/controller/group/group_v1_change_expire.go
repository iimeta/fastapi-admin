package group

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/group/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ChangeExpire(ctx context.Context, req *v1.ChangeExpireReq) (res *v1.ChangeExpireRes, err error) {

	err = service.Group().ChangeExpire(ctx, req.GroupChangeExpireReq)

	return
}
