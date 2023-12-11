package model

// 新建模型接口请求参数
type ModelCreateReq struct {
	Model
}

// 更新模型接口请求参数
type ModelUpdateReq struct {
}

// 模型详情接口响应参数
type ModelDetailRes struct {
	Model
}

type Model struct {
	Id        string   `json:"_id,omitempty"`        // ID
	Corp      string   `json:"corp,omitempty"`       // 公司
	Name      string   `json:"name,omitempty"`       // 模型名称
	Type      int      `json:"type,omitempty"`       // 模型类型[1:文生文; 2:文生图; 3:图生文; 4:图生图; 5:文生语音; 6:语音生文]
	BaseUrl   string   `json:"base_url,omitempty"`   // 默认官方模型地址
	Path      string   `json:"path,omitempty"`       // 默认官方模型地址路径
	Proxy     string   `json:"proxy,omitempty"`      // 代理
	Keys      []string `json:"keys,omitempty"`       // 密钥
	Status    int      `json:"status,omitempty"`     // 状态[1:正常; 2:禁用; -1:删除]
	Creator   string   `json:"creator,omitempty"`    // 创建人
	Updater   string   `json:"updater,omitempty"`    // 更新人
	CreatedAt int64    `json:"created_at,omitempty"` // 创建时间
	UpdatedAt int64    `json:"updated_at,omitempty"` // 更新时间
}
