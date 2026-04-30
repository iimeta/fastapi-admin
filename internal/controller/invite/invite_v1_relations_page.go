package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) RelationsPage(ctx context.Context, req *v1.RelationsPageReq) (res *v1.RelationsPageRes, err error) {

	pageRes, err := service.Invite().RelationsPage(ctx, req.InviteRelationPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.RelationsPageRes{
		InviteRelationPageRes: pageRes,
	}

	return
}
