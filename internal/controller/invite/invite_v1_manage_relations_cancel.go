package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ManageRelationsCancel(ctx context.Context, req *v1.ManageRelationsCancelReq) (res *v1.ManageRelationsCancelRes, err error) {

	err = service.Invite().ManageRelationsCancel(ctx, req.InviteRelationCancelReq)

	return
}
