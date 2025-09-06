package group

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/group/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) ChangePublic(ctx context.Context, req *v1.ChangePublicReq) (res *v1.ChangePublicRes, err error) {

	err = service.Group().ChangePublic(ctx, req.GroupChangePublicReq)

	return
}
