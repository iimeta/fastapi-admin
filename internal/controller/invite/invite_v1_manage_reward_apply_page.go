package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ManageRewardApplyPage(ctx context.Context, req *v1.ManageRewardApplyPageReq) (res *v1.ManageRewardApplyPageRes, err error) {

	pageRes, err := service.Invite().ManageRewardApplyPage(ctx, req.InviteRewardApplyPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ManageRewardApplyPageRes{
		InviteRewardApplyPageRes: pageRes,
	}

	return
}
