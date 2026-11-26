package do

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

type Reseller struct {
	gmeta.Meta              `collection:"reseller" bson:"-"`
	Id                      string   `bson:"_id,omitempty"`                       // ID
	UserId                  int      `bson:"user_id,omitempty"`                   // 用户ID
	Name                    string   `bson:"name,omitempty"`                      // 姓名
	Avatar                  string   `bson:"avatar,omitempty"`                    // 头像
	Email                   string   `bson:"email,omitempty"`                     // 邮箱
	Phone                   string   `bson:"phone,omitempty"`                     // 手机号
	Quota                   int      `bson:"quota"`                               // 剩余额度
	UsedQuota               int      `bson:"used_quota,omitempty"`                // 已用额度
	QuotaExpiresAt          int64    `bson:"quota_expires_at"`                    // 额度过期时间
	IsCycleResetQuota       bool     `bson:"is_cycle_reset_quota"`                // 是否循环重置额度
	ResetQuota              int      `bson:"reset_quota,omitempty"`               // 重置额度
	CyclePeriod             int      `bson:"cycle_period,omitempty"`              // 循环周期
	PeriodUnit              string   `bson:"period_unit,omitempty"`               // 周期单位[hour:小时, day:天]
	ResetMode               string   `bson:"reset_mode,omitempty"`                // 重置模式[natural:自然周期, relative:相对周期]
	ResetAt                 int64    `bson:"reset_at"`                            // 重置时间
	NextResetAt             int64    `bson:"next_reset_at"`                       // 下次重置时间
	Groups                  []string `bson:"groups"`                              // 分组权限
	QuotaWarning            bool     `bson:"quota_warning"`                       // 额度预警开关
	WarningThresholds       []int    `bson:"warning_thresholds,omitempty"`        // 预警档位, 单位: $
	ExpireWarningThresholds []int    `bson:"expire_warning_thresholds,omitempty"` // 过期预警档位, 单位: 天
	WarningNotices          []int    `bson:"warning_notices,omitempty"`           // 已发送的额度预警档位, 单位: $
	ExhaustionNotice        bool     `bson:"exhaustion_notice"`                   // 耗尽通知
	ExpireWarningNotices    []int    `bson:"expire_warning_notices,omitempty"`    // 已发送的过期预警档位, 单位: 天
	ExpireNotice            bool     `bson:"expire_notice"`                       // 额度过期通知
	NoticeQuota             int      `bson:"notice_quota,omitempty"`              // 已发送额度预警档位对应的剩余额度, 额度回升后档位失效
	NoticeQuotaExpiry       int64    `bson:"notice_quota_expiry,omitempty"`       // 已发送过期预警档位对应的过期时间, 与过期时间不一致时档位失效
	Remark                  string   `bson:"remark"`                              // 备注
	Status                  int      `bson:"status,omitempty"`                    // 状态[1:正常, 2:禁用, -1:删除]
	Creator                 string   `bson:"creator,omitempty"`                   // 创建人
	Updater                 string   `bson:"updater,omitempty"`                   // 更新人
	CreatedAt               int64    `bson:"created_at,omitempty"`                // 创建时间
	UpdatedAt               int64    `bson:"updated_at,omitempty"`                // 更新时间
}

type ResellerAccount struct {
	gmeta.Meta         `collection:"reseller_account" bson:"-"`
	Id                 string `bson:"_id,omitempty"`                  // ID
	Uid                string `bson:"uid,omitempty"`                  // 用户主键ID
	UserId             int    `bson:"user_id,omitempty"`              // 用户ID
	Account            string `bson:"account,omitempty"`              // 账号
	Password           string `bson:"password,omitempty"`             // 密码
	Salt               string `bson:"salt,omitempty"`                 // 盐
	LoginIP            string `bson:"login_ip,omitempty"`             // 登录IP
	LoginTime          int64  `bson:"login_time,omitempty"`           // 登录时间
	LoginDomain        string `bson:"login_domain,omitempty"`         // 登录域名
	NeedChangePassword bool   `bson:"need_change_password,omitempty"` // 是否需要修改密码
	PasswordUpdatedAt  int64  `bson:"password_updated_at,omitempty"`  // 密码最后修改时间
	Remark             string `bson:"remark,omitempty"`               // 备注
	Status             int    `bson:"status,omitempty"`               // 状态[1:正常, 2:禁用, -1:删除]
	Creator            string `bson:"creator,omitempty"`              // 创建人
	Updater            string `bson:"updater,omitempty"`              // 更新人
	CreatedAt          int64  `bson:"created_at,omitempty"`           // 创建时间
	UpdatedAt          int64  `bson:"updated_at,omitempty"`           // 更新时间
}
