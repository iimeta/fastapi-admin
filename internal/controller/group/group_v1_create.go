package group

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/group/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	req.Discount /= 100

	_, err = service.Group().Create(ctx, req.GroupCreateReq)

	return
}
