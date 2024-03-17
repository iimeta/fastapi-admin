package model

// 新建管理员接口请求参数
type SysAdminCreateReq struct {
	Name     string `json:"name,omitempty"`     // 名称
	Avatar   string `json:"avatar,omitempty"`   // 头像
	Phone    string `json:"phone,omitempty"`    // 手机号
	Email    string `json:"email,omitempty"`    // 邮箱
	Account  string `json:"account,omitempty"`  // 账号
	Password string `json:"password,omitempty"` // 密码
	Remark   string `json:"remark,omitempty"`   // 备注
	Status   int    `json:"status,omitempty"`   // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新管理员接口请求参数
type SysAdminUpdateReq struct {
	Id       string `json:"id,omitempty"`       // ID
	Name     string `json:"name,omitempty"`     // 名称
	Avatar   string `json:"avatar,omitempty"`   // 头像
	Phone    string `json:"phone,omitempty"`    // 手机号
	Email    string `json:"email,omitempty"`    // 邮箱
	Account  string `json:"account,omitempty"`  // 账号
	Password string `json:"password,omitempty"` // 密码
	Remark   string `json:"remark,omitempty"`   // 备注
	Status   int    `json:"status,omitempty"`   // 状态[1:正常, 2:禁用, -1:删除]
}

// 管理员详情接口响应参数
type SysAdminDetailRes struct {
	*SysAdmin
}

// 管理员分页列表接口请求参数
type SysAdminPageReq struct {
	Paging
	Name     string `json:"name,omitempty"`     // 名称
	Avatar   string `json:"avatar,omitempty"`   // 头像
	Phone    string `json:"phone,omitempty"`    // 手机号
	Email    string `json:"email,omitempty"`    // 邮箱
	Account  string `json:"account,omitempty"`  // 账号
	Password string `json:"password,omitempty"` // 密码
	Remark   string `json:"remark,omitempty"`   // 备注
	Status   int    `json:"status,omitempty"`   // 状态[1:正常, 2:禁用, -1:删除]
}

// 管理员分页列表接口响应参数
type SysAdminPageRes struct {
	Items  []*SysAdmin `json:"items"`
	Paging *Paging     `json:"paging"`
}

type SysAdmin struct {
	Id        string `json:"id,omitempty"`         // ID
	Name      string `json:"name,omitempty"`       // 名称
	Avatar    string `json:"avatar,omitempty"`     // 头像
	Phone     string `json:"phone,omitempty"`      // 手机号
	Email     string `json:"email,omitempty"`      // 邮箱
	Account   string `json:"account,omitempty"`    // 账号
	Password  string `json:"password,omitempty"`   // 密码
	Salt      string `json:"salt,omitempty"`       // 盐
	LoginIP   string `json:"login_ip,omitempty"`   // 登录IP
	LoginTime int64  `json:"login_time,omitempty"` // 登录时间
	Remark    string `json:"remark,omitempty"`     // 备注
	Status    int    `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Creator   string `json:"creator,omitempty"`    // 创建人
	Updater   string `json:"updater,omitempty"`    // 更新人
	CreatedAt int64  `json:"created_at,omitempty"` // 创建时间
	UpdatedAt int64  `json:"updated_at,omitempty"` // 更新时间
}
