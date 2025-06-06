package do

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

const (
	MODEL_COLLECTION = "model"
)

type Model struct {
	gmeta.Meta           `collection:"model" bson:"-"`
	Corp                 string                      `bson:"corp,omitempty"`            // 公司
	Name                 string                      `bson:"name,omitempty"`            // 模型名称
	Model                string                      `bson:"model,omitempty"`           // 模型
	Type                 int                         `bson:"type,omitempty"`            // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
	BaseUrl              string                      `bson:"base_url"`                  // 模型地址
	Path                 string                      `bson:"path"`                      // 模型路径
	IsEnablePresetConfig bool                        `bson:"is_enable_preset_config"`   // 是否启用预设配置
	PresetConfig         common.PresetConfig         `bson:"preset_config"`             // 预设配置
	TextQuota            common.TextQuota            `bson:"text_quota"`                // 文本额度
	ImageQuota           common.ImageQuota           `bson:"image_quota"`               // 图像额度
	AudioQuota           common.AudioQuota           `bson:"audio_quota"`               // 音频额度
	MultimodalQuota      common.MultimodalQuota      `bson:"multimodal_quota"`          // 多模态额度
	RealtimeQuota        common.RealtimeQuota        `bson:"realtime_quota"`            // 多模态实时额度
	MultimodalAudioQuota common.MultimodalAudioQuota `bson:"multimodal_audio_quota"`    // 多模态语音额度
	MidjourneyQuotas     []common.MidjourneyQuota    `bson:"midjourney_quotas"`         // Midjourney额度
	DataFormat           int                         `bson:"data_format,omitempty"`     // 数据格式[1:统一格式, 2:官方格式]
	IsPublic             bool                        `bson:"is_public"`                 // 是否公开
	IsEnableModelAgent   bool                        `bson:"is_enable_model_agent"`     // 是否启用模型代理
	LbStrategy           int                         `bson:"lb_strategy,omitempty"`     // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents          []string                    `bson:"model_agents"`              // 模型代理
	IsEnableForward      bool                        `bson:"is_enable_forward"`         // 是否启用模型转发
	ForwardConfig        *common.ForwardConfig       `bson:"forward_config,omitempty"`  // 模型转发配置
	IsEnableFallback     bool                        `bson:"is_enable_fallback"`        // 是否启用后备
	FallbackConfig       *common.FallbackConfig      `bson:"fallback_config,omitempty"` // 后备配置
	Remark               string                      `bson:"remark"`                    // 备注
	Status               int                         `bson:"status,omitempty"`          // 状态[1:正常, 2:禁用, -1:删除]
	Creator              string                      `bson:"creator,omitempty"`         // 创建人
	Updater              string                      `bson:"updater,omitempty"`         // 更新人
	CreatedAt            int64                       `bson:"created_at,omitempty"`      // 创建时间
	UpdatedAt            int64                       `bson:"updated_at,omitempty"`      // 更新时间
}
