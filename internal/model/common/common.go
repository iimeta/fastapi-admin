package common

import (
	smodel "github.com/iimeta/fastapi-sdk/v2/model"
)

type PresetConfig struct {
	IsSupportSystemRole bool   `bson:"is_support_system_role" json:"is_support_system_role,omitempty"` // 是否支持system角色
	SystemRolePrompt    string `bson:"system_role_prompt"     json:"system_role_prompt,omitempty"`     // system角色预设提示词
	MinTokens           int    `bson:"min_tokens"             json:"min_tokens,omitempty"`             // max_tokens取值的最小值
	MaxTokens           int    `bson:"max_tokens"             json:"max_tokens,omitempty"`             // max_tokens取值的最大值
	IsSupportStream     bool   `bson:"is_support_stream"      json:"is_support_stream,omitempty"`      // 是否支持流式
}

type ForwardConfig struct {
	ForwardRule       int      `bson:"forward_rule,omitempty"   json:"forward_rule,omitempty"`        // 转发规则[1:全部转发, 2:按关键字, 3:内容长度, 4:已用额度]
	ForwardMode       int      `bson:"forward_mode,omitempty"   json:"forward_mode,omitempty"`        // 转发模式[1:固定, 2:轮询]
	MatchRule         []int    `bson:"match_rule,omitempty"     json:"match_rule,omitempty"`          // 转发规则为2时的匹配规则[1:智能匹配, 2:正则匹配]
	TargetModel       string   `bson:"target_model,omitempty"   json:"target_model,omitempty"`        // 转发规则为[1,3,4]时的目标模型
	DecisionModel     string   `bson:"decision_model,omitempty" json:"decision_model,omitempty"`      // 转发规则为2时并且匹配规则为1时的判定模型
	Keywords          []string `bson:"keywords,omitempty"       json:"keywords,omitempty"`            // 转发规则为2时的关键字
	TargetModels      []string `bson:"target_models,omitempty"  json:"target_models,omitempty"`       // 转发规则为[1,2]时的目标模型
	ContentLength     int      `bson:"content_length,omitempty" json:"content_length,omitempty"`      // 转发规则为3时的内容长度
	UsedQuota         float64  `bson:"used_quota,omitempty"     json:"used_quota,omitempty"`          // 转发规则为4时的已用额度
	TargetModelName   string   `bson:"-"                        json:"target_model_name,omitempty"`   // 转发规则为[1,3,4]时的目标模型名称
	DecisionModelName string   `bson:"-"                        json:"decision_model_name,omitempty"` // 转发规则为2时并且匹配规则为1时的判定模型名称
	TargetModelNames  []string `bson:"-"                        json:"target_model_names,omitempty"`  // 转发规则为[1,2]时的目标模型名称
}

type FallbackConfig struct {
	ModelAgent     string `bson:"model_agent,omitempty"      json:"model_agent,omitempty"`      // 后备模型代理
	ModelAgentName string `bson:"model_agent_name,omitempty" json:"model_agent_name,omitempty"` // 后备模型代理名称
	Model          string `bson:"model,omitempty"            json:"model,omitempty"`            // 后备模型
	ModelName      string `bson:"model_name,omitempty"       json:"model_name,omitempty"`       // 后备模型名称
}

type Message struct {
	Role         string               `bson:"role,omitempty"          json:"role,omitempty"`    // 角色
	Content      string               `bson:"content,omitempty"       json:"content,omitempty"` // 内容
	Refusal      string               `bson:"refusal,omitempty"       json:"refusal,omitempty"`
	Name         string               `bson:"name,omitempty"          json:"name,omitempty"`
	FunctionCall *smodel.FunctionCall `bson:"function_call,omitempty" json:"function_call,omitempty"`
	ToolCalls    []smodel.ToolCall    `bson:"tool_calls,omitempty"    json:"tool_calls,omitempty"`
	ToolCallId   string               `bson:"tool_call_id,omitempty"  json:"tool_call_id,omitempty"`
	Audio        *smodel.Audio        `bson:"audio,omitempty"         json:"audio,omitempty"`
}

type Retry struct {
	IsRetry    bool   `bson:"is_retry,omitempty"    json:"is_retry,omitempty"`    // 是否重试
	RetryCount int    `bson:"retry_count,omitempty" json:"retry_count,omitempty"` // 重试次数
	ErrMsg     string `bson:"err_msg,omitempty"     json:"err_msg,omitempty"`     // 错误信息
}

type ImageData struct {
	URL           string `bson:"url,omitempty"`
	B64JSON       string `bson:"b64_json,omitempty"`
	RevisedPrompt string `bson:"revised_prompt,omitempty"`
}

type ModelStat struct {
	ModelId        string  `bson:"model_id,omitempty"        json:"model_id,omitempty"`        // 模型ID
	Model          string  `bson:"model,omitempty"           json:"model,omitempty"`           // 模型
	Total          int     `bson:"total,omitempty"           json:"total,omitempty"`           // 总数
	Tokens         float64 `bson:"tokens,omitempty"          json:"tokens,omitempty"`          // 令牌数
	Abnormal       int     `bson:"abnormal,omitempty"        json:"abnormal,omitempty"`        // 异常数
	AbnormalTokens float64 `bson:"abnormal_tokens,omitempty" json:"abnormal_tokens,omitempty"` // 异常令牌数
}

type Carousel struct {
	ImageUrl string `bson:"image_url,omitempty" json:"image_url,omitempty"` // 图像URL
	JumpUrl  string `bson:"jump_url,omitempty"  json:"jump_url,omitempty"`  // 跳转URL
}

type Announcement struct {
	Content string `bson:"content,omitempty"  json:"content,omitempty"`  // 公告内容
	JumpUrl string `bson:"jump_url,omitempty" json:"jump_url,omitempty"` // 跳转URL
}

type Document struct {
	Title   string `bson:"title,omitempty"    json:"title,omitempty"`    // 文档标题
	JumpUrl string `bson:"jump_url,omitempty" json:"jump_url,omitempty"` // 跳转URL
}
