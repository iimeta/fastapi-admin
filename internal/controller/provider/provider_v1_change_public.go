package provider

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/provider/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) ChangePublic(ctx context.Context, req *v1.ChangePublicReq) (res *v1.ChangePublicRes, err error) {

	err = service.Provider().ChangePublic(ctx, req.ProviderChangePublicReq)

	return
}
