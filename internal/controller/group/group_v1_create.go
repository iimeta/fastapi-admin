package group

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/group/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	for _, rule := range req.TimeRules {
		rule.Discount = util.Round(rule.Discount/100, 2)
	}

	_, err = service.Group().Create(ctx, req.GroupCreateReq)

	return
}
