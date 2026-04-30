package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {

	profileRes, err := service.Invite().Profile(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ProfileRes{
		InviteProfileRes: profileRes,
	}

	return
}
