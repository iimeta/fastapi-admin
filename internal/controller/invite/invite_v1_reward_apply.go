package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) RewardApply(ctx context.Context, req *v1.RewardApplyReq) (res *v1.RewardApplyRes, err error) {

	err = service.Invite().RewardApply(ctx, req.InviteRewardApplyReq)

	return
}
