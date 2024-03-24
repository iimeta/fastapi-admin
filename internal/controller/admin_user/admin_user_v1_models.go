package admin_user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/admin_user/v1"
)

func (c *ControllerV1) Models(ctx context.Context, req *v1.ModelsReq) (res *v1.ModelsRes, err error) {

	err = service.AdminUser().Models(ctx, req.UserModelsReq)

	return
}
