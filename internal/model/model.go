package model

import "github.com/iimeta/fastapi-admin/internal/model/common"

// 新建模型接口请求参数
type ModelCreateReq struct {
	Corp                 string                      `json:"corp,omitempty"`                    // 公司
	Name                 string                      `json:"name,omitempty"`                    // 模型名称
	Model                string                      `json:"model,omitempty"`                   // 模型
	Type                 int                         `json:"type,omitempty"`                    // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
	BaseUrl              string                      `json:"base_url,omitempty"`                // 模型地址
	Path                 string                      `json:"path,omitempty"`                    // 模型路径
	IsEnablePresetConfig bool                        `json:"is_enable_preset_config,omitempty"` // 是否启用预设配置
	PresetConfig         common.PresetConfig         `json:"preset_config,omitempty"`           // 预设配置
	TextQuota            common.TextQuota            `json:"text_quota,omitempty"`              // 文本额度
	ImageQuotas          []common.ImageQuota         `json:"image_quotas,omitempty"`            // 图像额度
	AudioQuota           common.AudioQuota           `json:"audio_quota,omitempty"`             // 音频额度
	MultimodalQuota      common.MultimodalQuota      `json:"multimodal_quota,omitempty"`        // 多模态额度
	RealtimeQuota        common.RealtimeQuota        `json:"realtime_quota,omitempty"`          // 多模态实时额度
	MultimodalAudioQuota common.MultimodalAudioQuota `json:"multimodal_audio_quota,omitempty"`  // 多模态语音额度
	MidjourneyQuotas     []common.MidjourneyQuota    `json:"midjourney_quotas,omitempty"`       // Midjourney额度
	DataFormat           int                         `json:"data_format,omitempty"`             // 数据格式[1:统一格式, 2:官方格式]
	IsPublic             bool                        `json:"is_public,omitempty"`               // 是否公开
	IsEnableModelAgent   bool                        `json:"is_enable_model_agent,omitempty"`   // 是否启用模型代理
	LbStrategy           int                         `json:"lb_strategy,omitempty" d:"1"`       // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents          []string                    `json:"model_agents,omitempty" d:"[]"`     // 模型代理
	IsEnableForward      bool                        `json:"is_enable_forward,omitempty"`       // 是否启用模型转发
	ForwardConfig        *common.ForwardConfig       `json:"forward_config,omitempty"`          // 模型转发配置
	IsEnableFallback     bool                        `json:"is_enable_fallback,omitempty"`      // 是否启用后备
	FallbackConfig       *common.FallbackConfig      `json:"fallback_config,omitempty"`         // 后备配置
	Remark               string                      `json:"remark,omitempty"`                  // 备注
	Status               int                         `json:"status,omitempty" d:"1"`            // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新模型接口请求参数
type ModelUpdateReq struct {
	Id                   string                      `json:"id" v:"required"`                   // ID
	Corp                 string                      `json:"corp,omitempty"`                    // 公司
	Name                 string                      `json:"name,omitempty"`                    // 模型名称
	Model                string                      `json:"model,omitempty"`                   // 模型
	Type                 int                         `json:"type,omitempty"`                    // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
	BaseUrl              string                      `json:"base_url,omitempty"`                // 模型地址
	Path                 string                      `json:"path,omitempty"`                    // 模型路径
	IsEnablePresetConfig bool                        `json:"is_enable_preset_config,omitempty"` // 是否启用预设配置
	PresetConfig         common.PresetConfig         `json:"preset_config,omitempty"`           // 预设配置
	TextQuota            common.TextQuota            `json:"text_quota,omitempty"`              // 文本额度
	ImageQuotas          []common.ImageQuota         `json:"image_quotas,omitempty"`            // 图像额度
	AudioQuota           common.AudioQuota           `json:"audio_quota,omitempty"`             // 音频额度
	MultimodalQuota      common.MultimodalQuota      `json:"multimodal_quota,omitempty"`        // 多模态额度
	RealtimeQuota        common.RealtimeQuota        `json:"realtime_quota,omitempty"`          // 多模态实时额度
	MultimodalAudioQuota common.MultimodalAudioQuota `json:"multimodal_audio_quota,omitempty"`  // 多模态语音额度
	MidjourneyQuotas     []common.MidjourneyQuota    `json:"midjourney_quotas,omitempty"`       // Midjourney额度
	DataFormat           int                         `json:"data_format,omitempty"`             // 数据格式[1:统一格式, 2:官方格式]
	IsPublic             bool                        `json:"is_public,omitempty"`               // 是否公开
	IsEnableModelAgent   bool                        `json:"is_enable_model_agent,omitempty"`   // 是否启用模型代理
	LbStrategy           int                         `json:"lb_strategy,omitempty" d:"1"`       // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents          []string                    `json:"model_agents,omitempty" d:"[]"`     // 模型代理
	IsEnableForward      bool                        `json:"is_enable_forward,omitempty"`       // 是否启用模型转发
	ForwardConfig        *common.ForwardConfig       `json:"forward_config,omitempty"`          // 模型转发配置
	IsEnableFallback     bool                        `json:"is_enable_fallback,omitempty"`      // 是否启用后备
	FallbackConfig       *common.FallbackConfig      `json:"fallback_config,omitempty"`         // 后备配置
	Remark               string                      `json:"remark,omitempty"`                  // 备注
	Status               int                         `json:"status,omitempty" d:"1"`            // 状态[1:正常, 2:禁用, -1:删除]
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
	Corp      string   `json:"corp,omitempty"`       // 公司
	Name      string   `json:"name,omitempty"`       // 模型名称
	Model     string   `json:"model,omitempty"`      // 模型
	Type      int      `json:"type,omitempty"`       // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
	Remark    string   `json:"remark,omitempty"`     // 备注
	Status    int      `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
}

// 模型分页列表接口响应参数
type ModelPageRes struct {
	Items  []*Model `json:"items"`
	Paging *Paging  `json:"paging"`
}

// 模型列表接口请求参数
type ModelListReq struct {
	Corp   string `json:"corp,omitempty"`         // 公司
	Name   string `json:"name,omitempty"`         // 模型名称
	Model  string `json:"model,omitempty"`        // 模型
	Type   int    `json:"type,omitempty"`         // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型列表接口响应参数
type ModelListRes struct {
	Items []*Model `json:"items"`
}

// 模型批量操作接口请求参数
type ModelBatchOperateReq struct {
	Action         string                 `json:"action"`                    // 动作
	Ids            []string               `json:"ids"`                       // 主键Ids
	Value          any                    `json:"value"`                     // 值
	ModelAgents    []string               `json:"model_agents,omitempty"`    // 模型代理
	TargetModel    string                 `json:"target_model,omitempty"`    // 目标模型
	FallbackConfig *common.FallbackConfig `json:"fallback_config,omitempty"` // 后备配置
}

// 模型树接口请求参数
type ModelTreeReq struct {
	Corp   string `json:"corp,omitempty"`         // 公司
	Name   string `json:"name,omitempty"`         // 模型名称
	Model  string `json:"model,omitempty"`        // 模型
	Type   int    `json:"type,omitempty"`         // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型树接口响应参数
type ModelTreeRes struct {
	Items []*Tree `json:"items"`
}

// 模型初始化同步接口请求参数
type ModelInitSyncReq struct {
	Url                string `json:"url"`                   // Fast API 模型接口
	Key                string `json:"key"`                   // Fast API 应用密钥
	IsConfigModelAgent bool   `json:"is_config_model_agent"` // 是否配置模型代理
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
	Corp                 string                      `json:"corp,omitempty"`                   // 公司名称
	Code                 string                      `json:"code,omitempty"`                   // 公司代码
	Model                string                      `json:"model,omitempty"`                  // 模型
	Type                 int                         `json:"type,omitempty"`                   // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
	BaseUrl              string                      `json:"base_url,omitempty"`               // 模型地址
	Path                 string                      `json:"path,omitempty"`                   // 模型路径
	TextQuota            common.TextQuota            `json:"text_quota,omitempty"`             // 文本额度
	ImageQuotas          []common.ImageQuota         `json:"image_quotas,omitempty"`           // 图像额度
	AudioQuota           common.AudioQuota           `json:"audio_quota,omitempty"`            // 音频额度
	MultimodalQuota      common.MultimodalQuota      `json:"multimodal_quota,omitempty"`       // 多模态额度
	RealtimeQuota        common.RealtimeQuota        `json:"realtime_quota,omitempty"`         // 多模态实时额度
	MultimodalAudioQuota common.MultimodalAudioQuota `json:"multimodal_audio_quota,omitempty"` // 多模态语音额度
	MidjourneyQuotas     []common.MidjourneyQuota    `json:"midjourney_quotas,omitempty"`      // Midjourney额度
	Remark               string                      `json:"remark,omitempty"`                 // 备注
}

type Model struct {
	Id                   string                      `json:"id,omitempty"`                      // ID
	Corp                 string                      `json:"corp,omitempty"`                    // 公司ID
	CorpName             string                      `json:"corp_name,omitempty"`               // 公司名称
	CorpCode             string                      `json:"corp_code,omitempty"`               // 公司代码
	Name                 string                      `json:"name,omitempty"`                    // 模型名称
	Model                string                      `json:"model,omitempty"`                   // 模型
	Type                 int                         `json:"type,omitempty"`                    // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
	BaseUrl              string                      `json:"base_url,omitempty"`                // 模型地址
	Path                 string                      `json:"path,omitempty"`                    // 模型路径
	IsEnablePresetConfig bool                        `json:"is_enable_preset_config,omitempty"` // 是否启用预设配置
	PresetConfig         common.PresetConfig         `json:"preset_config,omitempty"`           // 预设配置
	TextQuota            common.TextQuota            `json:"text_quota,omitempty"`              // 文本额度
	ImageQuotas          []common.ImageQuota         `json:"image_quotas,omitempty"`            // 图像额度
	AudioQuota           common.AudioQuota           `json:"audio_quota,omitempty"`             // 音频额度
	MultimodalQuota      common.MultimodalQuota      `json:"multimodal_quota,omitempty"`        // 多模态额度
	RealtimeQuota        common.RealtimeQuota        `json:"realtime_quota,omitempty"`          // 多模态实时额度
	MultimodalAudioQuota common.MultimodalAudioQuota `json:"multimodal_audio_quota,omitempty"`  // 多模态语音额度
	MidjourneyQuotas     []common.MidjourneyQuota    `json:"midjourney_quotas,omitempty"`       // Midjourney额度
	DataFormat           int                         `json:"data_format,omitempty"`             // 数据格式[1:统一格式, 2:官方格式]
	IsPublic             bool                        `json:"is_public"`                         // 是否公开
	IsEnableModelAgent   bool                        `json:"is_enable_model_agent"`             // 是否启用模型代理
	LbStrategy           int                         `json:"lb_strategy,omitempty"`             // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents          []string                    `json:"model_agents,omitempty"`            // 模型代理
	ModelAgentNames      []string                    `json:"model_agent_names,omitempty"`       // 模型代理名称
	IsEnableForward      bool                        `json:"is_enable_forward,omitempty"`       // 是否启用模型转发
	ForwardConfig        *common.ForwardConfig       `json:"forward_config,omitempty"`          // 模型转发配置
	IsEnableFallback     bool                        `json:"is_enable_fallback,omitempty"`      // 是否启用后备
	FallbackConfig       *common.FallbackConfig      `json:"fallback_config,omitempty"`         // 后备配置
	Remark               string                      `json:"remark,omitempty"`                  // 备注
	Status               int                         `json:"status,omitempty"`                  // 状态[1:正常, 2:禁用, -1:删除]
	Creator              string                      `json:"creator,omitempty"`                 // 创建人
	Updater              string                      `json:"updater,omitempty"`                 // 更新人
	CreatedAt            string                      `json:"created_at,omitempty"`              // 创建时间
	UpdatedAt            string                      `json:"updated_at,omitempty"`              // 更新时间
}
