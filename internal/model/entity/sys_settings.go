package entity

type SysSettings struct {
	Id          string   `bson:"_id,omitempty"`          // ID
	AppId       int      `bson:"app_id,omitempty"`       // 应用ID
	Name        string   `bson:"name,omitempty"`         // 应用名称
	Type        int      `bson:"type,omitempty"`         // 应用类型
	Models      []string `bson:"models,omitempty"`       // 模型
	Keys        []string `bson:"keys,omitempty"`         // 密钥
	IpWhitelist []string `bson:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `bson:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `bson:"remark,omitempty"`       // 备注
	Status      int      `bson:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
	Creator     string   `bson:"creator,omitempty"`      // 创建人
	Updater     string   `bson:"updater,omitempty"`      // 更新人
	CreatedAt   int64    `bson:"created_at,omitempty"`   // 创建时间
	UpdatedAt   int64    `bson:"updated_at,omitempty"`   // 更新时间
}
