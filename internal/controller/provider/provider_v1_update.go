package provider

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/provider/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	err = service.Provider().Update(ctx, req.ProviderUpdateReq)

	return
}
