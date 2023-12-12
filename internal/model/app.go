package model

// 新建应用接口请求参数
type AppCreateReq struct {
	Name        string   `json:"name,omitempty"`         // 应用名称
	Type        int      `json:"type,omitempty"`         // 应用类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
}

// 更新应用接口请求参数
type AppUpdateReq struct {
	Id          string   `json:"_id,omitempty"`          // ID
	Name        string   `json:"name,omitempty"`         // 应用名称
	Type        int      `json:"type,omitempty"`         // 应用类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
}

// 应用详情接口响应参数
type AppDetailRes struct {
	*App
}

// 应用分页列表接口请求参数
type AppPageReq struct {
	Paging
	Name        string   `json:"name,omitempty"`         // 应用名称
	Type        int      `json:"type,omitempty"`         // 应用类型
	Models      []string `json:"models,omitempty"`       // 模型
	Keys        []string `json:"keys,omitempty"`         // 密钥
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常; 2:禁用; -1:删除]
}

// 应用分页列表接口响应参数
type AppPageRes struct {
	Items  []*App  `json:"items"`
	Paging *Paging `json:"paging"`
}

type App struct {
	Id          string   `json:"_id,omitempty"`          // ID
	AppId       int      `json:"app_id,omitempty"`       // 应用ID
	Name        string   `json:"name,omitempty"`         // 应用名称
	Type        int      `json:"type,omitempty"`         // 应用类型
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
