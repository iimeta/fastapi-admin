package model

// 新建角色接口请求参数
type RoleCreateReq struct {
	Name        string   `json:"name,omitempty"`         // 角色名称
	Type        int      `json:"type,omitempty"`         // 角色类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
}

// 更新角色接口请求参数
type RoleUpdateReq struct {
	Id          string   `json:"_id,omitempty"`          // ID
	Name        string   `json:"name,omitempty"`         // 角色名称
	Type        int      `json:"type,omitempty"`         // 角色类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
}

// 角色详情接口响应参数
type RoleDetailRes struct {
	*Role
}

// 角色分页列表接口请求参数
type RolePageReq struct {
	Paging
	Name        string   `json:"name,omitempty"`         // 角色名称
	Type        int      `json:"type,omitempty"`         // 角色类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
}

// 角色分页列表接口响应参数
type RolePageRes struct {
	Items  []*Role `json:"items"`
	Paging *Paging `json:"paging"`
}

type Role struct {
	Id          string   `json:"_id,omitempty"`          // ID
	RoleId      int      `json:"app_id,omitempty"`       // 角色ID
	Name        string   `json:"name,omitempty"`         // 角色名称
	Type        int      `json:"type,omitempty"`         // 角色类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
	Creator     string   `json:"creator,omitempty"`      // 创建人
	Updater     string   `json:"updater,omitempty"`      // 更新人
	CreatedAt   int64    `json:"created_at,omitempty"`   // 创建时间
	UpdatedAt   int64    `json:"updated_at,omitempty"`   // 更新时间
}
