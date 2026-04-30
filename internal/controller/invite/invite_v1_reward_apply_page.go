package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) RewardApplyPage(ctx context.Context, req *v1.RewardApplyPageReq) (res *v1.RewardApplyPageRes, err error) {

	pageRes, err := service.Invite().RewardApplyPage(ctx, req.InviteRewardApplyPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.RewardApplyPageRes{
		InviteRewardApplyPageRes: pageRes,
	}

	return
}
