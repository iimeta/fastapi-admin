package model_agent

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model_agent/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	items, err := service.ModelAgent().List(ctx, req.ModelAgentListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListRes{
		ModelAgentListRes: &model.ModelAgentListRes{
			Items: items,
		},
	}

	return
}
