package admin_reseller

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/admin_reseller/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {

	err = service.AdminReseller().Delete(ctx, req.ResellerDeleteReq)

	return
}
