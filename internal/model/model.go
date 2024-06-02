package model

// 新建模型接口请求参数
type ModelCreateReq struct {
	Corp               string          `json:"corp,omitempty"`                   // 公司
	Name               string          `json:"name,omitempty"`                   // 模型名称
	Model              string          `json:"model,omitempty"`                  // 模型
	Type               int             `json:"type,omitempty"`                   // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态]
	BaseUrl            string          `json:"base_url,omitempty"`               // 模型地址
	Path               string          `json:"path,omitempty"`                   // 模型路径
	Prompt             string          `json:"prompt,omitempty"`                 // 预设提示词
	BillingMethod      int             `json:"billing_method,omitempty"`         // 计费方式[1:倍率, 2:固定额度]
	PromptRatio        float64         `json:"prompt_ratio,omitempty" d:"1"`     // 提示倍率(提问倍率)
	CompletionRatio    float64         `json:"completion_ratio,omitempty" d:"1"` // 补全倍率(回答倍率)
	FixedQuota         int             `json:"fixed_quota,omitempty"`            // 固定额度
	DataFormat         int             `json:"data_format,omitempty"`            // 数据格式[1:统一格式, 2:官方格式]
	IsPublic           bool            `json:"is_public,omitempty"`              // 是否公开
	IsEnableModelAgent bool            `json:"is_enable_model_agent,omitempty"`  // 是否启用模型代理
	ModelAgents        []string        `json:"model_agents,omitempty"`           // 模型代理
	IsEnableForward    bool            `json:"is_enable_forward,omitempty"`      // 是否启用模型转发
	ForwardConfig      *ForwardConfig  `json:"forward_config,omitempty"`         // 模型转发配置
	IsEnableFallback   bool            `json:"is_enable_fallback,omitempty"`     // 是否启用后备模型
	FallbackConfig     *FallbackConfig `json:"fallback_config,omitempty"`        // 后备模型配置
	Remark             string          `json:"remark,omitempty"`                 // 备注
	Status             int             `json:"status,omitempty" d:"1"`           // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新模型接口请求参数
type ModelUpdateReq struct {
	Id                 string          `json:"id" v:"required"`                  // ID
	Corp               string          `json:"corp,omitempty"`                   // 公司
	Name               string          `json:"name,omitempty"`                   // 模型名称
	Model              string          `json:"model,omitempty"`                  // 模型
	Type               int             `json:"type,omitempty"`                   // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态]
	BaseUrl            string          `json:"base_url,omitempty"`               // 模型地址
	Path               string          `json:"path,omitempty"`                   // 模型路径
	Prompt             string          `json:"prompt,omitempty"`                 // 预设提示词
	BillingMethod      int             `json:"billing_method,omitempty"`         // 计费方式[1:倍率, 2:固定额度]
	PromptRatio        float64         `json:"prompt_ratio,omitempty" d:"1"`     // 提示倍率(提问倍率)
	CompletionRatio    float64         `json:"completion_ratio,omitempty" d:"1"` // 补全倍率(回答倍率)
	FixedQuota         int             `json:"fixed_quota,omitempty"`            // 固定额度
	DataFormat         int             `json:"data_format,omitempty"`            // 数据格式[1:统一格式, 2:官方格式]
	IsPublic           bool            `json:"is_public,omitempty"`              // 是否公开
	IsEnableModelAgent bool            `json:"is_enable_model_agent,omitempty"`  // 是否启用模型代理
	ModelAgents        []string        `json:"model_agents,omitempty" d:"[]"`    // 模型代理
	IsEnableForward    bool            `json:"is_enable_forward,omitempty"`      // 是否启用模型转发
	ForwardConfig      *ForwardConfig  `json:"forward_config,omitempty"`         // 模型转发配置
	IsEnableFallback   bool            `json:"is_enable_fallback,omitempty"`     // 是否启用后备模型
	FallbackConfig     *FallbackConfig `json:"fallback_config,omitempty"`        // 后备模型配置
	Remark             string          `json:"remark,omitempty"`                 // 备注
	Status             int             `json:"status,omitempty" d:"1"`           // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改模型状态接口请求参数
type ModelChangeStatusReq struct {
	Id     string `json:"id" v:"required"`        // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型详情接口响应参数
type ModelDetailRes struct {
	*Model
}

// 模型分页列表接口请求参数
type ModelPageReq struct {
	Paging
	Corp            string   `json:"corp,omitempty"`             // 公司
	Name            string   `json:"name,omitempty"`             // 模型名称
	Model           string   `json:"model,omitempty"`            // 模型
	Type            int      `json:"type,omitempty"`             // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态]
	PromptRatio     float64  `json:"prompt_ratio,omitempty"`     // 提示倍率(提问倍率)
	CompletionRatio float64  `json:"completion_ratio,omitempty"` // 补全倍率(回答倍率)
	DataFormat      int      `json:"data_format,omitempty"`      // 数据格式[1:统一格式, 2:官方格式]
	IsPublic        bool     `json:"is_public"`                  // 是否公开
	Remark          string   `json:"remark,omitempty"`           // 备注
	Status          int      `json:"status,omitempty"`           // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt       []string `json:"created_at,omitempty"`       // 创建时间
}

// 模型分页列表接口响应参数
type ModelPageRes struct {
	Items  []*Model `json:"items"`
	Paging *Paging  `json:"paging"`
}

// 模型列表接口请求参数
type ModelListReq struct {
	Corp            string  `json:"corp,omitempty"`             // 公司
	Name            string  `json:"name,omitempty"`             // 模型名称
	Model           string  `json:"model,omitempty"`            // 模型
	Type            int     `json:"type,omitempty"`             // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态]
	PromptRatio     float64 `json:"prompt_ratio,omitempty"`     // 提示倍率(提问倍率)
	CompletionRatio float64 `json:"completion_ratio,omitempty"` // 补全倍率(回答倍率)
	DataFormat      int     `json:"data_format,omitempty"`      // 数据格式[1:统一格式, 2:官方格式]
	IsPublic        bool    `json:"is_public"`                  // 是否公开
	Remark          string  `json:"remark,omitempty"`           // 备注
	Status          int     `json:"status,omitempty" d:"1"`     // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型列表接口响应参数
type ModelListRes struct {
	Items []*Model `json:"items"`
}

// 模型批量操作接口请求参数
type ModelBatchOperateReq struct {
	Action      string   `json:"action"`                 // 动作
	Ids         []string `json:"ids"`                    // 主键Ids
	Value       any      `json:"value"`                  // 值
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
	TargetModel string   `json:"target_model,omitempty"` // 目标模型
}

type Model struct {
	Id                 string          `json:"id,omitempty"`                 // ID
	Corp               string          `json:"corp,omitempty"`               // 公司ID
	CorpName           string          `json:"corp_name,omitempty"`          // 公司名称
	Name               string          `json:"name,omitempty"`               // 模型名称
	Model              string          `json:"model,omitempty"`              // 模型
	Type               int             `json:"type,omitempty"`               // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态]
	BaseUrl            string          `json:"base_url,omitempty"`           // 模型地址
	Path               string          `json:"path,omitempty"`               // 模型路径
	Prompt             string          `json:"prompt,omitempty"`             // 预设提示词
	BillingMethod      int             `json:"billing_method,omitempty"`     // 计费方式[1:倍率, 2:固定额度]
	PromptRatio        float64         `json:"prompt_ratio,omitempty"`       // 提示倍率(提问倍率)
	PromptPrice        float64         `json:"prompt_price,omitempty"`       // 提示价格(提问价格)
	CompletionRatio    float64         `json:"completion_ratio,omitempty"`   // 补全倍率(回答倍率)
	CompletionPrice    float64         `json:"completion_price,omitempty"`   // 补全价格(回答价格)
	FixedQuota         int             `json:"fixed_quota"`                  // 固定额度
	FixedPrice         float64         `json:"fixed_price"`                  // 固定价格
	DataFormat         int             `json:"data_format,omitempty"`        // 数据格式[1:统一格式, 2:官方格式]
	IsPublic           bool            `json:"is_public"`                    // 是否公开
	IsEnableModelAgent bool            `json:"is_enable_model_agent"`        // 是否启用模型代理
	ModelAgents        []string        `json:"model_agents,omitempty"`       // 模型代理
	ModelAgentNames    []string        `json:"model_agent_names,omitempty"`  // 模型代理名称
	IsEnableForward    bool            `json:"is_enable_forward,omitempty"`  // 是否启用模型转发
	ForwardConfig      *ForwardConfig  `json:"forward_config,omitempty"`     // 模型转发配置
	IsEnableFallback   bool            `json:"is_enable_fallback,omitempty"` // 是否启用后备模型
	FallbackConfig     *FallbackConfig `json:"fallback_config,omitempty"`    // 后备模型配置
	Remark             string          `json:"remark,omitempty"`             // 备注
	Status             int             `json:"status,omitempty"`             // 状态[1:正常, 2:禁用, -1:删除]
	Creator            string          `json:"creator,omitempty"`            // 创建人
	Updater            string          `json:"updater,omitempty"`            // 更新人
	CreatedAt          string          `json:"created_at,omitempty"`         // 创建时间
	UpdatedAt          string          `json:"updated_at,omitempty"`         // 更新时间
}

type ForwardConfig struct {
	ForwardRule       int      `json:"forward_rule,omitempty"`        // 转发规则[1:全部转发, 2:按关键字]
	MatchRule         []int    `json:"match_rule,omitempty"`          // 匹配规则[1:智能匹配, 2:正则匹配]
	TargetModel       string   `json:"target_model,omitempty"`        // 转发规则为1时的目标模型
	TargetModelName   string   `json:"target_model_name,omitempty"`   // 转发规则为1时的目标模型名称
	DecisionModel     string   `json:"decision_model,omitempty"`      // 转发规则为2时并且匹配规则为1时的判定模型
	DecisionModelName string   `json:"decision_model_name,omitempty"` // 转发规则为2时并且匹配规则为1时的判定模型名称
	Keywords          []string `json:"keywords,omitempty"`            // 转发规则为2时的关键字
	TargetModels      []string `json:"target_models,omitempty"`       // 转发规则为2时的目标模型
	TargetModelNames  []string `json:"target_model_names,omitempty"`  // 转发规则为2时的目标模型名称
}

type FallbackConfig struct {
	FallbackModel     string `json:"fallback_model,omitempty"`      // 后备模型
	FallbackModelName string `json:"fallback_model_name,omitempty"` // 后备模型名称
}
