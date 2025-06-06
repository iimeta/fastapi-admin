package entity

import "github.com/gogf/gf/v2/util/gmeta"

type App struct {
	gmeta.Meta     `role:"*" bson:"-"`
	Id             string   `bson:"_id,omitempty"`              // ID
	UserId         int      `bson:"user_id,omitempty"`          // 用户ID
	AppId          int      `bson:"app_id,omitempty"`           // 应用ID
	Name           string   `bson:"name,omitempty"`             // 应用名称
	Models         []string `bson:"models,omitempty"`           // 模型权限
	IsLimitQuota   bool     `bson:"is_limit_quota,omitempty"`   // 是否限制额度
	Quota          int      `bson:"quota,omitempty"`            // 剩余额度
	UsedQuota      int      `bson:"used_quota,omitempty"`       // 已用额度
	QuotaExpiresAt int64    `bson:"quota_expires_at,omitempty"` // 额度过期时间
	IsBindGroup    bool     `bson:"is_bind_group,omitempty"`    // 是否绑定分组
	Group          string   `bson:"group,omitempty"`            // 绑定分组
	IpWhitelist    []string `bson:"ip_whitelist,omitempty"`     // IP白名单
	IpBlacklist    []string `bson:"ip_blacklist,omitempty"`     // IP黑名单
	Remark         string   `bson:"remark,omitempty"`           // 备注
	Status         int      `bson:"status,omitempty"`           // 状态[1:正常, 2:禁用, -1:删除]
	Rid            int      `bson:"rid,omitempty"`              // 代理商ID
	Creator        string   `bson:"creator,omitempty"`          // 创建人
	Updater        string   `bson:"updater,omitempty"`          // 更新人
	CreatedAt      int64    `bson:"created_at,omitempty"`       // 创建时间
	UpdatedAt      int64    `bson:"updated_at,omitempty"`       // 更新时间
}
