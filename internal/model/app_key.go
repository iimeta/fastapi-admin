package model

// 新建应用密钥接口请求参数
type AppKeyCreateReq struct {
	UserId int `json:"user_id,omitempty"` // 用户ID
	AppId  int `json:"app_id,omitempty"`  // 应用ID
}

// 新建应用密钥接口响应参数
type AppKeyCreateRes struct {
	AppId int    `json:"app_id,omitempty"` // 应用ID
	Key   string `json:"key,omitempty"`    // 密钥
}

// 应用密钥配置接口请求参数
type AppKeyConfigReq struct {
	Id                  string   `json:"id,omitempty"`                       // ID
	UserId              int      `json:"user_id,omitempty"`                  // 用户ID
	AppId               int      `json:"app_id,omitempty"`                   // 应用ID
	Key                 string   `json:"key,omitempty"`                      // 密钥
	BillingMethods      []int    `json:"billing_methods,omitempty"`          // 计费方式[1:按Tokens, 2:按次]
	Models              []string `json:"models,omitempty"`                   // 模型权限
	IsLimitQuota        bool     `json:"is_limit_quota,omitempty"`           // 是否限制额度
	Quota               int      `json:"quota,omitempty"`                    // 额度
	QuotaExpiresRule    int      `json:"quota_expires_rule,omitempty" d:"1"` // 额度过期规则[1:固定, 2:时长]
	QuotaExpiresAt      string   `json:"quota_expires_at,omitempty"`         // 额度过期时间
	QuotaExpiresMinutes int64    `json:"quota_expires_minutes,omitempty"`    // 额度过期分钟数
	IsBindGroup         bool     `json:"is_bind_group,omitempty"`            // 是否绑定分组
	Group               string   `json:"group,omitempty"`                    // 绑定分组
	IpWhitelist         string   `json:"ip_whitelist,omitempty"`             // IP白名单
	IpBlacklist         string   `json:"ip_blacklist,omitempty"`             // IP黑名单
	Remark              string   `json:"remark,omitempty"`                   // 备注
	Status              int      `json:"status,omitempty" d:"1"`             // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改应用密钥状态接口请求参数
type AppKeyChangeStatusReq struct {
	Id     string `json:"id,omitempty"`           // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 应用密钥详情接口响应参数
type AppKeyDetailRes struct {
	*AppKey
}

// 应用密钥分页列表接口请求参数
type AppKeyPageReq struct {
	Paging
	UserId         int      `json:"user_id,omitempty"`          // 用户ID
	AppId          int      `json:"app_id,omitempty"`           // 应用ID
	Key            string   `json:"key,omitempty"`              // 应用密钥
	BillingMethods []int    `json:"billing_methods,omitempty"`  // 计费方式[1:按Tokens, 2:按次]
	Quota          float64  `json:"quota,omitempty"`            // 额度
	QuotaExpiresAt []string `json:"quota_expires_at,omitempty"` // 额度过期时间
	Models         []string `json:"models,omitempty"`           // 模型
	Remark         string   `json:"remark,omitempty"`           // 备注
	Status         int      `json:"status,omitempty"`           // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt      []string `json:"created_at,omitempty"`       // 创建时间
}

// 应用密钥分页列表接口响应参数
type AppKeyPageRes struct {
	Items  []*AppKey `json:"items"`
	Paging *Paging   `json:"paging"`
}

// 应用密钥批量操作接口请求参数
type AppKeyBatchOperateReq struct {
	Action              string        `json:"action"`                             // 动作
	Ids                 []string      `json:"ids"`                                // 主键Ids
	Value               any           `json:"value"`                              // 值
	UserId              int           `json:"user_id,omitempty"`                  // 用户ID
	AppId               int           `json:"app_id,omitempty"`                   // 应用ID
	Key                 string        `json:"key,omitempty"`                      // 应用密钥
	BillingMethods      []int         `json:"billing_methods,omitempty"`          // 计费方式[1:按Tokens, 2:按次]
	N                   int           `json:"n,omitempty"`                        // 数量
	Models              []string      `json:"models,omitempty"`                   // 模型权限
	IsLimitQuota        bool          `json:"is_limit_quota,omitempty"`           // 是否限制额度
	Quota               int           `json:"quota,omitempty"`                    // 额度
	QuotaExpiresRule    int           `json:"quota_expires_rule,omitempty" d:"1"` // 额度过期规则[1:固定, 2:时长]
	QuotaExpiresAt      string        `json:"quota_expires_at,omitempty"`         // 额度过期时间
	QuotaExpiresMinutes int64         `json:"quota_expires_minutes,omitempty"`    // 额度过期分钟数
	IsBindGroup         bool          `json:"is_bind_group,omitempty"`            // 是否绑定分组
	Group               string        `json:"group,omitempty"`                    // 绑定分组
	IpWhitelist         string        `json:"ip_whitelist,omitempty"`             // IP白名单
	IpBlacklist         string        `json:"ip_blacklist,omitempty"`             // IP黑名单
	Remark              string        `json:"remark,omitempty"`                   // 备注
	Status              int           `json:"status,omitempty"`                   // 状态[1:正常, 2:禁用, -1:删除]
	QueryParams         AppKeyPageReq `json:"query_params,omitempty"`             // 查询参数
}

// 应用密钥导出接口请求参数
type AppKeyExportReq struct {
	Ids    []string `json:"ids"`               // 主键Ids
	UserId int      `json:"user_id,omitempty"` // 用户ID
	AppId  int      `json:"app_id,omitempty"`  // 应用ID
}

// 应用密钥导出
type AppKeyExport struct {
	UserId         int    `json:"user_id,omitempty"`          // 用户ID
	AppId          int    `json:"app_id,omitempty"`           // 应用ID
	AppName        string `json:"app_name,omitempty"`         // 应用名称
	Key            string `json:"key,omitempty"`              // 应用密钥
	Quota          string `json:"quota,omitempty"`            // 额度
	QuotaExpiresAt string `json:"quota_expires_at,omitempty"` // 额度过期时间
	Remark         string `json:"remark,omitempty"`           // 备注
}

// 应用密钥模型权限接口请求参数
type AppKeyModelsReq struct {
	Id     string   `json:"id,omitempty"`            // ID
	Models []string `json:"models,omitempty" d:"[]"` // 模型权限
}

// 应用密钥绑定分组接口请求参数
type AppKeyGroupReq struct {
	Id          string `json:"id,omitempty"`            // ID
	IsBindGroup bool   `json:"is_bind_group,omitempty"` // 是否绑定分组
	Group       string `json:"group,omitempty"`         // 绑定分组
}

type AppKey struct {
	Id                  string   `json:"id,omitempty"`                    // ID
	UserId              int      `json:"user_id,omitempty"`               // 用户ID
	AppId               int      `json:"app_id,omitempty"`                // 应用ID
	Key                 string   `json:"key,omitempty"`                   // 应用密钥
	BillingMethods      []int    `json:"billing_methods,omitempty"`       // 计费方式[1:按Tokens, 2:按次]
	Models              []string `json:"models,omitempty"`                // 模型权限
	ModelNames          []string `json:"model_names,omitempty"`           // 模型名称
	IsLimitQuota        bool     `json:"is_limit_quota"`                  // 是否限制额度
	Quota               int      `json:"quota"`                           // 剩余额度
	UsedQuota           int      `json:"used_quota"`                      // 已用额度
	QuotaExpiresRule    int      `json:"quota_expires_rule,omitempty"`    // 额度过期规则[1:固定, 2:时长]
	QuotaExpiresAt      string   `json:"quota_expires_at,omitempty"`      // 额度过期时间
	QuotaExpiresMinutes int64    `json:"quota_expires_minutes,omitempty"` // 额度过期分钟数
	IsBindGroup         bool     `json:"is_bind_group"`                   // 是否绑定分组
	Group               string   `json:"group,omitempty"`                 // 绑定分组
	GroupName           string   `json:"group_name,omitempty"`            // 绑定分组名称
	IpWhitelist         []string `json:"ip_whitelist,omitempty"`          // IP白名单
	IpBlacklist         []string `json:"ip_blacklist,omitempty"`          // IP黑名单
	Remark              string   `json:"remark,omitempty"`                // 备注
	Status              int      `json:"status,omitempty"`                // 状态[1:正常, 2:禁用, -1:删除]
	Creator             string   `json:"creator,omitempty"`               // 创建人
	Updater             string   `json:"updater,omitempty"`               // 更新人
	CreatedAt           string   `json:"created_at,omitempty"`            // 创建时间
	UpdatedAt           string   `json:"updated_at,omitempty"`            // 更新时间
}
