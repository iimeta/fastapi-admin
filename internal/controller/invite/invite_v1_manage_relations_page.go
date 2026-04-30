package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ManageRelationsPage(ctx context.Context, req *v1.ManageRelationsPageReq) (res *v1.ManageRelationsPageRes, err error) {

	pageRes, err := service.Invite().ManageRelationsPage(ctx, req.InviteRelationPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ManageRelationsPageRes{
		InviteRelationPageRes: pageRes,
	}

	return
}
