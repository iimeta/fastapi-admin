package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) RewardsPage(ctx context.Context, req *v1.RewardsPageReq) (res *v1.RewardsPageRes, err error) {

	pageRes, err := service.Invite().RewardsPage(ctx, req.InviteRewardPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.RewardsPageRes{
		InviteRewardPageRes: pageRes,
	}

	return
}
