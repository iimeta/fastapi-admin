package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ManageRewardApplyApprove(ctx context.Context, req *v1.ManageRewardApplyApproveReq) (res *v1.ManageRewardApplyApproveRes, err error) {

	err = service.Invite().ManageRewardApplyApprove(ctx, req.InviteRewardApplyAuditReq)

	return
}
