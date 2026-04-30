// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
)

type IInviteV1 interface {
	Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error)
	RelationsPage(ctx context.Context, req *v1.RelationsPageReq) (res *v1.RelationsPageRes, err error)
	RewardsPage(ctx context.Context, req *v1.RewardsPageReq) (res *v1.RewardsPageRes, err error)
	RewardApply(ctx context.Context, req *v1.RewardApplyReq) (res *v1.RewardApplyRes, err error)
	RewardApplyPage(ctx context.Context, req *v1.RewardApplyPageReq) (res *v1.RewardApplyPageRes, err error)
	ManageRelationsPage(ctx context.Context, req *v1.ManageRelationsPageReq) (res *v1.ManageRelationsPageRes, err error)
	ManageRewardsPage(ctx context.Context, req *v1.ManageRewardsPageReq) (res *v1.ManageRewardsPageRes, err error)
	ManageRewardApplyPage(ctx context.Context, req *v1.ManageRewardApplyPageReq) (res *v1.ManageRewardApplyPageRes, err error)
	ManageRewardApplyApprove(ctx context.Context, req *v1.ManageRewardApplyApproveReq) (res *v1.ManageRewardApplyApproveRes, err error)
	ManageRewardApplyReject(ctx context.Context, req *v1.ManageRewardApplyRejectReq) (res *v1.ManageRewardApplyRejectRes, err error)
	ManageRewardsCancel(ctx context.Context, req *v1.ManageRewardsCancelReq) (res *v1.ManageRewardsCancelRes, err error)
}
