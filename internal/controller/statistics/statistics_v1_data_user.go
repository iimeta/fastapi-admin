package statistics

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/statistics/v1"
)

func (c *ControllerV1) DataUser(ctx context.Context, req *v1.DataUserReq) (res *v1.DataUserRes, err error) {

	//_, err = service.Statistics().DataUser(ctx, req.StatisticsDataReq)
	//if err != nil {
	//	return nil, err
	//}

	service.Statistics().StatisticsChat(ctx, nil)

	return
}
