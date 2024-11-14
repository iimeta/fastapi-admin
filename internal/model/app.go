package model

// 新建应用接口请求参数
type AppCreateReq struct {
	UserId         int      `json:"user_id,omitempty"`          // 用户ID
	Name           string   `json:"name,omitempty"`             // 应用名称
	Models         []string `json:"models,omitempty"`           // 模型权限
	IsLimitQuota   bool     `json:"is_limit_quota,omitempty"`   // 是否限制额度
	Quota          int      `json:"quota,omitempty"`            // 额度
	QuotaExpiresAt string   `json:"quota_expires_at,omitempty"` // 额度过期时间
	IpWhitelist    string   `json:"ip_whitelist,omitempty"`     // IP白名单
	IpBlacklist    string   `json:"ip_blacklist,omitempty"`     // IP黑名单
	Remark         string   `json:"remark,omitempty"`           // 备注
	IsCreateKey    bool     `json:"is_create_key,omitempty"`    // 是否创建密钥
	Status         int      `json:"status,omitempty" d:"1"`     // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新应用接口请求参数
type AppUpdateReq struct {
	Id             string   `json:"id,omitempty"`               // ID
	Name           string   `json:"name,omitempty"`             // 应用名称
	Models         []string `json:"models,omitempty" d:"[]"`    // 模型权限
	IsLimitQuota   bool     `json:"is_limit_quota,omitempty"`   // 是否限制额度
	Quota          int      `json:"quota,omitempty"`            // 额度
	QuotaExpiresAt string   `json:"quota_expires_at,omitempty"` // 额度过期时间
	IpWhitelist    string   `json:"ip_whitelist,omitempty"`     // IP白名单
	IpBlacklist    string   `json:"ip_blacklist,omitempty"`     // IP黑名单
	Remark         string   `json:"remark,omitempty"`           // 备注
	Status         int      `json:"status,omitempty" d:"1"`     // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改应用状态接口请求参数
type AppChangeStatusReq struct {
	Id     string `json:"id,omitempty"`           // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 应用详情接口响应参数
type AppDetailRes struct {
	*App
}

// 应用分页列表接口请求参数
type AppPageReq struct {
	Paging
	UserId         int      `json:"user_id,omitempty"`          // 用户ID
	AppId          int      `json:"app_id,omitempty"`           // 应用ID
	Name           string   `json:"name,omitempty"`             // 应用名称
	AppKey         string   `json:"app_key,omitempty"`          // 应用密钥
	Models         []string `json:"models,omitempty"`           // 模型权限
	Quota          int      `json:"quota,omitempty"`            // 额度
	QuotaExpiresAt []string `json:"quota_expires_at,omitempty"` // 额度过期时间
	Status         int      `json:"status,omitempty"`           // 状态[1:正常, 2:禁用, -1:删除]
}

// 应用分页列表接口响应参数
type AppPageRes struct {
	Items  []*App  `json:"items"`
	Paging *Paging `json:"paging"`
}

// 应用列表接口请求参数
type AppListReq struct {
	UserId         int      `json:"user_id,omitempty"`          // 用户ID
	AppId          int      `json:"app_id,omitempty"`           // 应用ID
	Name           string   `json:"name,omitempty"`             // 应用名称
	AppKey         string   `json:"app_key,omitempty"`          // 应用密钥
	Models         []string `json:"models,omitempty"`           // 模型权限
	Quota          int      `json:"quota,omitempty"`            // 额度
	QuotaExpiresAt []string `json:"quota_expires_at,omitempty"` // 额度过期时间
	Status         int      `json:"status,omitempty"`           // 状态[1:正常, 2:禁用, -1:删除]
}

// 应用列表接口响应参数
type AppListRes struct {
	Items []*App `json:"items"`
}

// 新建应用密钥接口请求参数
type AppCreateKeyReq struct {
	UserId int `json:"user_id,omitempty"` // 用户ID
	AppId  int `json:"app_id,omitempty"`  // 应用ID
}

// 新建应用密钥接口响应参数
type AppCreateKeyRes struct {
	AppId int    `json:"app_id,omitempty"` // 应用ID
	Key   string `json:"key,omitempty"`    // 密钥
}

// 应用密钥配置接口请求参数
type AppKeyConfigReq struct {
	Id                  string   `json:"id,omitempty"`                       // ID
	UserId              int      `json:"user_id,omitempty"`                  // 用户ID
	AppId               int      `json:"app_id,omitempty"`                   // 应用ID
	Key                 string   `json:"key,omitempty"`                      // 密钥
	IsLimitQuota        bool     `json:"is_limit_quota,omitempty"`           // 是否限制额度
	Quota               int      `json:"quota,omitempty"`                    // 额度
	QuotaExpiresRule    int      `json:"quota_expires_rule,omitempty" d:"1"` // 额度过期规则[1:固定, 2:时长]
	QuotaExpiresAt      string   `json:"quota_expires_at,omitempty"`         // 额度过期时间
	QuotaExpiresMinutes int64    `json:"quota_expires_minutes,omitempty"`    // 额度过期分钟数
	Models              []string `json:"models,omitempty"`                   // 模型权限
	IpWhitelist         string   `json:"ip_whitelist,omitempty"`             // IP白名单
	IpBlacklist         string   `json:"ip_blacklist,omitempty"`             // IP黑名单
	Remark              string   `json:"remark,omitempty"`                   // 备注
	Status              int      `json:"status,omitempty" d:"1"`             // 状态[1:正常, 2:禁用, -1:删除]
}

// 应用模型权限接口请求参数
type AppModelsReq struct {
	AppId  int      `json:"app_id,omitempty"`        // 应用ID
	Models []string `json:"models,omitempty" d:"[]"` // 模型权限
}

type App struct {
	Id             string   `json:"id,omitempty"`               // ID
	AppId          int      `json:"app_id,omitempty"`           // 应用ID
	Name           string   `json:"name,omitempty"`             // 应用名称
	Models         []string `json:"models,omitempty"`           // 模型权限
	ModelNames     []string `json:"model_names,omitempty"`      // 模型名称
	IsLimitQuota   bool     `json:"is_limit_quota,omitempty"`   // 是否限制额度
	Quota          int      `json:"quota"`                      // 剩余额度
	UsedQuota      int      `json:"used_quota"`                 // 已用额度
	QuotaExpiresAt string   `json:"quota_expires_at,omitempty"` // 额度过期时间
	IpWhitelist    []string `json:"ip_whitelist,omitempty"`     // IP白名单
	IpBlacklist    []string `json:"ip_blacklist,omitempty"`     // IP黑名单
	Remark         string   `json:"remark,omitempty"`           // 备注
	Status         int      `json:"status,omitempty"`           // 状态[1:正常, 2:禁用, -1:删除]
	UserId         int      `json:"user_id,omitempty"`          // 用户ID
	Creator        string   `json:"creator,omitempty"`          // 创建人
	Updater        string   `json:"updater,omitempty"`          // 更新人
	CreatedAt      string   `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt      string   `json:"updated_at,omitempty"`       // 更新时间
}
