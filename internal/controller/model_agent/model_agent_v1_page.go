package model_agent

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/model_agent/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	modelAgentPageRes, err := service.ModelAgent().Page(ctx, req.ModelAgentPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		ModelAgentPageRes: modelAgentPageRes,
	}

	return
}
