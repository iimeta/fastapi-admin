package invite

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/invite/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	profileRes, err := service.Invite().Profile(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ProfileRes{InviteProfileRes: profileRes}, nil
}

func (c *ControllerV1) RelationsPage(ctx context.Context, req *v1.RelationsPageReq) (res *v1.RelationsPageRes, err error) {
	pageRes, err := service.Invite().RelationsPage(ctx, req.InviteRelationPageReq)
	if err != nil {
		return nil, err
	}
	return &v1.RelationsPageRes{InviteRelationPageRes: pageRes}, nil
}

func (c *ControllerV1) RewardsPage(ctx context.Context, req *v1.RewardsPageReq) (res *v1.RewardsPageRes, err error) {
	pageRes, err := service.Invite().RewardsPage(ctx, req.InviteRewardPageReq)
	if err != nil {
		return nil, err
	}
	return &v1.RewardsPageRes{InviteRewardPageRes: pageRes}, nil
}

func (c *ControllerV1) RewardApply(ctx context.Context, req *v1.RewardApplyReq) (res *v1.RewardApplyRes, err error) {
	if err = service.Invite().RewardApply(ctx, req.InviteRewardApplyReq); err != nil {
		return nil, err
	}
	return &v1.RewardApplyRes{}, nil
}

func (c *ControllerV1) RewardApplyPage(ctx context.Context, req *v1.RewardApplyPageReq) (res *v1.RewardApplyPageRes, err error) {
	pageRes, err := service.Invite().RewardApplyPage(ctx, req.InviteRewardApplyPageReq)
	if err != nil {
		return nil, err
	}
	return &v1.RewardApplyPageRes{InviteRewardApplyPageRes: pageRes}, nil
}

func (c *ControllerV1) ManageRelationsPage(ctx context.Context, req *v1.ManageRelationsPageReq) (res *v1.ManageRelationsPageRes, err error) {
	pageRes, err := service.Invite().ManageRelationsPage(ctx, req.InviteRelationPageReq)
	if err != nil {
		return nil, err
	}
	return &v1.ManageRelationsPageRes{InviteRelationPageRes: pageRes}, nil
}

func (c *ControllerV1) ManageRewardsPage(ctx context.Context, req *v1.ManageRewardsPageReq) (res *v1.ManageRewardsPageRes, err error) {
	pageRes, err := service.Invite().ManageRewardsPage(ctx, req.InviteRewardPageReq)
	if err != nil {
		return nil, err
	}
	return &v1.ManageRewardsPageRes{InviteRewardPageRes: pageRes}, nil
}

func (c *ControllerV1) ManageRewardApplyPage(ctx context.Context, req *v1.ManageRewardApplyPageReq) (res *v1.ManageRewardApplyPageRes, err error) {
	pageRes, err := service.Invite().ManageRewardApplyPage(ctx, req.InviteRewardApplyPageReq)
	if err != nil {
		return nil, err
	}
	return &v1.ManageRewardApplyPageRes{InviteRewardApplyPageRes: pageRes}, nil
}

func (c *ControllerV1) ManageRewardApplyApprove(ctx context.Context, req *v1.ManageRewardApplyApproveReq) (res *v1.ManageRewardApplyApproveRes, err error) {
	if err = service.Invite().ManageRewardApplyApprove(ctx, req.InviteRewardApplyAuditReq); err != nil {
		return nil, err
	}
	return &v1.ManageRewardApplyApproveRes{}, nil
}

func (c *ControllerV1) ManageRewardApplyReject(ctx context.Context, req *v1.ManageRewardApplyRejectReq) (res *v1.ManageRewardApplyRejectRes, err error) {
	if err = service.Invite().ManageRewardApplyReject(ctx, req.InviteRewardApplyAuditReq); err != nil {
		return nil, err
	}
	return &v1.ManageRewardApplyRejectRes{}, nil
}

func (c *ControllerV1) ManageRewardsCancel(ctx context.Context, req *v1.ManageRewardsCancelReq) (res *v1.ManageRewardsCancelRes, err error) {
	if err = service.Invite().ManageRewardsCancel(ctx, req.InviteRewardsCancelReq); err != nil {
		return nil, err
	}
	return &v1.ManageRewardsCancelRes{}, nil
}
