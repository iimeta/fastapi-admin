package model

// 新建密钥接口请求参数
type KeyCreateReq struct {
	Corp           string   `json:"corp,omitempty"`             // 公司
	Key            string   `json:"key,omitempty"`              // 密钥
	Weight         int      `json:"weight,omitempty"`           // 权重
	Models         []string `json:"models,omitempty"`           // 模型
	ModelAgents    []string `json:"model_agents,omitempty"`     // 模型代理
	IsAgentsOnly   bool     `json:"is_agents_only,omitempty"`   // 是否代理专用
	IsNeverDisable bool     `json:"is_never_disable,omitempty"` // 是否永不禁用
	Remark         string   `json:"remark,omitempty"`           // 备注
	Status         int      `json:"status,omitempty" d:"1"`     // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新密钥接口请求参数
type KeyUpdateReq struct {
	Id             string   `json:"id,omitempty"`                  // ID
	Corp           string   `json:"corp,omitempty"`                // 公司
	Key            string   `json:"key,omitempty"`                 // 密钥
	Weight         int      `json:"weight,omitempty"`              // 权重
	Models         []string `json:"models,omitempty" d:"[]"`       // 模型
	ModelAgents    []string `json:"model_agents,omitempty" d:"[]"` // 模型代理
	IsAgentsOnly   bool     `json:"is_agents_only,omitempty"`      // 是否代理专用
	IsNeverDisable bool     `json:"is_never_disable,omitempty"`    // 是否永不禁用
	Remark         string   `json:"remark,omitempty"`              // 备注
	Status         int      `json:"status,omitempty" d:"1"`        // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改密钥状态接口请求参数
type KeyChangeStatusReq struct {
	Id     string `json:"id,omitempty"`           // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 密钥详情接口响应参数
type KeyDetailRes struct {
	*Key
}

// 密钥分页列表接口请求参数
type KeyPageReq struct {
	Paging
	Type           int      `json:"type,omitempty"`             // 密钥类型[1:应用, 2:模型]
	UserId         int      `json:"user_id,omitempty"`          // 用户ID
	AppId          int      `json:"app_id,omitempty"`           // 应用ID
	Corp           string   `json:"corp,omitempty"`             // 公司
	Key            string   `json:"key,omitempty"`              // 密钥
	Quota          float64  `json:"quota,omitempty"`            // 额度
	QuotaExpiresAt []string `json:"quota_expires_at,omitempty"` // 额度过期时间
	Models         []string `json:"models,omitempty"`           // 模型
	ModelAgents    []string `json:"model_agents,omitempty"`     // 模型代理
	Remark         string   `json:"remark,omitempty"`           // 备注
	Status         int      `json:"status,omitempty"`           // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt      []string `json:"created_at,omitempty"`       // 创建时间
}

// 密钥分页列表接口响应参数
type KeyPageRes struct {
	Items  []*Key  `json:"items"`
	Paging *Paging `json:"paging"`
}

// 密钥列表接口请求参数
type KeyListReq struct {
	Type        int      `json:"type,omitempty"`         // 密钥类型[1:应用, 2:模型]
	AppId       int      `json:"app_id,omitempty"`       // 应用ID
	Corp        string   `json:"corp,omitempty"`         // 公司
	Key         string   `json:"key,omitempty"`          // 密钥
	Models      []string `json:"models,omitempty"`       // 模型
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
	Quota       int      `json:"quota,omitempty"`        // 额度
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
}

// 密钥列表接口响应参数
type KeyListRes struct {
	Items []*Key `json:"items"`
}

// 密钥批量操作接口请求参数
type KeyBatchOperateReq struct {
	Action      string   `json:"action"`                 // 动作
	Ids         []string `json:"ids"`                    // 主键Ids
	Value       any      `json:"value"`                  // 值
	Type        int      `json:"type,omitempty"`         // 密钥类型[1:应用, 2:模型]
	Corp        string   `json:"corp,omitempty"`         // 公司
	Key         string   `json:"key,omitempty"`          // 密钥
	Models      []string `json:"models,omitempty"`       // 模型
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
}

// 密钥模型权限接口请求参数
type KeyModelsReq struct {
	Id     string   `json:"id,omitempty"`            // ID
	Models []string `json:"models,omitempty" d:"[]"` // 模型权限
}

type Key struct {
	Id                  string   `json:"id,omitempty"`                    // ID
	UserId              int      `json:"user_id,omitempty"`               // 用户ID
	AppId               int      `json:"app_id,omitempty"`                // 应用ID
	Corp                string   `json:"corp,omitempty"`                  // 公司ID
	CorpName            string   `json:"corp_name,omitempty"`             // 公司名称
	Key                 string   `json:"key,omitempty"`                   // 密钥
	Type                int      `json:"type,omitempty"`                  // 密钥类型[1:应用, 2:模型]
	Weight              int      `json:"weight"`                          // 权重
	Models              []string `json:"models,omitempty"`                // 模型
	ModelNames          []string `json:"model_names,omitempty"`           // 模型名称
	ModelAgents         []string `json:"model_agents,omitempty"`          // 模型代理
	ModelAgentNames     []string `json:"model_agent_names,omitempty"`     // 模型代理名称
	IsAgentsOnly        bool     `json:"is_agents_only"`                  // 是否代理专用
	IsNeverDisable      bool     `json:"is_never_disable,omitempty"`      // 是否永不禁用
	IsLimitQuota        bool     `json:"is_limit_quota"`                  // 是否限制额度
	Quota               int      `json:"quota"`                           // 剩余额度
	UsedQuota           int      `json:"used_quota"`                      // 已用额度
	QuotaExpiresRule    int      `json:"quota_expires_rule,omitempty"`    // 额度过期规则[1:固定, 2:时长]
	QuotaExpiresAt      string   `json:"quota_expires_at,omitempty"`      // 额度过期时间
	QuotaExpiresMinutes int64    `json:"quota_expires_minutes,omitempty"` // 额度过期分钟数
	IpWhitelist         []string `json:"ip_whitelist,omitempty"`          // IP白名单
	IpBlacklist         []string `json:"ip_blacklist,omitempty"`          // IP黑名单
	Remark              string   `json:"remark,omitempty"`                // 备注
	Status              int      `json:"status,omitempty"`                // 状态[1:正常, 2:禁用, -1:删除]
	IsAutoDisabled      bool     `json:"is_auto_disabled,omitempty"`      // 是否自动禁用
	AutoDisabledReason  string   `json:"auto_disabled_reason,omitempty"`  // 自动禁用原因
	Creator             string   `json:"creator,omitempty"`               // 创建人
	Updater             string   `json:"updater,omitempty"`               // 更新人
	CreatedAt           string   `json:"created_at,omitempty"`            // 创建时间
	UpdatedAt           string   `json:"updated_at,omitempty"`            // 更新时间
}
