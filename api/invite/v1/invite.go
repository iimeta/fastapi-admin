package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

// 我的邀请信息接口请求参数
type ProfileReq struct {
	g.Meta `path:"/profile" method:"get" auth:"true" role:"user" tags:"invite" summary:"我的邀请信息接口"`
}

// 我的邀请信息接口响应参数
type ProfileRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.InviteProfileRes
}

// 我的邀请记录分页接口请求参数
type RelationsPageReq struct {
	g.Meta `path:"/relations/page" method:"post" auth:"true" role:"user" tags:"invite" summary:"我的邀请记录分页接口"`
	model.InviteRelationPageReq
}

// 我的邀请记录分页接口响应参数
type RelationsPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.InviteRelationPageRes
}

// 我的邀请收益分页接口请求参数
type RewardsPageReq struct {
	g.Meta `path:"/rewards/page" method:"post" auth:"true" role:"user" tags:"invite" summary:"我的邀请收益分页接口"`
	model.InviteRewardPageReq
}

// 我的邀请收益分页接口响应参数
type RewardsPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.InviteRewardPageRes
}

// 邀请收益入账申请接口请求参数
type RewardApplyReq struct {
	g.Meta `path:"/reward/apply" method:"post" auth:"true" role:"user" tags:"invite" summary:"邀请收益入账申请接口"`
	model.InviteRewardApplyReq
}

// 邀请收益入账申请接口响应参数
type RewardApplyRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 我的邀请收益入账申请分页接口请求参数
type RewardApplyPageReq struct {
	g.Meta `path:"/reward/apply/page" method:"post" auth:"true" role:"user" tags:"invite" summary:"我的邀请收益入账申请分页接口"`
	model.InviteRewardApplyPageReq
}

// 我的邀请收益入账申请分页接口响应参数
type RewardApplyPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.InviteRewardApplyPageRes
}

// 管理端邀请关系列表接口请求参数
type ManageRelationsPageReq struct {
	g.Meta `path:"/manage/relations/page" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请关系列表接口"`
	model.InviteRelationPageReq
}

// 管理端邀请关系列表接口响应参数
type ManageRelationsPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.InviteRelationPageRes
}

// 管理端邀请收益列表接口请求参数
type ManageRewardsPageReq struct {
	g.Meta `path:"/manage/rewards/page" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请收益列表接口"`
	model.InviteRewardPageReq
}

// 管理端邀请收益列表接口响应参数
type ManageRewardsPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.InviteRewardPageRes
}

// 管理端邀请收益入账申请列表接口请求参数
type ManageRewardApplyPageReq struct {
	g.Meta `path:"/manage/reward/apply/page" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请收益入账申请列表接口"`
	model.InviteRewardApplyPageReq
}

// 管理端邀请收益入账申请列表接口响应参数
type ManageRewardApplyPageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.InviteRewardApplyPageRes
}

// 邀请收益入账申请审核通过接口请求参数
type ManageRewardApplyApproveReq struct {
	g.Meta `path:"/manage/reward/apply/approve" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请收益入账申请审核通过接口"`
	model.InviteRewardApplyAuditReq
}

// 邀请收益入账申请审核通过接口响应参数
type ManageRewardApplyApproveRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 邀请收益入账申请审核驳回接口请求参数
type ManageRewardApplyRejectReq struct {
	g.Meta `path:"/manage/reward/apply/reject" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"邀请收益入账申请审核驳回接口"`
	model.InviteRewardApplyAuditReq
}

// 邀请收益入账申请审核驳回接口响应参数
type ManageRewardApplyRejectRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 作废邀请收益接口请求参数
type ManageRewardsCancelReq struct {
	g.Meta `path:"/manage/rewards/cancel" method:"post" auth:"true" role:"reseller,admin" tags:"invite" summary:"作废邀请收益接口"`
	model.InviteRewardsCancelReq
}

// 作废邀请收益接口响应参数
type ManageRewardsCancelRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
