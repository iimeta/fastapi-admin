package model

// 新建角色接口请求参数
type SysRoleCreateReq struct {
	Pid    string   `json:"pid,omitempty"`    // 父ID
	Name   string   `json:"name,omitempty"`   // 名称
	Type   int      `json:"type,omitempty"`   // 类型
	Perms  []string `json:"perms,omitempty"`  // 权限
	Remark string   `json:"remark,omitempty"` // 备注
	Status int      `json:"status,omitempty"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新角色接口请求参数
type SysRoleUpdateReq struct {
	Id     string   `json:"id,omitempty"`     // ID
	Pid    string   `json:"pid,omitempty"`    // 父ID
	Name   string   `json:"name,omitempty"`   // 名称
	Type   int      `json:"type,omitempty"`   // 类型
	Perms  []string `json:"perms,omitempty"`  // 权限
	Remark string   `json:"remark,omitempty"` // 备注
	Status int      `json:"status,omitempty"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 角色详情接口响应参数
type SysRoleDetailRes struct {
	*SysRole
}

// 角色分页列表接口请求参数
type SysRolePageReq struct {
	Paging
	Pid    string   `json:"pid,omitempty"`    // 父ID
	Name   string   `json:"name,omitempty"`   // 名称
	Type   int      `json:"type,omitempty"`   // 类型
	Perms  []string `json:"perms,omitempty"`  // 权限
	Remark string   `json:"remark,omitempty"` // 备注
	Status int      `json:"status,omitempty"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 角色分页列表接口响应参数
type SysRolePageRes struct {
	Items  []*SysRole `json:"items"`
	Paging *Paging    `json:"paging"`
}

type SysRole struct {
	Id        string   `json:"id,omitempty"`         // ID
	Pid       string   `json:"pid,omitempty"`        // 父ID
	Name      string   `json:"name,omitempty"`       // 名称
	Type      int      `json:"type,omitempty"`       // 类型
	Perms     []string `json:"perms,omitempty"`      // 权限
	Remark    string   `json:"remark,omitempty"`     // 备注
	Status    int      `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Creator   string   `json:"creator,omitempty"`    // 创建人
	Updater   string   `json:"updater,omitempty"`    // 更新人
	CreatedAt int64    `json:"created_at,omitempty"` // 创建时间
	UpdatedAt int64    `json:"updated_at,omitempty"` // 更新时间
}
