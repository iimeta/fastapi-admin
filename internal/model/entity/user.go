package entity

import "github.com/gogf/gf/v2/util/gmeta"

type User struct {
	gmeta.Meta       `role:"user,admin" bson:"-"`
	Id               string   `bson:"_id,omitempty"`               // ID
	UserId           int      `bson:"user_id,omitempty"`           // 用户ID
	Name             string   `bson:"name,omitempty"`              // 姓名
	Avatar           string   `bson:"avatar,omitempty"`            // 头像
	Email            string   `bson:"email,omitempty"`             // 邮箱
	Phone            string   `bson:"phone,omitempty"`             // 手机号
	Quota            int      `bson:"quota,omitempty"`             // 剩余额度
	UsedQuota        int      `bson:"used_quota,omitempty"`        // 已用额度
	QuotaExpiresAt   int64    `bson:"quota_expires_at,omitempty"`  // 额度过期时间
	Models           []string `bson:"models,omitempty"`            // 模型权限
	QuotaWarning     bool     `bson:"quota_warning,omitempty"`     // 额度预警
	WarningThreshold int      `bson:"warning_threshold,omitempty"` // 预警阈值
	WarningNotice    bool     `bson:"warning_notice,omitempty"`    // 预警通知
	ExhaustionNotice bool     `bson:"exhaustion_notice,omitempty"` // 耗尽通知
	Remark           string   `bson:"remark,omitempty"`            // 备注
	Status           int      `bson:"status,omitempty"`            // 状态[1:正常, 2:禁用, -1:删除]
	Creator          string   `bson:"creator,omitempty"`           // 创建人
	Updater          string   `bson:"updater,omitempty"`           // 更新人
	CreatedAt        int64    `bson:"created_at,omitempty"`        // 创建时间
	UpdatedAt        int64    `bson:"updated_at,omitempty"`        // 更新时间
}

type Account struct {
	gmeta.Meta `role:"*" bson:"-"`
	Id         string `bson:"_id,omitempty"`        // ID
	Uid        string `bson:"uid,omitempty"`        // 用户主键ID
	UserId     int    `bson:"user_id,omitempty"`    // 用户ID
	Account    string `bson:"account,omitempty"`    // 账号
	Password   string `bson:"password,omitempty"`   // 密码
	Salt       string `bson:"salt,omitempty"`       // 盐
	LoginIP    string `bson:"login_ip,omitempty"`   // 登录IP
	LoginTime  int64  `bson:"login_time,omitempty"` // 登录时间
	Remark     string `bson:"remark,omitempty"`     // 备注
	Status     int    `bson:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Creator    string `bson:"creator,omitempty"`    // 创建人
	Updater    string `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}
