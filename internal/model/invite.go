package model

// 我的邀请信息接口响应参数
type InviteProfileRes struct {
	InviteCode                  string  `json:"invite_code,omitempty"`                    // 当前用户的邀请码
	InviteLink                  string  `json:"invite_link,omitempty"`                    // 当前用户的邀请注册链接
	RuleText                    string  `json:"rule_text,omitempty"`                      // 站点配置的邀请规则说明
	RewardQuota                 float64 `json:"reward_quota,omitempty"`                   // 邀请人每次邀请获得的收益额度
	GrantQuota                  float64 `json:"grant_quota,omitempty"`                    // 被邀请人通过邀请注册额外获得的额度
	TotalInvites                int64   `json:"total_invites,omitempty"`                  // 当前用户累计邀请成功人数
	PendingQuota                float64 `json:"pending_quota,omitempty"`                  // 待申请入账的邀请收益额度
	ApplyingQuota               float64 `json:"applying_quota,omitempty"`                 // 已提交申请、审核中的邀请收益额度
	CreditedQuota               float64 `json:"credited_quota,omitempty"`                 // 已审核通过并入账的邀请收益额度
	CurrentQuota                float64 `json:"current_quota,omitempty"`                  // 当前用户现有可用额度
	MinApplyQuota               float64 `json:"min_apply_quota,omitempty"`                // 邀请收益最低申请入账额度
	RechargeRebateEnabled       bool    `json:"recharge_rebate_enabled,omitempty"`        // 是否开启邀请充值返利
	RechargeRebateFirstEnabled  bool    `json:"recharge_rebate_first_enabled,omitempty"`  // 首次充值返利是否开启
	RechargeRebateFirstType     string  `json:"recharge_rebate_first_type,omitempty"`     // 首次充值返利类型
	RechargeRebateFirstRate     float64 `json:"recharge_rebate_first_rate,omitempty"`     // 首次充值返利比例
	RechargeRebateFirstQuota    float64 `json:"recharge_rebate_first_quota,omitempty"`    // 首次充值固定返利额度
	RechargeRebateSecondEnabled bool    `json:"recharge_rebate_second_enabled,omitempty"` // 后续充值返利是否开启
	RechargeRebateSecondType    string  `json:"recharge_rebate_second_type,omitempty"`    // 后续充值返利类型
	RechargeRebateSecondRate    float64 `json:"recharge_rebate_second_rate,omitempty"`    // 后续充值返利比例
	RechargeRebateSecondQuota   float64 `json:"recharge_rebate_second_quota,omitempty"`   // 后续充值固定返利额度
}

// 我的邀请记录分页接口请求参数
type InviteRelationPageReq struct {
	Paging
	InviteeUserId int      `json:"invitee_user_id,omitempty"` // 被邀请人用户ID
	InviterUserId int      `json:"inviter_user_id,omitempty"` // 邀请人用户ID
	Rid           int      `json:"rid,omitempty"`             // 代理商ID, 0表示平台主管用户
	Status        int      `json:"status,omitempty"`          // 邀请关系状态[1:已注册, 2:有效, 3:无效, 4:已取消]
	CreatedAt     []string `json:"created_at,omitempty"`      // 创建时间范围
}

// 我的邀请记录分页接口响应参数
type InviteRelationPageRes struct {
	Items  []*InviteRelation `json:"items"`  // 邀请关系列表
	Paging *Paging           `json:"paging"` // 分页信息
}

// 我的邀请收益分页接口请求参数
type InviteRewardPageReq struct {
	Paging
	InviteeUserId int      `json:"invitee_user_id,omitempty"` // 被邀请人用户ID
	InviterUserId int      `json:"inviter_user_id,omitempty"` // 邀请人用户ID
	Rid           int      `json:"rid,omitempty"`             // 代理商ID, 0表示平台主管用户
	Status        int      `json:"status,omitempty"`          // 邀请收益状态[1:待申请, 2:审核中, 3:已入账, 4:已驳回, 5:已取消]
	CreatedAt     []string `json:"created_at,omitempty"`      // 创建时间范围
}

// 我的邀请收益分页接口响应参数
type InviteRewardPageRes struct {
	Items  []*InviteReward `json:"items"`  // 邀请收益列表
	Paging *Paging         `json:"paging"` // 分页信息
}

// 邀请收益入账申请接口请求参数
type InviteRewardApplyReq struct {
	RewardIds []string `json:"reward_ids,omitempty" v:"required"` // 需要申请入账的邀请收益ID列表
}

// 我的邀请收益入账申请分页接口请求参数
type InviteRewardApplyPageReq struct {
	Paging
	UserId    int      `json:"user_id,omitempty"`    // 申请入账的用户ID
	Rid       int      `json:"rid,omitempty"`        // 代理商ID, 0表示平台主管用户
	Status    int      `json:"status,omitempty"`     // 入账申请状态[1:待审核, 2:已通过, 3:已入账, 4:已驳回, 5:已取消]
	AppliedAt []string `json:"applied_at,omitempty"` // 申请时间范围
}

// 我的邀请收益入账申请分页接口响应参数
type InviteRewardApplyPageRes struct {
	Items  []*InviteRewardApply `json:"items"`  // 入账申请列表
	Paging *Paging              `json:"paging"` // 分页信息
}

// 邀请收益入账申请审核通过接口请求参数
type InviteRewardApplyAuditReq struct {
	Id            string `json:"id,omitempty" v:"required"` // 入账申请ID
	AuditRemark   string `json:"audit_remark,omitempty"`    // 审核备注
	RejectReason  string `json:"reject_reason,omitempty"`   // 驳回原因
	ReturnPending bool   `json:"return_pending,omitempty"`  // 驳回时是否将收益退回待申请状态
}

// 作废邀请收益接口请求参数
type InviteRewardsCancelReq struct {
	Ids             []string `json:"ids,omitempty" v:"required"` // 需要作废的邀请收益ID列表
	CancelledReason string   `json:"cancelled_reason,omitempty"` // 作废原因
}

// 取消邀请关系接口请求参数
type InviteRelationCancelReq struct {
	Ids []string `json:"ids,omitempty" v:"required"` // 需要取消的邀请关系ID列表
}

// 邀请关系
type InviteRelation struct {
	Id            string  `json:"id,omitempty"`              // 邀请关系ID
	InviteCode    string  `json:"invite_code,omitempty"`     // 注册时使用的邀请码
	InviterUserId int     `json:"inviter_user_id,omitempty"` // 邀请人用户ID
	InviteeUserId int     `json:"invitee_user_id,omitempty"` // 被邀请人用户ID
	Rid           int     `json:"rid,omitempty"`             // 代理商ID, 0表示平台主管用户
	Domain        string  `json:"domain,omitempty"`          // 注册时所属域名
	Terminal      string  `json:"terminal,omitempty"`        // 注册终端
	Channel       string  `json:"channel,omitempty"`         // 注册渠道
	Account       string  `json:"account,omitempty"`         // 被邀请人注册账号
	Ip            string  `json:"ip,omitempty"`              // 被邀请人注册IP
	Status        int     `json:"status,omitempty"`          // 邀请关系状态[1:已注册, 2:有效, 3:无效, 4:已取消]
	RewardQuota   float64 `json:"reward_quota,omitempty"`    // 该邀请关系产生的奖励额度
	RewardId      string  `json:"reward_id,omitempty"`       // 关联的邀请收益ID
	Remark        string  `json:"remark,omitempty"`          // 备注
	CreatedAt     string  `json:"created_at,omitempty"`      // 创建时间
	UpdatedAt     string  `json:"updated_at,omitempty"`      // 更新时间
}

// 邀请收益
type InviteReward struct {
	Id                 string  `json:"id,omitempty"`                    // 邀请收益ID
	RelationId         string  `json:"relation_id,omitempty"`           // 关联的邀请关系ID
	InviterUserId      int     `json:"inviter_user_id,omitempty"`       // 邀请人用户ID, 也是收益归属用户
	InviteeUserId      int     `json:"invitee_user_id,omitempty"`       // 被邀请人用户ID
	Rid                int     `json:"rid,omitempty"`                   // 代理商ID, 0表示平台主管用户
	Quota              float64 `json:"quota,omitempty"`                 // 邀请收益额度
	Status             int     `json:"status,omitempty"`                // 邀请收益状态[1:待申请, 2:审核中, 3:已入账, 4:已驳回, 5:已取消]
	TriggerType        string  `json:"trigger_type,omitempty"`          // 收益触发类型[register:注册奖励, recharge:充值返利]
	SourceDealRecordId string  `json:"source_deal_record_id,omitempty"` // 来源充值流水ID
	RechargeSequence   int     `json:"recharge_sequence,omitempty"`     // 第几次充值
	RechargeQuota      float64 `json:"recharge_quota,omitempty"`        // 本次充值额度
	RebateType         string  `json:"rebate_type,omitempty"`           // 返利类型[percent:百分比, fixed:固定额度]
	RebateRate         float64 `json:"rebate_rate,omitempty"`           // 返利比例
	RebateQuota        float64 `json:"rebate_quota,omitempty"`          // 固定返利额度
	ApplyOrderId       string  `json:"apply_order_id,omitempty"`        // 关联的入账申请单号
	DealRecordId       string  `json:"deal_record_id,omitempty"`        // 审核通过后生成的财务流水ID
	CreditedAt         string  `json:"credited_at,omitempty"`           // 入账时间
	RejectedReason     string  `json:"rejected_reason,omitempty"`       // 驳回原因
	CancelledReason    string  `json:"cancelled_reason,omitempty"`      // 作废原因
	CreatedAt          string  `json:"created_at,omitempty"`            // 创建时间
	UpdatedAt          string  `json:"updated_at,omitempty"`            // 更新时间
}

// 邀请收益入账申请
type InviteRewardApply struct {
	Id           string   `json:"id,omitempty"`             // 入账申请ID
	OrderNo      string   `json:"order_no,omitempty"`       // 入账申请单号
	UserId       int      `json:"user_id,omitempty"`        // 申请入账的用户ID
	Rid          int      `json:"rid,omitempty"`            // 代理商ID, 0表示平台主管用户
	RewardIds    []string `json:"reward_ids,omitempty"`     // 本次申请包含的邀请收益ID列表
	TotalQuota   float64  `json:"total_quota,omitempty"`    // 本次申请入账的总额度
	Status       int      `json:"status,omitempty"`         // 入账申请状态[1:待审核, 2:已通过, 3:已入账, 4:已驳回, 5:已取消]
	AuditRole    string   `json:"audit_role,omitempty"`     // 审核人角色[reseller:代理商, admin:管理员]
	AuditUserId  int      `json:"audit_user_id,omitempty"`  // 审核人用户ID
	AuditRemark  string   `json:"audit_remark,omitempty"`   // 审核备注
	RejectReason string   `json:"reject_reason,omitempty"`  // 驳回原因
	DealRecordId string   `json:"deal_record_id,omitempty"` // 审核通过后生成的财务流水ID
	AppliedAt    string   `json:"applied_at,omitempty"`     // 申请时间
	AuditedAt    string   `json:"audited_at,omitempty"`     // 审核时间
	CreditedAt   string   `json:"credited_at,omitempty"`    // 入账时间
	CreatedAt    string   `json:"created_at,omitempty"`     // 创建时间
	UpdatedAt    string   `json:"updated_at,omitempty"`     // 更新时间
}
