package model

// 新建提供商接口请求参数
type ProviderCreateReq struct {
	Name     string `json:"name,omitempty"`         // 名称
	Code     string `json:"code,omitempty"`         // 代码
	Sort     int    `json:"sort,omitempty"`         // 排序
	IsPublic bool   `json:"is_public,omitempty"`    // 是否公开
	Remark   string `json:"remark,omitempty"`       // 备注
	Status   int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新提供商接口请求参数
type ProviderUpdateReq struct {
	Id       string `json:"id" v:"required"`        // ID
	Name     string `json:"name,omitempty"`         // 名称
	Code     string `json:"code,omitempty"`         // 代码
	Sort     int    `json:"sort,omitempty"`         // 排序
	IsPublic bool   `json:"is_public,omitempty"`    // 是否公开
	Remark   string `json:"remark,omitempty"`       // 备注
	Status   int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改提供商状态接口请求参数
type ProviderChangeStatusReq struct {
	Id     string `json:"id" v:"required"`        // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改提供商公开状态接口请求参数
type ProviderChangePublicReq struct {
	Id       string `json:"id" v:"required"`     // ID
	IsPublic bool   `json:"is_public,omitempty"` // 是否公开
}

// 提供商详情接口响应参数
type ProviderDetailRes struct {
	*Provider
}

// 提供商分页列表接口请求参数
type ProviderPageReq struct {
	Paging
	Name      string   `json:"name,omitempty"`       // 名称
	Code      string   `json:"code,omitempty"`       // 代码
	Sort      int      `json:"sort,omitempty"`       // 排序
	IsPublic  string   `json:"is_public,omitempty"`  // 是否公开
	Remark    string   `json:"remark,omitempty"`     // 备注
	Status    int      `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
	UpdatedAt []string `json:"updated_at,omitempty"` // 更新时间
}

// 提供商分页列表接口响应参数
type ProviderPageRes struct {
	Items  []*Provider `json:"items"`
	Paging *Paging     `json:"paging"`
}

// 提供商列表接口请求参数
type ProviderListReq struct {
	Name     string `json:"name,omitempty"`         // 名称
	Code     string `json:"code,omitempty"`         // 代码
	Sort     int    `json:"sort,omitempty"`         // 排序
	IsPublic bool   `json:"is_public,omitempty"`    // 是否公开
	Remark   string `json:"remark,omitempty"`       // 备注
	Status   int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 提供商列表接口响应参数
type ProviderListRes struct {
	Items []*Provider `json:"items"`
}

// 提供商批量操作接口请求参数
type ProviderBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

type Provider struct {
	Id        string `json:"id,omitempty"`         // ID
	Name      string `json:"name,omitempty"`       // 名称
	Code      string `json:"code,omitempty"`       // 代码
	Sort      int    `json:"sort,omitempty"`       // 排序
	IsPublic  bool   `json:"is_public,omitempty"`  // 是否公开
	Remark    string `json:"remark,omitempty"`     // 备注
	Status    int    `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Creator   string `json:"creator,omitempty"`    // 创建人
	Updater   string `json:"updater,omitempty"`    // 更新人
	CreatedAt string `json:"created_at,omitempty"` // 创建时间
	UpdatedAt string `json:"updated_at,omitempty"` // 更新时间
}
