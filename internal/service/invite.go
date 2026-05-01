// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
)

type (
	IInvite interface {
		// 生成唯一随机邀请码
		GenerateInviteCode(ctx context.Context) (string, error)
		// 校验邀请码格式并查询邀请人用户ID
		ResolveInviteCode(ctx context.Context, code string) (int, error)
		// 查询当前用户邀请概览, 必要时为历史用户懒生成邀请码
		Profile(ctx context.Context) (*model.InviteProfileRes, error)
		// 查询当前用户作为邀请人的邀请关系列表
		RelationsPage(ctx context.Context, params model.InviteRelationPageReq) (*model.InviteRelationPageRes, error)
		// 查询当前用户可申请、审核中或已入账的邀请收益列表
		RewardsPage(ctx context.Context, params model.InviteRewardPageReq) (*model.InviteRewardPageRes, error)
		// 将当前用户选中的待申请邀请收益提交为入账申请
		RewardApply(ctx context.Context, params model.InviteRewardApplyReq) error
		// 查询当前用户的邀请收益入账申请记录
		RewardApplyPage(ctx context.Context, params model.InviteRewardApplyPageReq) (*model.InviteRewardApplyPageRes, error)
		// 管理端查询邀请关系列表, 代理商角色自动限制为自身rid
		ManageRelationsPage(ctx context.Context, params model.InviteRelationPageReq) (*model.InviteRelationPageRes, error)
		// 管理端查询邀请收益列表, 代理商角色自动限制为自身rid
		ManageRewardsPage(ctx context.Context, params model.InviteRewardPageReq) (*model.InviteRewardPageRes, error)
		// 根据被邀请人充值流水生成邀请充值返利
		CreateRechargeRebate(ctx context.Context, inviteeUserId int, sourceDealRecordId string, rechargeQuota int) error
		// 管理端查询邀请收益入账申请列表, 代理商角色自动限制为自身rid
		ManageRewardApplyPage(ctx context.Context, params model.InviteRewardApplyPageReq) (*model.InviteRewardApplyPageRes, error)
		// 审核通过邀请收益入账申请, 将额度加到用户quota并写财务流水
		ManageRewardApplyApprove(ctx context.Context, params model.InviteRewardApplyAuditReq) error
		// 驳回邀请收益入账申请, 并按参数决定收益退回待申请或标记驳回
		ManageRewardApplyReject(ctx context.Context, params model.InviteRewardApplyAuditReq) error
		// 作废尚未申请入账的邀请收益
		ManageRewardsCancel(ctx context.Context, params model.InviteRewardsCancelReq) error
		// 取消邀请关系, 同时撤销关联的待申请收益
		ManageRelationsCancel(ctx context.Context, params model.InviteRelationCancelReq) error
		// 校验邀请注册IP限制, 返回true表示已触发限制
		CheckInviteIpLimit(ctx context.Context, ip string, inviterUserId int, siteConfig *entity.SiteConfig) bool
		// 校验邀请人是否超出单日或累计收益次数上限, 返回true表示已达上限
		CheckRewardLimit(ctx context.Context, inviterUserId int, siteConfig *entity.SiteConfig) bool
		// 用户首次登录时激活邀请关系: REGISTERED(1)->VALID(2), 并创建注册奖励
		ActivateInviteRelation(ctx context.Context, inviteeUserId int)
	}
)

var (
	localInvite IInvite
)

func Invite() IInvite {
	if localInvite == nil {
		panic("implement not found for interface IInvite, forgot register?")
	}
	return localInvite
}

func RegisterInvite(i IInvite) {
	localInvite = i
}
