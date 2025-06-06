package do

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

const (
	IMAGE_COLLECTION = "image"
)

type Image struct {
	gmeta.Meta           `collection:"image" bson:"-"`
	Id                   string                 `bson:"_id,omitempty"`                     // ID
	TraceId              string                 `bson:"trace_id,omitempty"`                // 日志ID
	UserId               int                    `bson:"user_id,omitempty"`                 // 用户ID
	AppId                int                    `bson:"app_id,omitempty"`                  // 应用ID
	Corp                 string                 `bson:"corp,omitempty"`                    // 公司
	GroupId              string                 `bson:"group_id,omitempty"`                // 分组ID
	GroupName            string                 `bson:"group_name,omitempty"`              // 分组名称
	Discount             float64                `bson:"discount,omitempty"`                // 分组折扣
	ModelId              string                 `bson:"model_id,omitempty"`                // 模型ID
	Name                 string                 `bson:"name,omitempty"`                    // 模型名称
	Model                string                 `bson:"model,omitempty"`                   // 模型
	Type                 int                    `bson:"type,omitempty"`                    // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
	Key                  string                 `bson:"key,omitempty"`                     // 密钥
	IsEnablePresetConfig bool                   `bson:"is_enable_preset_config,omitempty"` // 是否启用预设配置
	PresetConfig         common.PresetConfig    `bson:"preset_config,omitempty"`           // 预设配置
	IsEnableModelAgent   bool                   `bson:"is_enable_model_agent,omitempty"`   // 是否启用模型代理
	ModelAgentId         string                 `bson:"model_agent_id,omitempty"`          // 模型代理ID
	ModelAgent           *ModelAgent            `bson:"model_agent,omitempty"`             // 模型代理信息
	IsEnableForward      bool                   `bson:"is_enable_forward,omitempty"`       // 是否启用模型转发
	ForwardConfig        *common.ForwardConfig  `bson:"forward_config,omitempty"`          // 模型转发配置
	IsSmartMatch         bool                   `bson:"is_smart_match,omitempty"`          // 是否智能匹配
	IsEnableFallback     bool                   `bson:"is_enable_fallback,omitempty"`      // 是否启用后备
	FallbackConfig       *common.FallbackConfig `bson:"fallback_config,omitempty"`         // 后备配置
	RealModelId          string                 `bson:"real_model_id,omitempty"`           // 真实模型ID
	RealModelName        string                 `bson:"real_model_name,omitempty"`         // 真实模型名称
	RealModel            string                 `bson:"real_model,omitempty"`              // 真实模型
	Prompt               string                 `bson:"prompt,omitempty"`                  // 提示(提问)
	Size                 string                 `bson:"size,omitempty"`                    // 尺寸大小
	N                    int                    `bson:"n,omitempty"`                       // 图像数
	Quality              string                 `bson:"quality,omitempty"`                 // 图像质量[high, medium, low, hd, standard]
	Style                string                 `bson:"style,omitempty"`                   // 图像样式[vivid, natural]
	ResponseFormat       string                 `bson:"response_format,omitempty"`         // 图像格式[url, b64_json]
	ImageData            []common.ImageData     `bson:"image_data,omitempty"`              // 生成图像数据
	ImageQuota           common.ImageQuota      `bson:"image_quota,omitempty"`             // 图像额度
	GenerationQuota      common.GenerationQuota `bson:"generation_quota,omitempty"`        // 生成额度
	InputTokens          int                    `bson:"input_tokens,omitempty"`            // 输入令牌数
	OutputTokens         int                    `bson:"output_tokens,omitempty"`           // 输出令牌数
	TextTokens           int                    `bson:"text_tokens,omitempty"`             // 文本令牌数
	ImageTokens          int                    `bson:"image_tokens,omitempty"`            // 图像令牌数
	TotalTokens          int                    `bson:"total_tokens,omitempty"`            // 总令牌数
	TotalTime            int64                  `bson:"total_time,omitempty"`              // 总时间
	InternalTime         int64                  `bson:"internal_time,omitempty"`           // 内耗时间
	ReqTime              int64                  `bson:"req_time,omitempty"`                // 请求时间
	ReqDate              string                 `bson:"req_date,omitempty"`                // 请求日期
	ClientIp             string                 `bson:"client_ip,omitempty"`               // 客户端IP
	RemoteIp             string                 `bson:"remote_ip,omitempty"`               // 远程IP
	LocalIp              string                 `bson:"local_ip,omitempty"`                // 本地IP
	ErrMsg               string                 `bson:"err_msg,omitempty"`                 // 错误信息
	IsRetry              bool                   `bson:"is_retry,omitempty"`                // 是否重试
	Retry                *common.Retry          `bson:"retry,omitempty"`                   // 重试
	Status               int                    `bson:"status,omitempty"`                  // 状态[1:成功, -1:失败, 2:中止, 3:重试]
	Host                 string                 `bson:"host,omitempty"`                    // Host
	Rid                  int                    `bson:"rid,omitempty"`                     // 代理商ID
	Creator              string                 `bson:"creator,omitempty"`                 // 创建人
	Updater              string                 `bson:"updater,omitempty"`                 // 更新人
	CreatedAt            int64                  `bson:"created_at,omitempty"`              // 创建时间
	UpdatedAt            int64                  `bson:"updated_at,omitempty"`              // 更新时间
}
