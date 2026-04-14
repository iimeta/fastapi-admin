// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package statistics

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/statistics/v1"
)

type IStatisticsV1 interface {
	DataUser(ctx context.Context, req *v1.DataUserReq) (res *v1.DataUserRes, err error)
	DataApp(ctx context.Context, req *v1.DataAppReq) (res *v1.DataAppRes, err error)
	DataAppKey(ctx context.Context, req *v1.DataAppKeyReq) (res *v1.DataAppKeyRes, err error)
	DataSummary(ctx context.Context, req *v1.DataSummaryReq) (res *v1.DataSummaryRes, err error)
	DataTrend(ctx context.Context, req *v1.DataTrendReq) (res *v1.DataTrendRes, err error)
	DataModelPercent(ctx context.Context, req *v1.DataModelPercentReq) (res *v1.DataModelPercentRes, err error)
	DataTop(ctx context.Context, req *v1.DataTopReq) (res *v1.DataTopRes, err error)
	DataDetail(ctx context.Context, req *v1.DataDetailReq) (res *v1.DataDetailRes, err error)
	DataOverview(ctx context.Context, req *v1.DataOverviewReq) (res *v1.DataOverviewRes, err error)
	DataModelTrend(ctx context.Context, req *v1.DataModelTrendReq) (res *v1.DataModelTrendRes, err error)
	DataLatencyTrend(ctx context.Context, req *v1.DataLatencyTrendReq) (res *v1.DataLatencyTrendRes, err error)
	DataTaskStatus(ctx context.Context, req *v1.DataTaskStatusReq) (res *v1.DataTaskStatusRes, err error)
	DataAgentStatus(ctx context.Context, req *v1.DataAgentStatusReq) (res *v1.DataAgentStatusRes, err error)
	DataKeyStatus(ctx context.Context, req *v1.DataKeyStatusReq) (res *v1.DataKeyStatusRes, err error)
}
