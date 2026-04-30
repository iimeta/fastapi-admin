package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ManageRewardApplyReject(ctx context.Context, req *v1.ManageRewardApplyRejectReq) (res *v1.ManageRewardApplyRejectRes, err error) {

	err = service.Invite().ManageRewardApplyReject(ctx, req.InviteRewardApplyAuditReq)

	return
}
