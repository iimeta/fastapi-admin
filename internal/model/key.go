package model

// 新建密钥接口请求参数
type KeyCreateReq struct {
	Corp   string   `json:"corp,omitempty"`         // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key    string   `json:"key,omitempty"`          // 密钥
	Models []string `json:"models,omitempty"`       // 模型
	Remark string   `json:"remark,omitempty"`       // 备注
	Status int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新密钥接口请求参数
type KeyUpdateReq struct {
	Id     string   `json:"id,omitempty"`           // ID
	Corp   string   `json:"corp,omitempty"`         // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key    string   `json:"key,omitempty"`          // 密钥
	Models []string `json:"models,omitempty"`       // 模型
	Remark string   `json:"remark,omitempty"`       // 备注
	Status int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 密钥详情接口响应参数
type KeyDetailRes struct {
	*Key
}

// 密钥分页列表接口请求参数
type KeyPageReq struct {
	Paging
	Corp      string   `json:"corp,omitempty"`       // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key       string   `json:"key,omitempty"`        // 密钥
	Models    []string `json:"models,omitempty"`     // 模型
	Remark    string   `json:"remark,omitempty"`     // 备注
	Status    int      `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
}

// 密钥分页列表接口响应参数
type KeyPageRes struct {
	Items  []*Key  `json:"items"`
	Paging *Paging `json:"paging"`
}

// 密钥列表接口请求参数
type KeyListReq struct {
	Corp   string   `json:"corp,omitempty"`   // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key    string   `json:"key,omitempty"`    // 密钥
	Models []string `json:"models,omitempty"` // 模型
	Quota  int      `json:"quota,omitempty"`  // 额度
	Remark string   `json:"remark,omitempty"` // 备注
	Status int      `json:"status,omitempty"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 密钥列表接口响应参数
type KeyListRes struct {
	Items []*Key `json:"items"`
}

type Key struct {
	Id        string   `json:"id,omitempty"`         // ID
	Corp      string   `json:"corp,omitempty"`       // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key       string   `json:"key,omitempty"`        // 密钥
	Models    []string `json:"models,omitempty"`     // 模型
	Quota     int      `json:"quota,omitempty"`      // 额度
	Remark    string   `json:"remark,omitempty"`     // 备注
	Status    int      `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Creator   string   `json:"creator,omitempty"`    // 创建人
	Updater   string   `json:"updater,omitempty"`    // 更新人
	CreatedAt int64    `json:"created_at,omitempty"` // 创建时间
	UpdatedAt int64    `json:"updated_at,omitempty"` // 更新时间
}
