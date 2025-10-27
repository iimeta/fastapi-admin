package model

// 新建密钥接口请求参数
type KeyCreateReq struct {
	ProviderId     string   `json:"provider_id,omitempty"`      // 提供商ID
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
	ProviderId     string   `json:"provider_id,omitempty"`         // 提供商ID
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
	ProviderId  string   `json:"provider_id,omitempty"`  // 提供商ID
	Key         string   `json:"key,omitempty"`          // 密钥
	Models      []string `json:"models,omitempty"`       // 模型
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt   []string `json:"created_at,omitempty"`   // 创建时间
}

// 密钥分页列表接口响应参数
type KeyPageRes struct {
	Items  []*Key  `json:"items"`
	Paging *Paging `json:"paging"`
}

// 密钥列表接口请求参数
type KeyListReq struct {
	ProviderId  string   `json:"provider_id,omitempty"`  // 提供商ID
	Key         string   `json:"key,omitempty"`          // 密钥
	Models      []string `json:"models,omitempty"`       // 模型
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
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
	ProviderId  string   `json:"provider_id,omitempty"`  // 提供商ID
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

// 密钥绑定分组接口请求参数
type KeyGroupReq struct {
	Id          string `json:"id,omitempty"`            // ID
	IsBindGroup bool   `json:"is_bind_group,omitempty"` // 是否绑定分组
	Group       string `json:"group,omitempty"`         // 绑定分组
}

type Key struct {
	Id                 string   `json:"id,omitempty"`                   // ID
	ProviderId         string   `json:"provider_id,omitempty"`          // 提供商ID
	ProviderName       string   `json:"provider_name,omitempty"`        // 提供商名称
	Key                string   `json:"key,omitempty"`                  // 密钥
	Weight             int      `json:"weight,omitempty"`               // 权重
	Models             []string `json:"models,omitempty"`               // 模型
	ModelNames         []string `json:"model_names,omitempty"`          // 模型名称
	ModelAgents        []string `json:"model_agents,omitempty"`         // 模型代理
	ModelAgentNames    []string `json:"model_agent_names,omitempty"`    // 模型代理名称
	IsAgentsOnly       bool     `json:"is_agents_only,omitempty"`       // 是否代理专用
	IsNeverDisable     bool     `json:"is_never_disable,omitempty"`     // 是否永不禁用
	UsedQuota          float64  `json:"used_quota,omitempty"`           // 已用额度
	Remark             string   `json:"remark,omitempty"`               // 备注
	Status             int      `json:"status,omitempty"`               // 状态[1:正常, 2:禁用, -1:删除]
	IsAutoDisabled     bool     `json:"is_auto_disabled,omitempty"`     // 是否自动禁用
	AutoDisabledReason string   `json:"auto_disabled_reason,omitempty"` // 自动禁用原因
	Creator            string   `json:"creator,omitempty"`              // 创建人
	Updater            string   `json:"updater,omitempty"`              // 更新人
	CreatedAt          string   `json:"created_at,omitempty"`           // 创建时间
	UpdatedAt          string   `json:"updated_at,omitempty"`           // 更新时间
}
