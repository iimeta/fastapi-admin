package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ManageRewardApplyDelete(ctx context.Context, req *v1.ManageRewardApplyDeleteReq) (res *v1.ManageRewardApplyDeleteRes, err error) {

	err = service.Invite().ManageRewardApplyDelete(ctx, req.Id)

	return
}
