package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ManageRewardsPage(ctx context.Context, req *v1.ManageRewardsPageReq) (res *v1.ManageRewardsPageRes, err error) {

	pageRes, err := service.Invite().ManageRewardsPage(ctx, req.InviteRewardPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ManageRewardsPageRes{
		InviteRewardPageRes: pageRes,
	}

	return
}
