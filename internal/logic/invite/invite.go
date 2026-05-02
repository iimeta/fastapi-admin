package invite

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	mrand "math/rand"
	"strings"
	"unicode"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	inviteCodeMinLength = 5
	inviteCodeMaxLength = 8
	inviteCodeDigits    = 2
	inviteCodeCharset   = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	inviteCodeNumbers   = "23456789"
	inviteCodeMaxRetry  = 32
)

type sInvite struct{}

func init() {
	service.RegisterInvite(New())
}

func New() service.IInvite {
	return &sInvite{}
}

// 生成唯一随机邀请码
func (s *sInvite) GenerateInviteCode(ctx context.Context) (string, error) {
	for i := 0; i < inviteCodeMaxRetry; i++ {
		code := s.generateRandomInviteCode()
		total, err := dao.User.CountDocuments(ctx, bson.M{"invite_code": code})
		if err != nil {
			logger.Error(ctx, err)
			return "", err
		}
		if total == 0 {
			return code, nil
		}
	}
	return "", errors.New("生成邀请码失败, 请稍后重试")
}

// 校验邀请码格式并查询邀请人用户ID
func (s *sInvite) ResolveInviteCode(ctx context.Context, code string) (int, error) {
	code = strings.TrimSpace(code)
	if !s.isValidInviteCodeFormat(code) {
		return 0, errors.New("无效的邀请码")
	}
	user, err := dao.User.FindOne(ctx, bson.M{"invite_code": code})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 0, errors.New("无效的邀请码")
		}
		logger.Error(ctx, err)
		return 0, err
	}
	return user.UserId, nil
}

func (s *sInvite) generateRandomInviteCode() string {
	length := inviteCodeMinLength + s.randomInt(inviteCodeMaxLength-inviteCodeMinLength+1)
	code := make([]byte, length)
	for i := range code {
		code[i] = inviteCodeCharset[s.randomInt(len(inviteCodeCharset))]
	}
	used := make(map[int]bool, inviteCodeDigits)
	for i := 0; i < inviteCodeDigits; i++ {
		pos := s.randomInt(length)
		for used[pos] {
			pos = s.randomInt(length)
		}
		used[pos] = true
		code[pos] = inviteCodeNumbers[s.randomInt(len(inviteCodeNumbers))]
	}
	return string(code)
}

func (s *sInvite) randomInt(max int) int {
	if max <= 0 {
		return 0
	}
	index, err := crand.Int(crand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return mrand.Intn(max)
	}
	return int(index.Int64())
}

func (s *sInvite) isValidInviteCodeFormat(code string) bool {
	if len(code) < inviteCodeMinLength || len(code) > inviteCodeMaxLength {
		return false
	}
	digitCount := 0
	for _, ch := range code {
		if !strings.ContainsRune(inviteCodeCharset, ch) {
			return false
		}
		if unicode.IsDigit(ch) {
			digitCount++
		}
	}
	return digitCount >= inviteCodeDigits
}

// 查询当前用户邀请概览, 必要时为历史用户懒生成邀请码
func (s *sInvite) Profile(ctx context.Context) (*model.InviteProfileRes, error) {
	userId := service.Session().GetUserId(ctx)
	user, err := dao.User.FindOne(ctx, bson.M{"user_id": userId})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	inviteCode := user.InviteCode
	if !s.isValidInviteCodeFormat(inviteCode) {
		inviteCode, err = s.GenerateInviteCode(ctx)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
		if _, err = dao.User.FindOneAndUpdate(ctx, bson.M{"user_id": userId}, bson.M{"invite_code": inviteCode}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}
	res := &model.InviteProfileRes{InviteCode: inviteCode, InviteLink: "/register/invite/" + inviteCode, CurrentQuota: common.ConvQuotaUnitReverse(user.Quota)}
	if r := g.RequestFromCtx(ctx); r != nil {
		if siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, r.GetHost()); siteConfig != nil {
			res.RuleText = siteConfig.InviteConfig.RuleText
			res.RewardQuota = common.ConvQuotaUnitReverse(siteConfig.InviteConfig.RewardQuota)
			res.GrantQuota = common.ConvQuotaUnitReverse(siteConfig.InviteConfig.GrantQuota)
			res.MinApplyQuota = common.ConvQuotaUnitReverse(siteConfig.InviteConfig.MinApplyQuota)
			res.RechargeRebateEnabled = siteConfig.InviteConfig.RechargeRebateEnabled
			res.RechargeRebateFirstEnabled = siteConfig.InviteConfig.RechargeRebateFirstEnabled
			res.RechargeRebateFirstType = siteConfig.InviteConfig.RechargeRebateFirstType
			res.RechargeRebateFirstRate = siteConfig.InviteConfig.RechargeRebateFirstRate
			res.RechargeRebateFirstQuota = common.ConvQuotaUnitReverse(siteConfig.InviteConfig.RechargeRebateFirstQuota)
			res.RechargeRebateSecondEnabled = siteConfig.InviteConfig.RechargeRebateSecondEnabled
			res.RechargeRebateSecondType = siteConfig.InviteConfig.RechargeRebateSecondType
			res.RechargeRebateSecondRate = siteConfig.InviteConfig.RechargeRebateSecondRate
			res.RechargeRebateSecondQuota = common.ConvQuotaUnitReverse(siteConfig.InviteConfig.RechargeRebateSecondQuota)
		}
	}
	if total, err := dao.InviteRelation.CountDocuments(ctx, bson.M{"inviter_user_id": userId}); err == nil {
		res.TotalInvites = total
	}
	res.PendingQuota = common.ConvQuotaUnitReverse(s.sumRewardQuota(ctx, bson.M{"inviter_user_id": userId, "status": consts.INVITE_REWARD_STATUS_PENDING}))
	res.ApplyingQuota = common.ConvQuotaUnitReverse(s.sumRewardQuota(ctx, bson.M{"inviter_user_id": userId, "status": consts.INVITE_REWARD_STATUS_APPLYING}))
	res.CreditedQuota = common.ConvQuotaUnitReverse(s.sumRewardQuota(ctx, bson.M{"inviter_user_id": userId, "status": consts.INVITE_REWARD_STATUS_CREDITED}))
	return res, nil
}

// 查询当前用户作为邀请人的邀请关系列表
func (s *sInvite) RelationsPage(ctx context.Context, params model.InviteRelationPageReq) (*model.InviteRelationPageRes, error) {
	params.InviterUserId = service.Session().GetUserId(ctx)
	res, err := s.relationsPage(ctx, params)
	if err != nil {
		return nil, err
	}
	s.sanitizeUserRelationPage(res)
	return res, nil
}

// 查询当前用户可申请、审核中或已入账的邀请收益列表
func (s *sInvite) RewardsPage(ctx context.Context, params model.InviteRewardPageReq) (*model.InviteRewardPageRes, error) {
	params.InviterUserId = service.Session().GetUserId(ctx)
	res, err := s.rewardsPage(ctx, params)
	if err != nil {
		return nil, err
	}
	s.sanitizeUserRewardPage(res)
	return res, nil
}

// 将当前用户选中的待申请邀请收益提交为入账申请
func (s *sInvite) RewardApply(ctx context.Context, params model.InviteRewardApplyReq) error {
	userId := service.Session().GetUserId(ctx)
	if len(params.RewardIds) == 0 {
		return errors.New("请选择申请入账的邀请收益")
	}
	rewards, err := dao.InviteReward.Find(ctx, bson.M{"_id": bson.M{"$in": params.RewardIds}, "inviter_user_id": userId, "status": consts.INVITE_REWARD_STATUS_PENDING})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	if len(rewards) != len(params.RewardIds) {
		return errors.New("存在不可申请的邀请收益")
	}
	totalQuota := 0
	rid := 0
	for _, reward := range rewards {
		totalQuota += reward.Quota
		rid = reward.Rid
	}
	user, err := dao.User.FindOne(ctx, bson.M{"user_id": userId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	if siteConfig := s.getUserSiteConfig(ctx, user); siteConfig != nil && siteConfig.InviteConfig.MinApplyQuota > 0 && totalQuota < siteConfig.InviteConfig.MinApplyQuota {
		return errors.Newf("邀请收益满 %g 才可申请入账", common.ConvQuotaUnitReverse(siteConfig.InviteConfig.MinApplyQuota))
	}
	now := gtime.TimestampMilli()
	apply := &do.InviteRewardApply{Id: util.GenerateId(), OrderNo: fmt.Sprintf("IA%d%d", userId, now), UserId: userId, Rid: rid, RewardIds: params.RewardIds, TotalQuota: totalQuota, Status: consts.INVITE_REWARD_APPLY_STATUS_PENDING, AppliedAt: now, CreatedAt: now, UpdatedAt: now}
	if _, err = dao.InviteRewardApply.Insert(ctx, apply); err != nil {
		logger.Error(ctx, err)
		return err
	}
	if err = dao.InviteReward.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": params.RewardIds}, "status": consts.INVITE_REWARD_STATUS_PENDING}, bson.M{"status": consts.INVITE_REWARD_STATUS_APPLYING, "apply_order_id": apply.OrderNo}); err != nil {
		logger.Error(ctx, err)
		return err
	}
	return nil
}

// 查询当前用户的邀请收益入账申请记录
func (s *sInvite) RewardApplyPage(ctx context.Context, params model.InviteRewardApplyPageReq) (*model.InviteRewardApplyPageRes, error) {
	params.UserId = service.Session().GetUserId(ctx)
	res, err := s.applyPage(ctx, params)
	if err != nil {
		return nil, err
	}
	s.sanitizeUserApplyPage(res)
	return res, nil
}

// 管理端查询邀请关系列表, 代理商角色自动限制为自身rid
func (s *sInvite) ManageRelationsPage(ctx context.Context, params model.InviteRelationPageReq) (*model.InviteRelationPageRes, error) {
	if service.Session().IsResellerRole(ctx) {
		params.Rid = service.Session().GetRid(ctx)
	}
	res, err := s.relationsPage(ctx, params)
	if err != nil {
		return nil, err
	}
	if service.Session().IsResellerRole(ctx) {
		s.sanitizeResellerRelationPage(res)
	}
	return res, nil
}

// 管理端查询邀请收益列表, 代理商角色自动限制为自身rid
func (s *sInvite) ManageRewardsPage(ctx context.Context, params model.InviteRewardPageReq) (*model.InviteRewardPageRes, error) {
	if service.Session().IsResellerRole(ctx) {
		params.Rid = service.Session().GetRid(ctx)
	}
	return s.rewardsPage(ctx, params)
}

// 根据被邀请人充值流水生成邀请充值返利
func (s *sInvite) CreateRechargeRebate(ctx context.Context, inviteeUserId int, sourceDealRecordId string, rechargeQuota int) error {
	if inviteeUserId == 0 || sourceDealRecordId == "" || rechargeQuota <= 0 {
		return nil
	}
	invitee, err := dao.User.FindOne(ctx, bson.M{"user_id": inviteeUserId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	if invitee.InviterUserId == 0 {
		return nil
	}
	relation, err := dao.InviteRelation.FindOne(ctx, bson.M{"inviter_user_id": invitee.InviterUserId, "invitee_user_id": inviteeUserId, "status": consts.INVITE_RELATION_STATUS_VALID})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}
		logger.Error(ctx, err)
		return err
	}
	siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, relation.Domain)
	if siteConfig == nil && relation.Rid != 0 {
		siteConfigs := service.SiteConfig().GetSiteConfigsByRid(ctx, relation.Rid)
		if len(siteConfigs) > 0 {
			siteConfig = siteConfigs[0]
		}
	}
	if siteConfig == nil || !siteConfig.InviteEnabled || !siteConfig.InviteConfig.RechargeRebateEnabled {
		return nil
	}
	rechargeCount, err := dao.DealRecord.CountDocuments(ctx, bson.M{"user_id": inviteeUserId, "type": 1, "quota": bson.M{"$gt": 0}, "status": 1, "created_at": bson.M{"$lte": gtime.TimestampMilli()}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	sequence := int(rechargeCount)
	enabled := false
	rebateType := ""
	rate := 0.0
	rebateQuota := 0
	if sequence == 1 {
		enabled = siteConfig.InviteConfig.RechargeRebateFirstEnabled
		rebateType = siteConfig.InviteConfig.RechargeRebateFirstType
		rate = siteConfig.InviteConfig.RechargeRebateFirstRate
		rebateQuota = siteConfig.InviteConfig.RechargeRebateFirstQuota
	} else if sequence >= 2 {
		enabled = siteConfig.InviteConfig.RechargeRebateSecondEnabled
		rebateType = siteConfig.InviteConfig.RechargeRebateSecondType
		rate = siteConfig.InviteConfig.RechargeRebateSecondRate
		rebateQuota = siteConfig.InviteConfig.RechargeRebateSecondQuota
	}
	if !enabled {
		return nil
	}
	if rebateType == "" {
		rebateType = "percent"
	}
	exists, err := dao.InviteReward.FindOne(ctx, bson.M{"trigger_type": consts.SCENE_RECHARGE, "source_deal_record_id": sourceDealRecordId, "inviter_user_id": invitee.InviterUserId, "invitee_user_id": inviteeUserId})
	if err == nil && exists != nil {
		return nil
	}
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
		return err
	}
	rewardQuota := 0
	switch rebateType {
	case "fixed":
		rewardQuota = rebateQuota
	case "percent":
		if rate <= 0 {
			return nil
		}
		rewardQuota = int(math.Floor(float64(rechargeQuota) * rate / 100))
	default:
		return nil
	}
	if rewardQuota <= 0 {
		return nil
	}
	if s.CheckRewardLimit(ctx, invitee.InviterUserId, siteConfig) {
		return nil
	}
	now := gtime.TimestampMilli()
	_, err = dao.InviteReward.Insert(ctx, &do.InviteReward{Id: util.GenerateId(), RelationId: relation.Id, InviterUserId: invitee.InviterUserId, InviteeUserId: inviteeUserId, Rid: relation.Rid, Quota: rewardQuota, Status: consts.INVITE_REWARD_STATUS_PENDING, TriggerType: consts.SCENE_RECHARGE, SourceDealRecordId: sourceDealRecordId, RechargeSequence: sequence, RechargeQuota: rechargeQuota, RebateType: rebateType, RebateRate: rate, RebateQuota: rebateQuota, CreatedAt: now, UpdatedAt: now})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	return nil
}

// 管理端查询邀请收益入账申请列表, 代理商角色自动限制为自身rid
func (s *sInvite) ManageRewardApplyPage(ctx context.Context, params model.InviteRewardApplyPageReq) (*model.InviteRewardApplyPageRes, error) {
	if service.Session().IsResellerRole(ctx) {
		params.Rid = service.Session().GetRid(ctx)
	}
	return s.applyPage(ctx, params)
}

// 审核通过邀请收益入账申请, 将额度加到用户quota并写财务流水
func (s *sInvite) ManageRewardApplyApprove(ctx context.Context, params model.InviteRewardApplyAuditReq) error {
	filter := bson.M{"_id": params.Id, "status": consts.INVITE_REWARD_APPLY_STATUS_PENDING}
	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}
	now := gtime.TimestampMilli()
	apply, err := dao.InviteRewardApply.FindOne(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	rewards, err := dao.InviteReward.Find(ctx, bson.M{"_id": bson.M{"$in": apply.RewardIds}, "status": consts.INVITE_REWARD_STATUS_APPLYING})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	if len(rewards) != len(apply.RewardIds) {
		return errors.New("存在不可入账的邀请收益")
	}
	apply, err = dao.InviteRewardApply.FindOneAndUpdate(ctx, filter, bson.M{"status": consts.INVITE_REWARD_APPLY_STATUS_APPROVED, "audit_role": service.Session().GetRole(ctx), "audit_user_id": service.Session().GetUserId(ctx), "audit_remark": params.AuditRemark, "audited_at": now})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	oldUser, err := dao.User.FindOne(ctx, bson.M{"user_id": apply.UserId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	newUser, err := dao.User.FindOneAndUpdate(ctx, bson.M{"user_id": apply.UserId}, bson.M{"$inc": bson.M{"quota": apply.TotalQuota}, "warning_notice": false, "exhaustion_notice": false, "expire_warning_notice": false, "expire_notice": false})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, apply.UserId), consts.USER_QUOTA_FIELD, int64(apply.TotalQuota)); err != nil {
		logger.Error(ctx, err)
		return err
	}
	dealId, err := dao.DealRecord.Insert(ctx, &do.DealRecord{UserId: apply.UserId, Quota: apply.TotalQuota, Type: 5, Remark: "邀请奖励: " + apply.OrderNo, Status: 1, Rid: newUser.Rid})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	if err = dao.InviteReward.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": apply.RewardIds}, "status": consts.INVITE_REWARD_STATUS_APPLYING}, bson.M{"status": consts.INVITE_REWARD_STATUS_CREDITED, "deal_record_id": dealId, "credited_at": now}); err != nil {
		logger.Error(ctx, err)
		return err
	}
	if err = dao.InviteRewardApply.UpdateOne(ctx, bson.M{"_id": apply.Id, "status": consts.INVITE_REWARD_APPLY_STATUS_APPROVED}, bson.M{"status": consts.INVITE_REWARD_APPLY_STATUS_CREDITED, "deal_record_id": dealId, "credited_at": now}); err != nil {
		logger.Error(ctx, err)
		return err
	}
	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{Action: consts.ACTION_UPDATE, OldData: oldUser, NewData: newUser}); err != nil {
		logger.Error(ctx, err)
		return err
	}
	return nil
}

// 驳回邀请收益入账申请, 并按参数决定收益退回待申请或标记驳回
func (s *sInvite) ManageRewardApplyReject(ctx context.Context, params model.InviteRewardApplyAuditReq) error {
	filter := bson.M{"_id": params.Id, "status": consts.INVITE_REWARD_APPLY_STATUS_PENDING}
	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}
	now := gtime.TimestampMilli()
	apply, err := dao.InviteRewardApply.FindOneAndUpdate(ctx, filter, bson.M{"status": consts.INVITE_REWARD_APPLY_STATUS_REJECTED, "audit_role": service.Session().GetRole(ctx), "audit_user_id": service.Session().GetUserId(ctx), "reject_reason": params.RejectReason, "audit_remark": params.AuditRemark, "audited_at": now})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	rewardStatus := consts.INVITE_REWARD_STATUS_REJECTED
	if params.ReturnPending {
		rewardStatus = consts.INVITE_REWARD_STATUS_PENDING
	}
	if err = dao.InviteReward.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": apply.RewardIds}, "status": consts.INVITE_REWARD_STATUS_APPLYING}, bson.M{"status": rewardStatus, "rejected_reason": params.RejectReason, "apply_order_id": ""}); err != nil {
		logger.Error(ctx, err)
		return err
	}
	return nil
}

// 作废尚未申请入账的邀请收益
func (s *sInvite) ManageRewardsCancel(ctx context.Context, params model.InviteRewardsCancelReq) error {
	filter := bson.M{"_id": bson.M{"$in": params.Ids}, "status": consts.INVITE_REWARD_STATUS_PENDING}
	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}
	if err := dao.InviteReward.UpdateMany(ctx, filter, bson.M{"status": consts.INVITE_REWARD_STATUS_CANCELLED, "cancelled_reason": params.CancelledReason}); err != nil {
		logger.Error(ctx, err)
		return err
	}
	return nil
}

// 取消邀请关系, 同时撤销关联的待申请收益
func (s *sInvite) ManageRelationsCancel(ctx context.Context, params model.InviteRelationCancelReq) error {

	filter := bson.M{
		"_id":    bson.M{"$in": params.Ids},
		"status": bson.M{"$in": []int{consts.INVITE_RELATION_STATUS_REGISTERED, consts.INVITE_RELATION_STATUS_VALID}},
	}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	// 查出即将被取消的关系, 用于后续撤销关联收益
	relations, err := dao.InviteRelation.Find(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if len(relations) == 0 {
		return nil
	}

	// 更新邀请关系状态为已取消
	if err = dao.InviteRelation.UpdateMany(ctx, filter, bson.M{"status": consts.INVITE_RELATION_STATUS_CANCELLED}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 收集关系ID, 撤销这些关系下的待申请收益
	relationIds := make([]string, 0, len(relations))
	for _, r := range relations {
		relationIds = append(relationIds, r.Id)
	}

	rewardFilter := bson.M{
		"relation_id": bson.M{"$in": relationIds},
		"status":      consts.INVITE_REWARD_STATUS_PENDING,
	}

	if err = dao.InviteReward.UpdateMany(ctx, rewardFilter, bson.M{"status": consts.INVITE_REWARD_STATUS_CANCELLED, "cancelled_reason": "邀请关系已取消"}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 按条件分页查询邀请关系并转换为前端展示模型
func (s *sInvite) relationsPage(ctx context.Context, params model.InviteRelationPageReq) (*model.InviteRelationPageRes, error) {
	paging := &db.Paging{Page: params.Page, PageSize: params.PageSize}
	filter := bson.M{}
	if params.InviterUserId != 0 {
		filter["inviter_user_id"] = params.InviterUserId
	}
	if params.InviteeUserId != 0 {
		filter["invitee_user_id"] = params.InviteeUserId
	}
	if params.Rid != 0 {
		filter["rid"] = params.Rid
	}
	if params.Status != 0 {
		filter["status"] = params.Status
	}
	if len(params.CreatedAt) > 0 {
		filter["created_at"] = bson.M{"$gte": gtime.NewFromStrFormat(params.CreatedAt[0], "2006-01-02").TimestampMilli(), "$lte": gtime.NewFromStrLayout(params.CreatedAt[1], "2006-01-02").EndOfDay(true).TimestampMilli()}
	}
	results, err := dao.InviteRelation.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	items := make([]*model.InviteRelation, 0)
	for _, result := range results {
		items = append(items, &model.InviteRelation{Id: result.Id, InviteCode: result.InviteCode, InviterUserId: result.InviterUserId, InviteeUserId: result.InviteeUserId, Rid: result.Rid, Domain: result.Domain, Terminal: result.Terminal, Channel: result.Channel, Account: result.Account, Ip: result.Ip, Status: result.Status, RewardQuota: common.ConvQuotaUnitReverse(result.RewardQuota), RewardId: result.RewardId, Remark: result.Remark, CreatedAt: util.FormatDateTime(result.CreatedAt), UpdatedAt: util.FormatDateTime(result.UpdatedAt)})
	}
	return &model.InviteRelationPageRes{Items: items, Paging: &model.Paging{Page: paging.Page, PageSize: paging.PageSize, Total: paging.Total}}, nil
}

// 按条件分页查询邀请收益并转换为前端展示模型
func (s *sInvite) rewardsPage(ctx context.Context, params model.InviteRewardPageReq) (*model.InviteRewardPageRes, error) {
	paging := &db.Paging{Page: params.Page, PageSize: params.PageSize}
	filter := bson.M{}
	if params.InviterUserId != 0 {
		filter["inviter_user_id"] = params.InviterUserId
	}
	if params.InviteeUserId != 0 {
		filter["invitee_user_id"] = params.InviteeUserId
	}
	if params.Rid != 0 {
		filter["rid"] = params.Rid
	}
	if params.Status != 0 {
		filter["status"] = params.Status
	}
	if len(params.CreatedAt) > 0 {
		filter["created_at"] = bson.M{"$gte": gtime.NewFromStrFormat(params.CreatedAt[0], "2006-01-02").TimestampMilli(), "$lte": gtime.NewFromStrLayout(params.CreatedAt[1], "2006-01-02").EndOfDay(true).TimestampMilli()}
	}
	results, err := dao.InviteReward.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	items := make([]*model.InviteReward, 0)
	for _, result := range results {
		items = append(items, &model.InviteReward{Id: result.Id, RelationId: result.RelationId, InviterUserId: result.InviterUserId, InviteeUserId: result.InviteeUserId, Rid: result.Rid, Quota: common.ConvQuotaUnitReverse(result.Quota), Status: result.Status, TriggerType: result.TriggerType, SourceDealRecordId: result.SourceDealRecordId, RechargeSequence: result.RechargeSequence, RechargeQuota: common.ConvQuotaUnitReverse(result.RechargeQuota), RebateType: result.RebateType, RebateRate: result.RebateRate, RebateQuota: common.ConvQuotaUnitReverse(result.RebateQuota), ApplyOrderId: result.ApplyOrderId, DealRecordId: result.DealRecordId, CreditedAt: util.FormatDateTime(result.CreditedAt), RejectedReason: result.RejectedReason, CancelledReason: result.CancelledReason, CreatedAt: util.FormatDateTime(result.CreatedAt), UpdatedAt: util.FormatDateTime(result.UpdatedAt)})
	}
	return &model.InviteRewardPageRes{Items: items, Paging: &model.Paging{Page: paging.Page, PageSize: paging.PageSize, Total: paging.Total}}, nil
}

// 按条件分页查询邀请收益入账申请并转换为前端展示模型
func (s *sInvite) applyPage(ctx context.Context, params model.InviteRewardApplyPageReq) (*model.InviteRewardApplyPageRes, error) {
	paging := &db.Paging{Page: params.Page, PageSize: params.PageSize}
	filter := bson.M{}
	if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}
	if params.Rid != 0 {
		filter["rid"] = params.Rid
	}
	if params.Status != 0 {
		filter["status"] = params.Status
	}
	if len(params.AppliedAt) > 0 {
		filter["applied_at"] = bson.M{"$gte": gtime.NewFromStrFormat(params.AppliedAt[0], "2006-01-02").TimestampMilli(), "$lte": gtime.NewFromStrLayout(params.AppliedAt[1], "2006-01-02").EndOfDay(true).TimestampMilli()}
	}
	results, err := dao.InviteRewardApply.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-applied_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	items := make([]*model.InviteRewardApply, 0)
	for _, result := range results {
		items = append(items, &model.InviteRewardApply{Id: result.Id, OrderNo: result.OrderNo, UserId: result.UserId, Rid: result.Rid, RewardIds: result.RewardIds, TotalQuota: common.ConvQuotaUnitReverse(result.TotalQuota), Status: result.Status, AuditRole: result.AuditRole, AuditUserId: result.AuditUserId, AuditRemark: result.AuditRemark, RejectReason: result.RejectReason, DealRecordId: result.DealRecordId, AppliedAt: util.FormatDateTime(result.AppliedAt), AuditedAt: util.FormatDateTime(result.AuditedAt), CreditedAt: util.FormatDateTime(result.CreditedAt), CreatedAt: util.FormatDateTime(result.CreatedAt), UpdatedAt: util.FormatDateTime(result.UpdatedAt)})
	}
	return &model.InviteRewardApplyPageRes{Items: items, Paging: &model.Paging{Page: paging.Page, PageSize: paging.PageSize, Total: paging.Total}}, nil
}

// 汇总符合条件的邀请收益内部额度
func (s *sInvite) sanitizeUserRelationPage(res *model.InviteRelationPageRes) {
	if res == nil {
		return
	}
	for _, item := range res.Items {
		item.Id = ""
		item.InviteCode = ""
		item.InviterUserId = 0
		item.Rid = 0
		item.Domain = ""
		item.Terminal = ""
		item.Channel = ""
		item.Account = ""
		item.Ip = ""
		item.RewardId = ""
		item.Remark = ""
		item.UpdatedAt = ""
	}
}

func (s *sInvite) sanitizeUserRewardPage(res *model.InviteRewardPageRes) {
	if res == nil {
		return
	}
	for _, item := range res.Items {
		item.RelationId = ""
		item.InviterUserId = 0
		item.Rid = 0
		item.SourceDealRecordId = ""
		item.DealRecordId = ""
		item.UpdatedAt = ""
	}
}

// 代理商管理端: 隐藏注册IP, 仅管理员可见
func (s *sInvite) sanitizeResellerRelationPage(res *model.InviteRelationPageRes) {
	if res == nil {
		return
	}
	for _, item := range res.Items {
		item.Ip = ""
	}
}

func (s *sInvite) sanitizeUserApplyPage(res *model.InviteRewardApplyPageRes) {
	if res == nil {
		return
	}
	for _, item := range res.Items {
		item.UserId = 0
		item.Rid = 0
		item.RewardIds = nil
		item.AuditRole = ""
		item.AuditUserId = 0
		item.AuditRemark = ""
		item.DealRecordId = ""
		item.CreatedAt = ""
		item.UpdatedAt = ""
	}
}
func (s *sInvite) sumRewardQuota(ctx context.Context, filter bson.M) int {
	rewards, err := dao.InviteReward.Find(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}
	total := 0
	for _, reward := range rewards {
		total += reward.Quota
	}
	return total
}

// 校验邀请注册IP限制, 返回true表示已触发限制
func (s *sInvite) CheckInviteIpLimit(ctx context.Context, ip string, inviterUserId int, siteConfig *entity.SiteConfig) bool {
	if siteConfig == nil || ip == "" {
		return false
	}
	if siteConfig.InviteConfig.IpDailyLimit > 0 {
		todayStart := gtime.Now().StartOfDay().TimestampMilli()
		count, err := dao.InviteRelation.CountDocuments(ctx, bson.M{"ip": ip, "created_at": bson.M{"$gte": todayStart}})
		if err != nil {
			logger.Error(ctx, err)
			return true
		}
		if int(count) >= siteConfig.InviteConfig.IpDailyLimit {
			return true
		}
	}
	if siteConfig.InviteConfig.IpTotalLimit > 0 {
		count, err := dao.InviteRelation.CountDocuments(ctx, bson.M{"ip": ip})
		if err != nil {
			logger.Error(ctx, err)
			return true
		}
		if int(count) >= siteConfig.InviteConfig.IpTotalLimit {
			return true
		}
	}
	if siteConfig.InviteConfig.IpPerInviterLimit > 0 {
		count, err := dao.InviteRelation.CountDocuments(ctx, bson.M{"ip": ip, "inviter_user_id": inviterUserId})
		if err != nil {
			logger.Error(ctx, err)
			return true
		}
		if int(count) >= siteConfig.InviteConfig.IpPerInviterLimit {
			return true
		}
	}
	return false
}

// 校验邀请人是否超出单日或累计收益次数上限, 返回true表示已达上限
func (s *sInvite) CheckRewardLimit(ctx context.Context, inviterUserId int, siteConfig *entity.SiteConfig) bool {
	if siteConfig == nil {
		return false
	}
	if siteConfig.InviteConfig.DailyLimit > 0 {
		todayStart := gtime.Now().StartOfDay().TimestampMilli()
		dailyCount, err := dao.InviteReward.CountDocuments(ctx, bson.M{"inviter_user_id": inviterUserId, "created_at": bson.M{"$gte": todayStart}})
		if err != nil {
			logger.Error(ctx, err)
			return true
		}
		if int(dailyCount) >= siteConfig.InviteConfig.DailyLimit {
			return true
		}
	}
	if siteConfig.InviteConfig.TotalLimit > 0 {
		totalCount, err := dao.InviteReward.CountDocuments(ctx, bson.M{"inviter_user_id": inviterUserId})
		if err != nil {
			logger.Error(ctx, err)
			return true
		}
		if int(totalCount) >= siteConfig.InviteConfig.TotalLimit {
			return true
		}
	}
	return false
}

// 获取用户所属站点配置, 优先使用当前请求域名, 代理商用户回退到rid配置
func (s *sInvite) getUserSiteConfig(ctx context.Context, user *entity.User) *entity.SiteConfig {
	if r := g.RequestFromCtx(ctx); r != nil {
		if siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, r.GetHost()); siteConfig != nil {
			return siteConfig
		}
	}
	if user.Rid == 0 {
		return nil
	}
	siteConfigs := service.SiteConfig().GetSiteConfigsByRid(ctx, user.Rid)
	if len(siteConfigs) == 0 {
		return nil
	}
	return siteConfigs[0]
}

// 用户首次登录时激活邀请关系: REGISTERED(1)->VALID(2), 并创建注册奖励
func (s *sInvite) ActivateInviteRelation(ctx context.Context, inviteeUserId int) {
	relation, err := dao.InviteRelation.FindOneAndUpdate(ctx, bson.M{"invitee_user_id": inviteeUserId, "status": consts.INVITE_RELATION_STATUS_REGISTERED}, bson.M{"status": consts.INVITE_RELATION_STATUS_VALID})
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
		}
		return
	}
	if relation.RewardQuota <= 0 || relation.InviterUserId == 0 {
		return
	}
	inviter, err := dao.User.FindOne(ctx, bson.M{"user_id": relation.InviterUserId, "status": 1})
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	siteConfig := s.getUserSiteConfig(ctx, inviter)
	if siteConfig != nil && !siteConfig.InviteEnabled {
		return
	}
	if s.CheckRewardLimit(ctx, relation.InviterUserId, siteConfig) {
		return
	}
	now := gtime.TimestampMilli()
	reward := &do.InviteReward{Id: util.GenerateId(), RelationId: relation.Id, InviterUserId: relation.InviterUserId, InviteeUserId: inviteeUserId, Rid: relation.Rid, Quota: relation.RewardQuota, Status: consts.INVITE_REWARD_STATUS_PENDING, TriggerType: consts.SCENE_REGISTER, CreatedAt: now, UpdatedAt: now}
	rewardId, err := dao.InviteReward.Insert(ctx, reward)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	if err = dao.InviteRelation.UpdateById(ctx, relation.Id, bson.M{"reward_id": rewardId}); err != nil {
		logger.Error(ctx, err)
	}
}
