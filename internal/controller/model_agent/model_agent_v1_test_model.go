package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/model_agent/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) TestModel(ctx context.Context, req *v1.TestModelReq) (res *v1.TestModelRes, err error) {

	result, err := service.ModelAgent().TestModel(ctx, req.ModelAgentTestModelReq)

	res = &v1.TestModelRes{
		ModelAgentTestModelRes: result,
	}

	return
}
