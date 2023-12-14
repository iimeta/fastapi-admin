package model

// 新建菜单接口请求参数
type SysMenuCreateReq struct {
	Pid    string `json:"pid,omitempty"`    // 父ID
	Name   string `json:"name,omitempty"`   // 名称
	Perm   string `json:"perm,omitempty"`   // 权限
	Type   int    `json:"type,omitempty"`   // 类型
	Route  string `json:"route,omitempty"`  // 路由/API
	Sort   int    `json:"sort,omitempty"`   // 排序
	Level  int    `json:"level,omitempty"`  // 层级
	Hidden bool   `json:"hidden,omitempty"` // 隐藏
	Remark string `json:"remark,omitempty"` // 备注
	Status int    `json:"status,omitempty"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新菜单接口请求参数
type SysMenuUpdateReq struct {
	Id     string `json:"id,omitempty"`     // ID
	Pid    string `json:"pid,omitempty"`    // 父ID
	Name   string `json:"name,omitempty"`   // 名称
	Perm   string `json:"perm,omitempty"`   // 权限
	Type   int    `json:"type,omitempty"`   // 类型
	Route  string `json:"route,omitempty"`  // 路由/API
	Sort   int    `json:"sort,omitempty"`   // 排序
	Level  int    `json:"level,omitempty"`  // 层级
	Hidden bool   `json:"hidden,omitempty"` // 隐藏
	Remark string `json:"remark,omitempty"` // 备注
	Status int    `json:"status,omitempty"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 菜单详情接口响应参数
type SysMenuDetailRes struct {
	*SysMenu
}

// 菜单分页列表接口请求参数
type SysMenuPageReq struct {
	Paging
	Pid    string `json:"pid,omitempty"`    // 父ID
	Name   string `json:"name,omitempty"`   // 名称
	Perm   string `json:"perm,omitempty"`   // 权限
	Type   int    `json:"type,omitempty"`   // 类型
	Route  string `json:"route,omitempty"`  // 路由/API
	Sort   int    `json:"sort,omitempty"`   // 排序
	Level  int    `json:"level,omitempty"`  // 层级
	Hidden bool   `json:"hidden,omitempty"` // 隐藏
	Remark string `json:"remark,omitempty"` // 备注
	Status int    `json:"status,omitempty"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 菜单分页列表接口响应参数
type SysMenuPageRes struct {
	Items  []*SysMenu `json:"items"`
	Paging *Paging    `json:"paging"`
}

type SysMenu struct {
	Id        string `json:"id,omitempty"`         // ID
	Pid       string `json:"pid,omitempty"`        // 父ID
	Name      string `json:"name,omitempty"`       // 名称
	Perm      string `json:"perm,omitempty"`       // 权限
	Type      int    `json:"type,omitempty"`       // 类型
	Route     string `json:"route,omitempty"`      // 路由/API
	Sort      int    `json:"sort,omitempty"`       // 排序
	Level     int    `json:"level,omitempty"`      // 层级
	Hidden    bool   `json:"hidden,omitempty"`     // 隐藏
	Remark    string `json:"remark,omitempty"`     // 备注
	Status    int    `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Creator   string `json:"creator,omitempty"`    // 创建人
	Updater   string `json:"updater,omitempty"`    // 更新人
	CreatedAt int64  `json:"created_at,omitempty"` // 创建时间
	UpdatedAt int64  `json:"updated_at,omitempty"` // 更新时间
}
