package do

import "github.com/gogf/gf/v2/util/gmeta"

// 邀请关系
type InviteRelation struct {
	gmeta.Meta    `collection:"invite_relation" bson:"-"`
	Id            string `bson:"_id,omitempty"`             // 邀请关系ID
	InviteCode    string `bson:"invite_code,omitempty"`     // 注册时使用的邀请码
	InviterUserId int    `bson:"inviter_user_id,omitempty"` // 邀请人用户ID
	InviteeUserId int    `bson:"invitee_user_id,omitempty"` // 被邀请人用户ID
	Rid           int    `bson:"rid,omitempty"`             // 代理商ID, 0表示平台主管用户
	Domain        string `bson:"domain,omitempty"`          // 注册时所属域名
	Terminal      string `bson:"terminal,omitempty"`        // 注册终端
	Channel       string `bson:"channel,omitempty"`         // 注册渠道
	Account       string `bson:"account,omitempty"`         // 被邀请人注册账号
	Ip            string `bson:"ip,omitempty"`              // 被邀请人注册IP
	Status        int    `bson:"status,omitempty"`          // 邀请关系状态[1:已注册, 2:有效, 3:无效, 4:已取消]
	RewardQuota   int    `bson:"reward_quota,omitempty"`    // 该邀请关系产生的奖励额度, 按系统内部整数额度存储
	RewardId      string `bson:"reward_id,omitempty"`       // 关联的邀请收益ID
	Remark        string `bson:"remark,omitempty"`          // 备注
	CreatedAt     int64  `bson:"created_at,omitempty"`      // 创建时间
	UpdatedAt     int64  `bson:"updated_at,omitempty"`      // 更新时间
}

// 邀请收益
type InviteReward struct {
	gmeta.Meta         `collection:"invite_reward" bson:"-"`
	Id                 string  `bson:"_id,omitempty"`                   // 邀请收益ID
	RelationId         string  `bson:"relation_id,omitempty"`           // 关联的邀请关系ID
	InviterUserId      int     `bson:"inviter_user_id,omitempty"`       // 邀请人用户ID, 也是收益归属用户
	InviteeUserId      int     `bson:"invitee_user_id,omitempty"`       // 被邀请人用户ID
	Rid                int     `bson:"rid,omitempty"`                   // 代理商ID, 0表示平台主管用户
	Quota              int     `bson:"quota,omitempty"`                 // 邀请收益额度, 按系统内部整数额度存储
	Status             int     `bson:"status,omitempty"`                // 邀请收益状态[1:待申请, 2:审核中, 3:已入账, 4:已驳回, 5:已取消]
	TriggerType        string  `bson:"trigger_type,omitempty"`          // 收益触发类型[register:注册奖励, recharge:充值返利]
	SourceDealRecordId string  `bson:"source_deal_record_id,omitempty"` // 来源充值流水ID
	RechargeSequence   int     `bson:"recharge_sequence,omitempty"`     // 第几次充值
	RechargeQuota      int     `bson:"recharge_quota,omitempty"`        // 本次充值额度, 按系统内部整数额度存储
	RebateType         string  `bson:"rebate_type,omitempty"`           // 返利类型[percent:百分比, fixed:固定额度]
	RebateRate         float64 `bson:"rebate_rate,omitempty"`           // 返利比例
	RebateQuota        int     `bson:"rebate_quota,omitempty"`          // 固定返利额度, 按系统内部整数额度存储
	ApplyOrderId       string  `bson:"apply_order_id,omitempty"`        // 关联的入账申请单号
	DealRecordId       string  `bson:"deal_record_id,omitempty"`        // 审核通过后生成的财务流水ID
	CreditedAt         int64   `bson:"credited_at,omitempty"`           // 入账时间
	RejectedReason     string  `bson:"rejected_reason,omitempty"`       // 驳回原因
	CancelledReason    string  `bson:"cancelled_reason,omitempty"`      // 作废原因
	CreatedAt          int64   `bson:"created_at,omitempty"`            // 创建时间
	UpdatedAt          int64   `bson:"updated_at,omitempty"`            // 更新时间
}

// 邀请收益入账申请
type InviteRewardApply struct {
	gmeta.Meta   `collection:"invite_reward_apply" bson:"-"`
	Id           string   `bson:"_id,omitempty"`            // 入账申请ID
	OrderNo      string   `bson:"order_no,omitempty"`       // 入账申请单号
	UserId       int      `bson:"user_id,omitempty"`        // 申请入账的用户ID
	Rid          int      `bson:"rid,omitempty"`            // 代理商ID, 0表示平台主管用户
	RewardIds    []string `bson:"reward_ids,omitempty"`     // 本次申请包含的邀请收益ID列表
	TotalQuota   int      `bson:"total_quota,omitempty"`    // 本次申请入账总额度, 按系统内部整数额度存储
	Status       int      `bson:"status,omitempty"`         // 入账申请状态[1:待审核, 2:已通过, 3:已入账, 4:已驳回, 5:已取消]
	AuditRole    string   `bson:"audit_role,omitempty"`     // 审核人角色[reseller:代理商, admin:管理员]
	AuditUserId  int      `bson:"audit_user_id,omitempty"`  // 审核人用户ID
	AuditRemark  string   `bson:"audit_remark,omitempty"`   // 审核备注
	RejectReason string   `bson:"reject_reason,omitempty"`  // 驳回原因
	DealRecordId string   `bson:"deal_record_id,omitempty"` // 审核通过后生成的财务流水ID
	AppliedAt    int64    `bson:"applied_at,omitempty"`     // 申请时间
	AuditedAt    int64    `bson:"audited_at,omitempty"`     // 审核时间
	CreditedAt   int64    `bson:"credited_at,omitempty"`    // 入账时间
	CreatedAt    int64    `bson:"created_at,omitempty"`     // 创建时间
	UpdatedAt    int64    `bson:"updated_at,omitempty"`     // 更新时间
}
