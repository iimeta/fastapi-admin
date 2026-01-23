package model

// 新建模型代理接口请求参数
type ModelAgentCreateReq struct {
	ProviderId           string   `json:"provider_id,omitempty"`             // 提供商ID
	Name                 string   `json:"name,omitempty"`                    // 模型代理名称
	BaseUrl              string   `json:"base_url,omitempty"`                // 模型代理地址
	Path                 string   `json:"path,omitempty"`                    // 模型代理地址路径
	Weight               int      `json:"weight,omitempty"`                  // 权重
	Groups               []string `json:"groups,omitempty"`                  // 绑定分组
	Models               []string `json:"models,omitempty" d:"[]"`           // 绑定模型
	IsEnableModelReplace bool     `json:"is_enable_model_replace,omitempty"` // 是否启用模型替换
	ReplaceModels        []string `json:"replace_models,omitempty"`          // 替换模型
	TargetModels         []string `json:"target_models,omitempty"`           // 目标模型
	IsNeverDisable       bool     `json:"is_never_disable,omitempty"`        // 是否永不禁用
	LbStrategy           int      `json:"lb_strategy,omitempty" d:"1"`       // 密钥负载均衡策略[1:轮询, 2:权重]
	Key                  string   `json:"key,omitempty"`                     // 密钥
	IsAgentsOnly         bool     `json:"is_agents_only,omitempty"`          // 是否代理专用
	IsNeverDisableKey    bool     `json:"is_never_disable_key,omitempty"`    // 是否永不禁用密钥
	Remark               string   `json:"remark,omitempty"`                  // 备注
	Status               int      `json:"status,omitempty" d:"1"`            // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新模型代理接口请求参数
type ModelAgentUpdateReq struct {
	Id                   string   `json:"id" v:"required"`                   // ID
	ProviderId           string   `json:"provider_id,omitempty"`             // 提供商ID
	Name                 string   `json:"name,omitempty"`                    // 模型代理名称
	BaseUrl              string   `json:"base_url,omitempty"`                // 模型代理地址
	Path                 string   `json:"path,omitempty"`                    // 模型代理地址路径
	Weight               int      `json:"weight,omitempty"`                  // 权重
	Groups               []string `json:"groups,omitempty"`                  // 绑定分组
	Models               []string `json:"models,omitempty" d:"[]"`           // 绑定模型
	IsEnableModelReplace bool     `json:"is_enable_model_replace,omitempty"` // 是否启用模型替换
	ReplaceModels        []string `json:"replace_models,omitempty"`          // 替换模型
	TargetModels         []string `json:"target_models,omitempty"`           // 目标模型
	IsNeverDisable       bool     `json:"is_never_disable,omitempty"`        // 是否永不禁用
	LbStrategy           int      `json:"lb_strategy,omitempty" d:"1"`       // 密钥负载均衡策略[1:轮询, 2:权重]
	Key                  string   `json:"key,omitempty"`                     // 密钥
	IsAgentsOnly         bool     `json:"is_agents_only,omitempty"`          // 是否代理专用
	IsNeverDisableKey    bool     `json:"is_never_disable_key,omitempty"`    // 是否永不禁用密钥
	Remark               string   `json:"remark,omitempty"`                  // 备注
	Status               int      `json:"status,omitempty" d:"1"`            // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改模型代理状态接口请求参数
type ModelAgentChangeStatusReq struct {
	Id     string `json:"id" v:"required"`        // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型代理详情接口响应参数
type ModelAgentDetailRes struct {
	*ModelAgent
}

// 模型代理分页列表接口请求参数
type ModelAgentPageReq struct {
	Paging
	ProviderId string   `json:"provider_id,omitempty"` // 提供商ID
	Name       string   `json:"name,omitempty"`        // 模型代理名称
	BaseUrl    string   `json:"base_url,omitempty"`    // 模型代理地址
	Path       string   `json:"path,omitempty"`        // 模型代理地址路径
	Models     []string `json:"models,omitempty"`      // 绑定模型
	Remark     string   `json:"remark,omitempty"`      // 备注
	Status     int      `json:"status,omitempty"`      // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt  []string `json:"created_at,omitempty"`  // 创建时间
}

// 模型代理分页列表接口响应参数
type ModelAgentPageRes struct {
	Items  []*ModelAgent `json:"items"`
	Paging *Paging       `json:"paging"`
}

// 模型代理列表接口请求参数
type ModelAgentListReq struct {
	ProviderId string `json:"provider_id,omitempty"`  // 提供商ID
	Name       string `json:"name,omitempty"`         // 模型代理名称
	Status     int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型代理列表接口响应参数
type ModelAgentListRes struct {
	Items []*ModelAgent `json:"items"`
}

// 模型代理批量操作接口请求参数
type ModelAgentBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

// 快速填入模型接口请求参数
type ModelAgentQuickFillModelReq struct {
	BaseUrl string `json:"base_url,omitempty"` // 模型代理地址
	Key     string `json:"key,omitempty"`      // 密钥
}

// 快速填入模型接口响应参数
type ModelAgentQuickFillModelRes struct {
	Models []string `json:"models,omitempty"` // 模型
}

type ModelAgent struct {
	Id                   string   `json:"id,omitempty"`                      // ID
	ProviderId           string   `json:"provider_id,omitempty"`             // 提供商ID
	ProviderName         string   `json:"provider_name,omitempty"`           // 提供商名称
	Name                 string   `json:"name,omitempty"`                    // 模型代理名称
	BaseUrl              string   `json:"base_url,omitempty"`                // 模型代理地址
	Path                 string   `json:"path,omitempty"`                    // 模型代理地址路径
	Weight               int      `json:"weight,omitempty"`                  // 权重
	Groups               []string `json:"groups,omitempty"`                  // 绑定分组
	GroupNames           []string `json:"group_names,omitempty"`             // 分组名称
	Models               []string `json:"models,omitempty"`                  // 绑定模型
	ModelNames           []string `json:"model_names,omitempty"`             // 模型名称
	FallbackModels       []string `json:"fallback_models,omitempty"`         // 后备模型
	FallbackModelNames   []string `json:"fallback_model_names,omitempty"`    // 后备模型名称
	IsEnableModelReplace bool     `json:"is_enable_model_replace,omitempty"` // 是否启用模型替换
	ReplaceModels        []string `json:"replace_models,omitempty"`          // 替换模型
	TargetModels         []string `json:"target_models,omitempty"`           // 目标模型
	IsNeverDisable       bool     `json:"is_never_disable,omitempty"`        // 是否永不禁用
	LbStrategy           int      `json:"lb_strategy,omitempty"`             // 密钥负载均衡策略[1:轮询, 2:权重]
	Key                  string   `json:"key,omitempty"`                     // 密钥
	Remark               string   `json:"remark,omitempty"`                  // 备注
	Status               int      `json:"status,omitempty"`                  // 状态[1:正常, 2:禁用, -1:删除]
	IsAutoDisabled       bool     `json:"is_auto_disabled,omitempty"`        // 是否自动禁用
	AutoDisabledReason   string   `json:"auto_disabled_reason,omitempty"`    // 自动禁用原因
	Creator              string   `json:"creator,omitempty"`                 // 创建人
	Updater              string   `json:"updater,omitempty"`                 // 更新人
	CreatedAt            string   `json:"created_at,omitempty"`              // 创建时间
	UpdatedAt            string   `json:"updated_at,omitempty"`              // 更新时间
}
