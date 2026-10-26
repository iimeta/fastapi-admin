package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ManageRewardApplyBatchDelete(ctx context.Context, req *v1.ManageRewardApplyBatchDeleteReq) (res *v1.ManageRewardApplyBatchDeleteRes, err error) {

	err = service.Invite().ManageRewardApplyBatchDelete(ctx, req.InviteRewardApplyBatchDeleteReq)

	return
}
