package model_agent

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model_agent/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	modelAgent, err := service.ModelAgent().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		ModelAgentDetailRes: &model.ModelAgentDetailRes{
			ModelAgent: modelAgent,
		},
	}

	return
}
