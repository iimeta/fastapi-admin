package model

import "github.com/iimeta/fastapi-admin/internal/model/common"

// 新建模型接口请求参数
type ModelCreateReq struct {
	ProviderId           string                 `json:"provider_id,omitempty"`             // 提供商ID
	Name                 string                 `json:"name,omitempty"`                    // 模型名称
	Model                string                 `json:"model,omitempty"`                   // 模型
	Type                 int                    `json:"type,omitempty"`                    // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 7:文本向量化, 8:文生视频, 100:多模态, 101:多模态实时, 102:多模态语音, 103:多模态向量化]
	BaseUrl              string                 `json:"base_url,omitempty"`                // 模型地址
	Path                 string                 `json:"path,omitempty"`                    // 模型路径
	IsEnablePresetConfig bool                   `json:"is_enable_preset_config,omitempty"` // 是否启用预设配置
	PresetConfig         common.PresetConfig    `json:"preset_config,omitempty"`           // 预设配置
	Pricing              common.Pricing         `json:"pricing,omitempty"`                 // 定价
	RequestDataFormat    int                    `json:"request_data_format,omitempty"`     // 请求数据格式[1:统一格式, 2:官方格式]
	ResponseDataFormat   int                    `json:"response_data_format,omitempty"`    // 响应数据格式[1:统一格式, 2:官方格式]
	IsPublic             bool                   `json:"is_public,omitempty"`               // 是否公开
	Groups               []string               `json:"groups,omitempty"`                  // 分组权限
	IsEnableModelAgent   bool                   `json:"is_enable_model_agent,omitempty"`   // 是否启用模型代理
	LbStrategy           int                    `json:"lb_strategy,omitempty" d:"1"`       // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents          []string               `json:"model_agents,omitempty" d:"[]"`     // 模型代理
	IsEnableForward      bool                   `json:"is_enable_forward,omitempty"`       // 是否启用模型转发
	ForwardConfig        *common.ForwardConfig  `json:"forward_config,omitempty"`          // 模型转发配置
	IsEnableFallback     bool                   `json:"is_enable_fallback,omitempty"`      // 是否启用后备
	FallbackConfig       *common.FallbackConfig `json:"fallback_config,omitempty"`         // 后备配置
	Remark               string                 `json:"remark,omitempty"`                  // 备注
	Status               int                    `json:"status,omitempty" d:"1"`            // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新模型接口请求参数
type ModelUpdateReq struct {
	Id                   string                 `json:"id" v:"required"`                   // ID
	ProviderId           string                 `json:"provider_id,omitempty"`             // 提供商ID
	Name                 string                 `json:"name,omitempty"`                    // 模型名称
	Model                string                 `json:"model,omitempty"`                   // 模型
	Type                 int                    `json:"type,omitempty"`                    // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 7:文本向量化, 8:文生视频, 100:多模态, 101:多模态实时, 102:多模态语音, 103:多模态向量化]
	BaseUrl              string                 `json:"base_url,omitempty"`                // 模型地址
	Path                 string                 `json:"path,omitempty"`                    // 模型路径
	IsEnablePresetConfig bool                   `json:"is_enable_preset_config,omitempty"` // 是否启用预设配置
	PresetConfig         common.PresetConfig    `json:"preset_config,omitempty"`           // 预设配置
	Pricing              common.Pricing         `json:"pricing,omitempty"`                 // 定价
	RequestDataFormat    int                    `json:"request_data_format,omitempty"`     // 请求数据格式[1:统一格式, 2:官方格式]
	ResponseDataFormat   int                    `json:"response_data_format,omitempty"`    // 响应数据格式[1:统一格式, 2:官方格式]
	IsPublic             bool                   `json:"is_public,omitempty"`               // 是否公开
	Groups               []string               `json:"groups,omitempty"`                  // 分组权限
	IsEnableModelAgent   bool                   `json:"is_enable_model_agent,omitempty"`   // 是否启用模型代理
	LbStrategy           int                    `json:"lb_strategy,omitempty" d:"1"`       // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents          []string               `json:"model_agents,omitempty" d:"[]"`     // 模型代理
	IsEnableForward      bool                   `json:"is_enable_forward,omitempty"`       // 是否启用模型转发
	ForwardConfig        *common.ForwardConfig  `json:"forward_config,omitempty"`          // 模型转发配置
	IsEnableFallback     bool                   `json:"is_enable_fallback,omitempty"`      // 是否启用后备
	FallbackConfig       *common.FallbackConfig `json:"fallback_config,omitempty"`         // 后备配置
	Remark               string                 `json:"remark,omitempty"`                  // 备注
	Status               int                    `json:"status,omitempty" d:"1"`            // 状态[1:正常, 2:禁用, -1:删除]
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
	ProviderId string   `json:"provider_id,omitempty"` // 提供商ID
	Name       string   `json:"name,omitempty"`        // 模型名称
	Model      string   `json:"model,omitempty"`       // 模型
	Type       int      `json:"type,omitempty"`        // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 7:文本向量化, 8:文生视频, 100:多模态, 101:多模态实时, 102:多模态语音, 103:多模态向量化]
	Group      string   `json:"group,omitempty"`       // 分组
	Remark     string   `json:"remark,omitempty"`      // 备注
	Status     int      `json:"status,omitempty"`      // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt  []string `json:"created_at,omitempty"`  // 创建时间
}

// 模型分页列表接口响应参数
type ModelPageRes struct {
	Items  []*Model `json:"items"`
	Paging *Paging  `json:"paging"`
}

// 模型列表接口请求参数
type ModelListReq struct {
	ProviderId string   `json:"provider_id,omitempty"` // 提供商ID
	Name       string   `json:"name,omitempty"`        // 模型名称
	Model      string   `json:"model,omitempty"`       // 模型
	Type       int      `json:"type,omitempty"`        // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 7:文本向量化, 8:文生视频, 100:多模态, 101:多模态实时, 102:多模态语音, 103:多模态向量化]
	Status     int      `json:"status,omitempty"`      // 状态[1:正常, 2:禁用, -1:删除]
	Models     []string `json:"models,omitempty"`      // 模型权限
}

// 模型列表接口响应参数
type ModelListRes struct {
	Items []*Model `json:"items"`
}

// 模型批量操作接口请求参数
type ModelBatchOperateReq struct {
	Action         string                 `json:"action"`                      // 动作
	Ids            []string               `json:"ids"`                         // 主键Ids
	Value          any                    `json:"value"`                       // 值
	LbStrategy     int                    `json:"lb_strategy,omitempty" d:"1"` // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents    []string               `json:"model_agents,omitempty"`      // 模型代理
	TargetModel    string                 `json:"target_model,omitempty"`      // 目标模型
	FallbackConfig *common.FallbackConfig `json:"fallback_config,omitempty"`   // 后备配置
}

// 模型树接口请求参数
type ModelTreeReq struct {
	ProviderId string `json:"provider_id,omitempty"`  // 提供商ID
	Name       string `json:"name,omitempty"`         // 模型名称
	Model      string `json:"model,omitempty"`        // 模型
	Type       int    `json:"type,omitempty"`         // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 7:文本向量化, 8:文生视频, 100:多模态, 101:多模态实时, 102:多模态语音, 103:多模态向量化]
	Status     int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型树接口响应参数
type ModelTreeRes struct {
	Items []*Tree `json:"items"`
}

// 模型权限列表接口请求参数
type ModelPermissionsReq struct {
	Id         string `json:"id"`          // 主键Id
	Action     string `json:"action"`      // 动作
	ProviderId string `json:"provider_id"` // 提供商ID
	Name       string `json:"name"`        // 模型名称
	Model      string `json:"model"`       // 模型
	Type       int    `json:"type"`        // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 7:文本向量化, 8:文生视频, 100:多模态, 101:多模态实时, 102:多模态语音, 103:多模态向量化]
	Status     int    `json:"status"`      // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型权限列表接口响应参数
type ModelPermissionsRes struct {
	Items []*Model `json:"items"`
}

// 模型初始化同步接口请求参数
type ModelInitSyncReq struct {
	Url                string `json:"url"`                   // Fast API 模型接口
	Key                string `json:"key"`                   // Fast API 应用密钥
	IsConfigModelAgent bool   `json:"is_config_model_agent"` // 是否配置模型代理
	IsCoverPrice       bool   `json:"is_cover_price"`        // 是否覆盖价格
}

// Models接口响应参数
type ModelsRes struct {
	Object string       `json:"object"`
	Data   []ModelsData `json:"data"`
}

type ModelsData struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	OwnedBy string   `json:"owned_by"`
	Created int      `json:"created"`
	FastAPI *FastAPI `json:"fastapi,omitempty"`
}

type FastAPI struct {
	Provider string         `json:"provider,omitempty"` // 提供商名称
	Code     string         `json:"code,omitempty"`     // 提供商代码
	Model    string         `json:"model,omitempty"`    // 模型
	Type     int            `json:"type,omitempty"`     // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 7:文本向量化, 8:文生视频, 100:多模态, 101:多模态实时, 102:多模态语音, 103:多模态向量化]
	BaseUrl  string         `json:"base_url,omitempty"` // 模型地址
	Path     string         `json:"path,omitempty"`     // 模型路径
	Pricing  common.Pricing `json:"pricing,omitempty"`  // 定价
	Remark   string         `json:"remark,omitempty"`   // 备注
}

type Model struct {
	Id                   string                 `json:"id,omitempty"`                      // ID
	ProviderId           string                 `json:"provider_id,omitempty"`             // 提供商ID
	ProviderName         string                 `json:"provider_name,omitempty"`           // 提供商名称
	ProviderCode         string                 `json:"provider_code,omitempty"`           // 提供商代码
	Name                 string                 `json:"name,omitempty"`                    // 模型名称
	Model                string                 `json:"model,omitempty"`                   // 模型
	Type                 int                    `json:"type,omitempty"`                    // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 7:文本向量化, 8:文生视频, 100:多模态, 101:多模态实时, 102:多模态语音, 103:多模态向量化]
	BaseUrl              string                 `json:"base_url,omitempty"`                // 模型地址
	Path                 string                 `json:"path,omitempty"`                    // 模型路径
	Groups               []string               `json:"groups,omitempty"`                  // 分组权限
	GroupNames           []string               `json:"group_names,omitempty"`             // 分组名称
	IsEnablePresetConfig bool                   `json:"is_enable_preset_config,omitempty"` // 是否启用预设配置
	PresetConfig         common.PresetConfig    `json:"preset_config,omitempty"`           // 预设配置
	Pricing              common.Pricing         `json:"pricing,omitempty"`                 // 定价
	RequestDataFormat    int                    `json:"request_data_format,omitempty"`     // 请求数据格式[1:统一格式, 2:官方格式]
	ResponseDataFormat   int                    `json:"response_data_format,omitempty"`    // 响应数据格式[1:统一格式, 2:官方格式]
	IsPublic             bool                   `json:"is_public"`                         // 是否公开
	IsEnableModelAgent   bool                   `json:"is_enable_model_agent"`             // 是否启用模型代理
	LbStrategy           int                    `json:"lb_strategy,omitempty"`             // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents          []string               `json:"model_agents,omitempty"`            // 模型代理
	ModelAgentNames      []string               `json:"model_agent_names,omitempty"`       // 模型代理名称
	IsEnableForward      bool                   `json:"is_enable_forward,omitempty"`       // 是否启用模型转发
	ForwardConfig        *common.ForwardConfig  `json:"forward_config,omitempty"`          // 模型转发配置
	IsEnableFallback     bool                   `json:"is_enable_fallback,omitempty"`      // 是否启用后备
	FallbackConfig       *common.FallbackConfig `json:"fallback_config,omitempty"`         // 后备配置
	Remark               string                 `json:"remark,omitempty"`                  // 备注
	Status               int                    `json:"status,omitempty"`                  // 状态[1:正常, 2:禁用, -1:删除]
	Creator              string                 `json:"creator,omitempty"`                 // 创建人
	Updater              string                 `json:"updater,omitempty"`                 // 更新人
	CreatedAt            string                 `json:"created_at,omitempty"`              // 创建时间
	UpdatedAt            string                 `json:"updated_at,omitempty"`              // 更新时间
}
