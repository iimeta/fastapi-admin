package model

// 新建公司接口请求参数
type CorpCreateReq struct {
	Name   string `json:"name,omitempty"`         // 名称
	Code   string `json:"code,omitempty"`         // 代码
	Remark string `json:"remark,omitempty"`       // 备注
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新公司接口请求参数
type CorpUpdateReq struct {
	Id     string `json:"id" v:"required"`        // ID
	Name   string `json:"name,omitempty"`         // 名称
	Code   string `json:"code,omitempty"`         // 代码
	Remark string `json:"remark,omitempty"`       // 备注
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改公司状态接口请求参数
type CorpChangeStatusReq struct {
	Id     string `json:"id" v:"required"`        // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 公司详情接口响应参数
type CorpDetailRes struct {
	*Corp
}

// 公司分页列表接口请求参数
type CorpPageReq struct {
	Paging
	Name      string   `json:"name,omitempty"`       // 名称
	Code      string   `json:"code,omitempty"`       // 代码
	Remark    string   `json:"remark,omitempty"`     // 备注
	Status    int      `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
}

// 公司分页列表接口响应参数
type CorpPageRes struct {
	Items  []*Corp `json:"items"`
	Paging *Paging `json:"paging"`
}

// 公司列表接口请求参数
type CorpListReq struct {
	Name   string `json:"name,omitempty"`         // 名称
	Code   string `json:"code,omitempty"`         // 代码
	Remark string `json:"remark,omitempty"`       // 备注
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 公司列表接口响应参数
type CorpListRes struct {
	Items []*Corp `json:"items"`
}

// 公司批量操作接口请求参数
type CorpBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

type Corp struct {
	Id        string `json:"id,omitempty"`         // ID
	Name      string `json:"name,omitempty"`       // 名称
	Code      string `json:"code,omitempty"`       // 代码
	Remark    string `json:"remark,omitempty"`     // 备注
	Status    int    `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Creator   string `json:"creator,omitempty"`    // 创建人
	Updater   string `json:"updater,omitempty"`    // 更新人
	CreatedAt string `json:"created_at,omitempty"` // 创建时间
	UpdatedAt string `json:"updated_at,omitempty"` // 更新时间
}
