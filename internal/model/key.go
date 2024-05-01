package model

// 新建密钥接口请求参数
type KeyCreateReq struct {
	Corp         string   `json:"corp,omitempty"`           // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key          string   `json:"key,omitempty"`            // 密钥
	Models       []string `json:"models,omitempty"`         // 模型
	ModelAgents  []string `json:"model_agents,omitempty"`   // 模型代理
	IsAgentsOnly bool     `json:"is_agents_only,omitempty"` // 是否代理专用
	Remark       string   `json:"remark,omitempty"`         // 备注
	Status       int      `json:"status,omitempty" d:"1"`   // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新密钥接口请求参数
type KeyUpdateReq struct {
	Id           string   `json:"id,omitempty"`                  // ID
	Corp         string   `json:"corp,omitempty"`                // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key          string   `json:"key,omitempty"`                 // 密钥
	Models       []string `json:"models,omitempty" d:"[]"`       // 模型
	ModelAgents  []string `json:"model_agents,omitempty" d:"[]"` // 模型代理
	IsAgentsOnly bool     `json:"is_agents_only,omitempty"`      // 是否代理专用
	Remark       string   `json:"remark,omitempty"`              // 备注
	Status       int      `json:"status,omitempty" d:"1"`        // 状态[1:正常, 2:禁用, -1:删除]
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
	Type        int      `json:"type,omitempty"`         // 密钥类型[1:应用, 2:模型]
	AppId       int      `json:"app_id,omitempty"`       // 应用ID
	Corp        string   `json:"corp,omitempty"`         // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
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
	Type        int      `json:"type,omitempty"`         // 密钥类型[1:应用, 2:模型]
	AppId       int      `json:"app_id,omitempty"`       // 应用ID
	Corp        string   `json:"corp,omitempty"`         // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key         string   `json:"key,omitempty"`          // 密钥
	Models      []string `json:"models,omitempty"`       // 模型
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
	Quota       int      `json:"quota,omitempty"`        // 额度
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
}

// 密钥列表接口响应参数
type KeyListRes struct {
	Items []*Key `json:"items"`
}

// 密钥批量操作接口请求参数
type KeyBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

type Key struct {
	Id              string   `json:"id,omitempty"`                // ID
	AppId           int      `json:"app_id,omitempty"`            // 应用ID
	Corp            string   `json:"corp,omitempty"`              // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key             string   `json:"key,omitempty"`               // 密钥
	Type            int      `json:"type,omitempty"`              // 密钥类型[1:应用, 2:模型]
	Models          []string `json:"models,omitempty"`            // 模型
	ModelNames      []string `json:"model_names,omitempty"`       // 模型名称
	ModelAgents     []string `json:"model_agents,omitempty"`      // 模型代理
	ModelAgentNames []string `json:"model_agent_names,omitempty"` // 模型代理名称
	IsAgentsOnly    bool     `json:"is_agents_only"`              // 是否代理专用
	IsLimitQuota    bool     `json:"is_limit_quota"`              // 是否限制额度
	Quota           int      `json:"quota,omitempty"`             // 剩余额度
	UsedQuota       int      `json:"used_quota,omitempty"`        // 已用额度
	IpWhitelist     []string `json:"ip_whitelist,omitempty"`      // IP白名单
	IpBlacklist     []string `json:"ip_blacklist,omitempty"`      // IP黑名单
	Remark          string   `json:"remark,omitempty"`            // 备注
	Status          int      `json:"status,omitempty"`            // 状态[1:正常, 2:禁用, -1:删除]
	Creator         string   `json:"creator,omitempty"`           // 创建人
	Updater         string   `json:"updater,omitempty"`           // 更新人
	CreatedAt       string   `json:"created_at,omitempty"`        // 创建时间
	UpdatedAt       string   `json:"updated_at,omitempty"`        // 更新时间
}
