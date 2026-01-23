package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model_agent/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) QuickFillModel(ctx context.Context, req *v1.QuickFillModelReq) (res *v1.QuickFillModelRes, err error) {

	models, err := service.ModelAgent().QuickFillModel(ctx, req.ModelAgentQuickFillModelReq)
	if err != nil {
		return nil, err
	}

	res = &v1.QuickFillModelRes{
		ModelAgentQuickFillModelRes: &model.ModelAgentQuickFillModelRes{
			Models: models,
		},
	}

	return
}
