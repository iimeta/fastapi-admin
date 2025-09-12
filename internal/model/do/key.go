package do

import "github.com/gogf/gf/v2/util/gmeta"

type Key struct {
	gmeta.Meta         `collection:"key" bson:"-"`
	ProviderId         string   `bson:"provider_id,omitempty"` // 提供商ID
	Key                string   `bson:"key,omitempty"`         // 密钥
	Weight             int      `bson:"weight"`                // 权重
	Models             []string `bson:"models"`                // 模型
	ModelAgents        []string `bson:"model_agents"`          // 模型代理
	IsAgentsOnly       bool     `bson:"is_agents_only"`        // 是否代理专用
	IsNeverDisable     bool     `bson:"is_never_disable"`      // 是否永不禁用
	UsedQuota          int      `bson:"used_quota,omitempty"`  // 已用额度
	Remark             string   `bson:"remark"`                // 备注
	Status             int      `bson:"status,omitempty"`      // 状态[1:正常, 2:禁用, -1:删除]
	IsAutoDisabled     bool     `bson:"is_auto_disabled"`      // 是否自动禁用
	AutoDisabledReason string   `bson:"auto_disabled_reason"`  // 自动禁用原因
	Creator            string   `bson:"creator,omitempty"`     // 创建人
	Updater            string   `bson:"updater,omitempty"`     // 更新人
	CreatedAt          int64    `bson:"created_at,omitempty"`  // 创建时间
	UpdatedAt          int64    `bson:"updated_at,omitempty"`  // 更新时间
}
