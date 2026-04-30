package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ManageRewardsCancel(ctx context.Context, req *v1.ManageRewardsCancelReq) (res *v1.ManageRewardsCancelRes, err error) {

	err = service.Invite().ManageRewardsCancel(ctx, req.InviteRewardsCancelReq)

	return
}
