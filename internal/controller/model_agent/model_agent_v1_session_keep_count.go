package model_agent

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"

	"github.com/iimeta/fastapi-admin/v2/api/model_agent/v1"
)

func (c *ControllerV1) SessionKeepCount(ctx context.Context, req *v1.SessionKeepCountReq) (res *v1.SessionKeepCountRes, err error) {

	count, err := service.ModelAgent().SessionKeepCount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.SessionKeepCountRes{
		ModelAgentSessionKeepCacheStatsRes: &model.ModelAgentSessionKeepCacheStatsRes{
			Count: count,
		},
	}

	return
}
