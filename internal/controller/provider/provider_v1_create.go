package provider

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/provider/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	_, err = service.Provider().Create(ctx, req.ProviderCreateReq)

	return
}
