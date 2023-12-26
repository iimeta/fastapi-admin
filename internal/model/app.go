package model

// 新建应用接口请求参数
type AppCreateReq struct {
	Name        string   `json:"name,omitempty"`         // 应用名称
	Type        int      `json:"type,omitempty"`         // 应用类型
	Models      []string `json:"models,omitempty"`       // 模型权限
	Quota       int      `json:"quota,omitempty"`        // 额度
	IpWhitelist string   `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist string   `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新应用接口请求参数
type AppUpdateReq struct {
	Id          string   `json:"id,omitempty"`           // ID
	Name        string   `json:"name,omitempty"`         // 应用名称
	Type        int      `json:"type,omitempty"`         // 应用类型
	Models      []string `json:"models,omitempty"`       // 模型权限
	Quota       int      `json:"quota,omitempty"`        // 额度
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
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
	Models      []string `json:"models,omitempty"`       // 模型权限
	Quota       int      `json:"quota,omitempty"`        // 额度
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt   []string `json:"created_at,omitempty"`   // 创建时间
}

// 应用分页列表接口响应参数
type AppPageRes struct {
	Items  []*App  `json:"items"`
	Paging *Paging `json:"paging"`
}

// 应用列表接口请求参数
type AppListReq struct {
	Name        string   `json:"name,omitempty"`         // 应用名称
	Type        int      `json:"type,omitempty"`         // 应用类型
	Models      []string `json:"models,omitempty"`       // 模型权限
	Quota       int      `json:"quota,omitempty"`        // 额度
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
}

// 应用列表接口响应参数
type AppListRes struct {
	Items []*App `json:"items"`
}

// 新建应用密钥接口请求参数
type AppCreateKeyReq struct {
	AppId int `json:"app_id,omitempty"` // 应用ID
}

// 新建应用密钥接口响应参数
type AppCreateKeyRes struct {
	AppId int    `json:"app_id,omitempty"` // 应用ID
	Key   string `json:"key,omitempty"`    // 密钥
}

// 应用密钥配置接口请求参数
type AppKeyConfigReq struct {
	Id          string   `json:"id,omitempty"`           // ID
	AppId       int      `json:"app_id,omitempty"`       // 应用ID
	Key         string   `json:"key,omitempty"`          // 密钥
	Quota       int      `json:"quota,omitempty"`        // 额度
	Models      []string `json:"models,omitempty"`       // 模型权限
	IpWhitelist string   `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist string   `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

type App struct {
	Id          string   `json:"id,omitempty"`           // ID
	AppId       int      `json:"app_id,omitempty"`       // 应用ID
	Name        string   `json:"name,omitempty"`         // 应用名称
	Type        int      `json:"type,omitempty"`         // 应用类型
	Models      []string `json:"models,omitempty"`       // 模型权限
	Quota       int      `json:"quota,omitempty"`        // 额度
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
	Creator     string   `json:"creator,omitempty"`      // 创建人
	Updater     string   `json:"updater,omitempty"`      // 更新人
	CreatedAt   string   `json:"created_at,omitempty"`   // 创建时间
	UpdatedAt   string   `json:"updated_at,omitempty"`   // 更新时间
}
