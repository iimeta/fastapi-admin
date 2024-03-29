package model

// 聊天详情接口响应参数
type ChatDetailRes struct {
	*Chat
}

// 聊天分页列表接口请求参数
type ChatPageReq struct {
	Paging
	UserId    int      `json:"user_id,omitempty"`    // 用户ID
	AppId     int      `json:"app_id,omitempty"`     // 应用ID
	Key       string   `json:"key,omitempty"`        // 密钥
	Models    []string `json:"models,omitempty"`     // 模型
	TotalTime int64    `json:"total_time,omitempty"` // 总时间
	Status    int      `json:"status,omitempty"`     // 状态[1:成功, -1:失败]
	ReqTimes  []string `json:"req_times,omitempty"`  // 请求时间
}

// 聊天分页列表接口响应参数
type ChatPageRes struct {
	Items  []*Chat `json:"items"`
	Paging *Paging `json:"paging"`
}

type Chat struct {
	Id                 string      `json:"id,omitempty"`                    // ID
	TraceId            string      `json:"trace_id,omitempty"`              // 日志ID
	UserId             int         `json:"user_id,omitempty"`               // 用户ID
	AppId              int         `json:"app_id,omitempty"`                // 应用ID
	Corp               string      `json:"corp,omitempty"`                  // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	ModelId            string      `json:"model_id,omitempty"`              // 模型ID
	Name               string      `json:"name,omitempty"`                  // 模型名称
	Model              string      `json:"model,omitempty"`                 // 模型
	Type               int         `json:"type,omitempty"`                  // 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文]
	Key                string      `json:"key,omitempty"`                   // 密钥
	IsEnableModelAgent bool        `json:"is_enable_model_agent,omitempty"` // 是否启用模型代理
	ModelAgentId       string      `json:"model_agent_id,omitempty"`        // 模型代理ID
	ModelAgent         *ModelAgent `json:"model_agent,omitempty"`           // 模型代理信息
	Proxy              string      `json:"proxy,omitempty"`                 // 代理
	Stream             bool        `json:"stream"`                          // 是否流式
	Messages           []Message   `json:"messages,omitempty"`              // 完整提示(提问)
	Prompt             string      `json:"prompt,omitempty"`                // 提示(提问)
	Completion         string      `json:"completion,omitempty"`            // 补全(回答)
	PromptRatio        float64     `json:"prompt_ratio,omitempty"`          // 提示倍率(提问倍率)
	CompletionRatio    float64     `json:"completion_ratio,omitempty"`      // 补全倍率(回答倍率)
	PromptTokens       int         `json:"prompt_tokens,omitempty"`         // 提示令牌数(提问令牌数)
	CompletionTokens   int         `json:"completion_tokens,omitempty"`     // 补全令牌数(回答令牌数)
	TotalTokens        int         `json:"total_tokens,omitempty"`          // 总令牌数
	ConnTime           int64       `json:"conn_time,omitempty"`             // 连接时间
	Duration           int64       `json:"duration,omitempty"`              // 持续时间
	TotalTime          int64       `json:"total_time,omitempty"`            // 总时间
	InternalTime       int64       `json:"internal_time,omitempty"`         // 内耗时间
	ReqTime            string      `json:"req_time,omitempty"`              // 请求时间
	ReqDate            string      `json:"req_date,omitempty"`              // 请求日期
	ClientIp           string      `json:"client_ip,omitempty"`             // 客户端IP
	RemoteIp           string      `json:"remote_ip,omitempty"`             // 远程IP
	ErrMsg             string      `json:"err_msg,omitempty"`               // 错误信息
	Status             int         `json:"status,omitempty"`                // 状态[1:成功, -1:失败]
	Creator            string      `json:"creator,omitempty"`               // 创建人
	Updater            string      `json:"updater,omitempty"`               // 更新人
	CreatedAt          string      `json:"created_at,omitempty"`            // 创建时间
	UpdatedAt          string      `json:"updated_at,omitempty"`            // 更新时间
}

type Message struct {
	Role    string `json:"role,omitempty"`    // 角色
	Content string `json:"content,omitempty"` // 内容
}
