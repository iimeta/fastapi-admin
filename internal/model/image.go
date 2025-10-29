package model

import "github.com/iimeta/fastapi-admin/internal/model/common"

// 绘图详情接口响应参数
type ImageDetailRes struct {
	*Image
}

// 绘图分页列表接口请求参数
type ImagePageReq struct {
	Paging
	TraceId     string   `json:"trace_id,omitempty"`     // 日志ID
	UserId      int      `json:"user_id,omitempty"`      // 用户ID
	AppId       int      `json:"app_id,omitempty"`       // 应用ID
	Key         string   `json:"key,omitempty"`          // 密钥
	Models      []string `json:"models,omitempty"`       // 模型
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
	TotalTime   int64    `json:"total_time,omitempty"`   // 总时间
	Status      int      `json:"status,omitempty"`       // 状态[1:成功, -1:失败]
	ReqTime     []string `json:"req_time,omitempty"`     // 请求时间
}

// 绘图分页列表接口响应参数
type ImagePageRes struct {
	Items  []*Image `json:"items"`
	Paging *Paging  `json:"paging"`
}

// 绘图日志详情复制字段值接口请求参数
type ImageCopyFieldReq struct {
	Id    string `json:"id"`
	Field string `json:"field"`
}

// 绘图日志详情复制字段值接口响应参数
type ImageCopyFieldRes struct {
	Value string `json:"value"`
}

type Image struct {
	Id                   string                 `json:"id,omitempty"`                      // ID
	TraceId              string                 `json:"trace_id,omitempty"`                // 日志ID
	UserId               int                    `json:"user_id,omitempty"`                 // 用户ID
	AppId                int                    `json:"app_id,omitempty"`                  // 应用ID
	ProviderId           string                 `json:"provider_id,omitempty"`             // 提供商ID
	ProviderName         string                 `json:"provider_name,omitempty"`           // 提供商名称
	ModelId              string                 `json:"model_id,omitempty"`                // 模型ID
	ModelName            string                 `json:"model_name,omitempty"`              // 模型名称
	Model                string                 `json:"model,omitempty"`                   // 模型
	ModelType            int                    `json:"model_type,omitempty"`              // 模型类型
	Key                  string                 `json:"key,omitempty"`                     // 密钥
	IsEnablePresetConfig bool                   `json:"is_enable_preset_config,omitempty"` // 是否启用预设配置
	PresetConfig         common.PresetConfig    `json:"preset_config,omitempty"`           // 预设配置
	IsEnableModelAgent   bool                   `json:"is_enable_model_agent,omitempty"`   // 是否启用模型代理
	ModelAgentId         string                 `json:"model_agent_id,omitempty"`          // 模型代理ID
	ModelAgent           *ModelAgent            `json:"model_agent,omitempty"`             // 模型代理信息
	IsEnableForward      bool                   `json:"is_enable_forward,omitempty"`       // 是否启用模型转发
	ForwardConfig        *common.ForwardConfig  `json:"forward_config,omitempty"`          // 模型转发配置
	IsSmartMatch         bool                   `json:"is_smart_match,omitempty"`          // 是否智能匹配
	IsEnableFallback     bool                   `json:"is_enable_fallback,omitempty"`      // 是否启用后备
	FallbackConfig       *common.FallbackConfig `json:"fallback_config,omitempty"`         // 后备配置
	RealModelId          string                 `json:"real_model_id,omitempty"`           // 真实模型ID
	RealModelName        string                 `json:"real_model_name,omitempty"`         // 真实模型名称
	RealModel            string                 `json:"real_model,omitempty"`              // 真实模型
	Prompt               string                 `json:"prompt,omitempty"`                  // 提示(提问)
	Size                 string                 `json:"size,omitempty"`                    // 尺寸大小
	N                    int                    `json:"n,omitempty"`                       // 图像数
	Quality              string                 `json:"quality,omitempty"`                 // 图像质量[hd]
	Style                string                 `json:"style,omitempty"`                   // 图像样式[vivid, natural]
	ResponseFormat       string                 `json:"response_format,omitempty"`         // 图像格式[url, b64_json]
	Images               []string               `json:"images,omitempty"`                  // 生成图像url
	ImageData            []common.ImageData     `json:"image_data,omitempty"`              // 生成图像数据
	Spend                common.Spend           `json:"spend,omitempty"`                   // 花费
	TotalTime            int64                  `json:"total_time,omitempty"`              // 总时间
	InternalTime         int64                  `json:"internal_time,omitempty"`           // 内耗时间
	ReqTime              string                 `json:"req_time,omitempty"`                // 请求时间
	ReqDate              string                 `json:"req_date,omitempty"`                // 请求日期
	ClientIp             string                 `json:"client_ip,omitempty"`               // 客户端IP
	RemoteIp             string                 `json:"remote_ip,omitempty"`               // 远程IP
	LocalIp              string                 `json:"local_ip,omitempty"`                // 本地IP
	ErrMsg               string                 `json:"err_msg,omitempty"`                 // 错误信息
	IsRetry              bool                   `json:"is_retry,omitempty"`                // 是否重试
	Retry                *common.Retry          `json:"retry,omitempty"`                   // 重试
	Status               int                    `json:"status,omitempty"`                  // 状态[1:成功, -1:失败, 2:中止, 3:重试]
	Host                 string                 `json:"host,omitempty"`                    // Host
	Creator              string                 `json:"creator,omitempty"`                 // 创建人
	Updater              string                 `json:"updater,omitempty"`                 // 更新人
	CreatedAt            string                 `json:"created_at,omitempty"`              // 创建时间
	UpdatedAt            string                 `json:"updated_at,omitempty"`              // 更新时间
}
