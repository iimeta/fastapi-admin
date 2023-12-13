package model

// 新建菜单接口请求参数
type SysMenuCreateReq struct {
	Name        string   `json:"name,omitempty"`         // 菜单名称
	Type        int      `json:"type,omitempty"`         // 菜单类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
}

// 更新菜单接口请求参数
type SysMenuUpdateReq struct {
	Id          string   `json:"_id,omitempty"`          // ID
	Name        string   `json:"name,omitempty"`         // 菜单名称
	Type        int      `json:"type,omitempty"`         // 菜单类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
}

// 菜单详情接口响应参数
type SysMenuDetailRes struct {
	*SysMenu
}

// 菜单分页列表接口请求参数
type SysMenuPageReq struct {
	Paging
	Name        string   `json:"name,omitempty"`         // 菜单名称
	Type        int      `json:"type,omitempty"`         // 菜单类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
}

// 菜单分页列表接口响应参数
type SysMenuPageRes struct {
	Items  []*SysMenu `json:"items"`
	Paging *Paging    `json:"paging"`
}

type SysMenu struct {
	Id          string   `json:"_id,omitempty"`          // ID
	MenuId      int      `json:"app_id,omitempty"`       // 菜单ID
	Name        string   `json:"name,omitempty"`         // 菜单名称
	Type        int      `json:"type,omitempty"`         // 菜单类型
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
