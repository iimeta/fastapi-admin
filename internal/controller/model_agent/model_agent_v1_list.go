package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model_agent/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

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
