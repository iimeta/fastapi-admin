package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

// ProfileReq 我的邀请信息请求。
type ProfileReq struct {
	g.Meta `path:"/profile" method:"get" auth:"true" role:"user,reseller,admin" tags:"invite" summary:"我的邀请信息接口"`
}

// ProfileRes 我的邀请信息响应。
type ProfileRes struct {
	g.Meta                  `mime:"application/json" example:"json"`
	*model.InviteProfileRes // 邀请码、邀请链接、邀请统计和收益额度概览。
}

// RelationsPageReq 我的邀请记录分页请求。
type RelationsPageReq struct {
	g.Meta                      `path:"/relations/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"invite" summary:"我的邀请记录分页接口"`
	model.InviteRelationPageReq // 邀请关系筛选条件和分页参数。
}

// RelationsPageRes 我的邀请记录分页响应。
type RelationsPageRes struct {
	g.Meta                       `mime:"application/json" example:"json"`
	*model.InviteRelationPageRes // 邀请关系列表和分页信息。
}

// RewardsPageReq 我的邀请收益分页请求。
type RewardsPageReq struct {
	g.Meta                    `path:"/rewards/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"invite" summary:"我的邀请收益分页接口"`
	model.InviteRewardPageReq // 邀请收益筛选条件和分页参数。
}

// RewardsPageRes 我的邀请收益分页响应。
type RewardsPageRes struct {
	g.Meta                     `mime:"application/json" example:"json"`
	*model.InviteRewardPageRes // 邀请收益列表和分页信息。
}

// RewardApplyReq 邀请收益入账申请请求。
type RewardApplyReq struct {
	g.Meta                     `path:"/reward/apply" method:"post" auth:"true" role:"user,reseller,admin" tags:"invite" summary:"邀请收益入账申请接口"`
	model.InviteRewardApplyReq // 待申请入账的邀请收益ID列表。
}

// RewardApplyRes 邀请收益入账申请响应。
type RewardApplyRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// RewardApplyPageReq 我的邀请收益入账申请分页请求。
type RewardApplyPageReq struct {
	g.Meta                         `path:"/reward/apply/page" method:"post" auth:"true" role:"user,reseller,admin" tags:"invite" summary:"我的邀请收益入账申请分页接口"`
	model.InviteRewardApplyPageReq // 入账申请筛选条件和分页参数。
}

// RewardApplyPageRes 我的邀请收益入账申请分页响应。
type RewardApplyPageRes struct {
	g.Meta                          `mime:"application/json" example:"json"`
	*model.InviteRewardApplyPageRes // 入账申请列表和分页信息。
}

// ManageRelationsPageReq 管理端邀请关系列表请求。
type ManageRelationsPageReq struct {
	g.Meta                      `path:"/manage/relations/page" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请关系列表接口"`
	model.InviteRelationPageReq // 邀请关系筛选条件和分页参数。
}

// ManageRelationsPageRes 管理端邀请关系列表响应。
type ManageRelationsPageRes struct {
	g.Meta                       `mime:"application/json" example:"json"`
	*model.InviteRelationPageRes // 邀请关系列表和分页信息。
}

// ManageRewardsPageReq 管理端邀请收益列表请求。
type ManageRewardsPageReq struct {
	g.Meta                    `path:"/manage/rewards/page" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请收益列表接口"`
	model.InviteRewardPageReq // 邀请收益筛选条件和分页参数。
}

// ManageRewardsPageRes 管理端邀请收益列表响应。
type ManageRewardsPageRes struct {
	g.Meta                     `mime:"application/json" example:"json"`
	*model.InviteRewardPageRes // 邀请收益列表和分页信息。
}

// ManageRewardApplyPageReq 管理端邀请收益入账申请列表请求。
type ManageRewardApplyPageReq struct {
	g.Meta                         `path:"/manage/reward/apply/page" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请收益入账申请列表接口"`
	model.InviteRewardApplyPageReq // 入账申请筛选条件和分页参数。
}

// ManageRewardApplyPageRes 管理端邀请收益入账申请列表响应。
type ManageRewardApplyPageRes struct {
	g.Meta                          `mime:"application/json" example:"json"`
	*model.InviteRewardApplyPageRes // 入账申请列表和分页信息。
}

// ManageRewardApplyApproveReq 邀请收益入账申请审核通过请求。
type ManageRewardApplyApproveReq struct {
	g.Meta                          `path:"/manage/reward/apply/approve" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请收益入账申请审核通过接口"`
	model.InviteRewardApplyAuditReq // 入账申请ID和审核备注。
}

// ManageRewardApplyApproveRes 邀请收益入账申请审核通过响应。
type ManageRewardApplyApproveRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// ManageRewardApplyRejectReq 邀请收益入账申请审核驳回请求。
type ManageRewardApplyRejectReq struct {
	g.Meta                          `path:"/manage/reward/apply/reject" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请收益入账申请审核驳回接口"`
	model.InviteRewardApplyAuditReq // 入账申请ID、驳回原因和收益退回策略。
}

// ManageRewardApplyRejectRes 邀请收益入账申请审核驳回响应。
type ManageRewardApplyRejectRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// ManageRewardsCancelReq 作废邀请收益请求。
type ManageRewardsCancelReq struct {
	g.Meta                       `path:"/manage/rewards/cancel" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"作废邀请收益接口"`
	model.InviteRewardsCancelReq // 待作废的邀请收益ID列表和作废原因。
}

// ManageRewardsCancelRes 作废邀请收益响应。
type ManageRewardsCancelRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
