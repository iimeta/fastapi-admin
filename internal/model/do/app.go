package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	APP_COLLECTION = "app"
)

type App struct {
	gmeta.Meta     `collection:"app" bson:"-"`
	AppId          int      `bson:"app_id,omitempty"`     // 应用ID
	Name           string   `bson:"name,omitempty"`       // 应用名称
	Models         []string `bson:"models"`               // 模型权限
	IsLimitQuota   bool     `bson:"is_limit_quota"`       // 是否限制额度
	Quota          int      `bson:"quota"`                // 剩余额度
	UsedQuota      int      `bson:"used_quota,omitempty"` // 已用额度
	QuotaExpiresAt int64    `bson:"quota_expires_at"`     // 额度过期时间
	IpWhitelist    []string `bson:"ip_whitelist"`         // IP白名单
	IpBlacklist    []string `bson:"ip_blacklist"`         // IP黑名单
	Remark         string   `bson:"remark"`               // 备注
	Status         int      `bson:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	UserId         int      `bson:"user_id,omitempty"`    // 用户ID
	Creator        string   `bson:"creator,omitempty"`    // 创建人
	Updater        string   `bson:"updater,omitempty"`    // 更新人
	CreatedAt      int64    `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt      int64    `bson:"updated_at,omitempty"` // 更新时间
}
