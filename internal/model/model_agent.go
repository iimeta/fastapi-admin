package model

// 新建模型代理接口请求参数
type ModelAgentCreateReq struct {
	Name    string   `json:"name,omitempty"`         // 模型代理名称
	BaseUrl string   `json:"base_url,omitempty"`     // 模型代理地址
	Path    string   `json:"path,omitempty"`         // 模型代理地址路径
	Weight  int      `json:"weight,omitempty"`       // 权重
	Models  []string `json:"models,omitempty"`       // 绑定模型
	Key     string   `json:"key,omitempty"`          // 密钥
	Remark  string   `json:"remark,omitempty"`       // 备注
	Status  int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新模型代理接口请求参数
type ModelAgentUpdateReq struct {
	Id      string   `json:"id" v:"required"`        // ID
	Name    string   `json:"name,omitempty"`         // 模型代理名称
	BaseUrl string   `json:"base_url,omitempty"`     // 模型代理地址
	Path    string   `json:"path,omitempty"`         // 模型代理地址路径
	Weight  int      `json:"weight,omitempty"`       // 权重
	Models  []string `json:"models,omitempty"`       // 绑定模型
	Key     string   `json:"key,omitempty"`          // 密钥
	Remark  string   `json:"remark,omitempty"`       // 备注
	Status  int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改模型代理状态接口请求参数
type ModelAgentChangeStatusReq struct {
	Id     string `json:"id" v:"required"`        // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型代理详情接口响应参数
type ModelAgentDetailRes struct {
	*ModelAgent
}

// 模型代理分页列表接口请求参数
type ModelAgentPageReq struct {
	Paging
	Name      string   `json:"name,omitempty"`         // 模型代理名称
	BaseUrl   string   `json:"base_url,omitempty"`     // 模型代理地址
	Path      string   `json:"path,omitempty"`         // 模型代理地址路径
	Weight    int      `json:"weight,omitempty"`       // 权重
	Models    []string `json:"models,omitempty"`       // 绑定模型
	Remark    string   `json:"remark,omitempty"`       // 备注
	Status    int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt []string `json:"created_at,omitempty"`   // 创建时间
}

// 模型代理分页列表接口响应参数
type ModelAgentPageRes struct {
	Items  []*ModelAgent `json:"items"`
	Paging *Paging       `json:"paging"`
}

// 模型代理列表接口请求参数
type ModelAgentListReq struct {
	Name    string `json:"name,omitempty"`         // 模型代理名称
	BaseUrl string `json:"base_url,omitempty"`     // 模型代理地址
	Path    string `json:"path,omitempty"`         // 模型代理地址路径
	Weight  int    `json:"weight,omitempty"`       // 权重
	Remark  string `json:"remark,omitempty"`       // 备注
	Status  int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 模型代理列表接口响应参数
type ModelAgentListRes struct {
	Items []*ModelAgent `json:"items"`
}

type ModelAgent struct {
	Id         string   `json:"id,omitempty"`          // ID
	Name       string   `json:"name,omitempty"`        // 模型代理名称
	BaseUrl    string   `json:"base_url,omitempty"`    // 模型代理地址
	Path       string   `json:"path,omitempty"`        // 模型代理地址路径
	Weight     int      `json:"weight,omitempty"`      // 权重
	Models     []string `json:"models,omitempty"`      // 绑定模型
	ModelNames []string `json:"model_names,omitempty"` // 模型名称
	Key        string   `json:"key,omitempty"`         // 密钥
	Remark     string   `json:"remark,omitempty"`      // 备注
	Status     int      `json:"status,omitempty"`      // 状态[1:正常, 2:禁用, -1:删除]
	Creator    string   `json:"creator,omitempty"`     // 创建人
	Updater    string   `json:"updater,omitempty"`     // 更新人
	CreatedAt  string   `json:"created_at,omitempty"`  // 创建时间
	UpdatedAt  string   `json:"updated_at,omitempty"`  // 更新时间
}
