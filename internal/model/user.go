package model

// 登录用户信息接口响应参数
type UserInfoRes struct {
	Id     string `json:"id"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Gender int    `json:"gender"`
	Role   string `json:"role"`
}

// 用户配置信息响应参数
type UserSettingRes struct {
	User    *User        `json:"user,omitempty"`
	Setting *SettingInfo `json:"setting,omitempty"`
}

// 用户信息更新接口请求参数
type UserDetailUpdateReq struct {
	Name     string `json:"name,omitempty" v:"required|max-length:30"`
	Avatar   string `json:"avatar,omitempty"`
	Gender   int    `json:"gender,omitempty" v:"in:0,1,2"`
	Motto    string `json:"motto,omitempty" v:"max-length:1024"`
	Birthday string `json:"birthday,omitempty" v:"length:0,10"`
}

// 用户密码更新接口请求参数
type UserPasswordUpdateReq struct {
	OldPassword string `json:"old_password,omitempty" v:"required"`
	NewPassword string `json:"new_password,omitempty" v:"required|min-length:6"`
}

// 用户手机号更新接口请求参数
type UserPhoneUpdateReq struct {
	Phone    string `json:"phone,omitempty" v:"required"`
	Password string `json:"password,omitempty" v:"required"`
	Code     string `json:"code,omitempty" v:"required|length:0,6"`
}

// 用户邮箱更新接口请求参数
type UserEmailUpdateReq struct {
	Email    string `json:"email,omitempty" v:"required"`
	Password string `json:"password,omitempty" v:"required"`
	Code     string `json:"code,omitempty" v:"required|length:0,6"`
}

// 新建用户接口请求参数
type UserCreateReq struct {
	Name         string   `json:"name,omitempty"`           // 用户名称
	Type         int      `json:"type,omitempty"`           // 用户类型
	Models       []string `json:"models,omitempty"`         // 模型权限
	IsLimitQuota bool     `json:"is_limit_quota,omitempty"` // 是否限制额度
	Quota        int      `json:"quota,omitempty"`          // 额度
	IpWhitelist  string   `json:"ip_whitelist,omitempty"`   // IP白名单
	IpBlacklist  string   `json:"ip_blacklist,omitempty"`   // IP黑名单
	Remark       string   `json:"remark,omitempty"`         // 备注
	Status       int      `json:"status,omitempty" d:"1"`   // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新用户接口请求参数
type UserUpdateReq struct {
	Id           string   `json:"id,omitempty"`             // ID
	Name         string   `json:"name,omitempty"`           // 用户名称
	Type         int      `json:"type,omitempty"`           // 用户类型
	Models       []string `json:"models,omitempty"`         // 模型权限
	IsLimitQuota bool     `json:"is_limit_quota,omitempty"` // 是否限制额度
	Quota        int      `json:"quota,omitempty"`          // 额度
	IpWhitelist  string   `json:"ip_whitelist,omitempty"`   // IP白名单
	IpBlacklist  string   `json:"ip_blacklist,omitempty"`   // IP黑名单
	Remark       string   `json:"remark,omitempty"`         // 备注
	Status       int      `json:"status,omitempty" d:"1"`   // 状态[1:正常, 2:禁用, -1:删除]
}

// 用户详情接口响应参数
type UserDetailRes struct {
	*User
}

// 用户分页列表接口请求参数
type UserPageReq struct {
	Paging
	UserId      int      `json:"user_id,omitempty"`      // 用户ID
	Name        string   `json:"name,omitempty"`         // 用户名称
	Type        int      `json:"type,omitempty"`         // 用户类型
	Models      []string `json:"models,omitempty"`       // 模型权限
	Quota       int      `json:"quota,omitempty"`        // 额度
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt   []string `json:"created_at,omitempty"`   // 创建时间
}

// 用户分页列表接口响应参数
type UserPageRes struct {
	Items  []*User `json:"items"`
	Paging *Paging `json:"paging"`
}

// 用户列表接口请求参数
type UserListReq struct {
	UserId      int      `json:"user_id,omitempty"`      // 用户ID
	Name        string   `json:"name,omitempty"`         // 用户名称
	Type        int      `json:"type,omitempty"`         // 用户类型
	Models      []string `json:"models,omitempty"`       // 模型权限
	Quota       int      `json:"quota,omitempty"`        // 额度
	IpWhitelist []string `json:"ip_whitelist,omitempty"` // IP白名单
	IpBlacklist []string `json:"ip_blacklist,omitempty"` // IP黑名单
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
}

// 用户列表接口响应参数
type UserListRes struct {
	Items []*User `json:"items"`
}

type User struct {
	Id        string `json:"id,omitempty"`         // ID
	UserId    int    `json:"user_id,omitempty"`    // 用户ID
	Name      string `json:"name,omitempty"`       // 姓名
	Avatar    string `json:"avatar,omitempty"`     // 用户头像地址
	Gender    int    `json:"gender,omitempty"`     // 用户性别  0:未知  1:男   2:女
	Phone     string `json:"phone,omitempty"`      // 手机号
	Email     string `json:"email,omitempty"`      // 用户邮箱
	Quota     int    `json:"quota,omitempty"`      // 额度
	Remark    string `json:"remark,omitempty"`     // 备注
	CreatedAt int64  `json:"created_at,omitempty"` // 创建时间
	UpdatedAt int64  `json:"updated_at,omitempty"` // 更新时间
}

type SettingInfo struct {
	ThemeMode           string `json:"theme_mode,omitempty"`
	ThemeBagImg         string `json:"theme_bag_img,omitempty"`
	ThemeColor          string `json:"theme_color,omitempty"`
	NotifyCueTone       string `json:"notify_cue_tone,omitempty"`
	KeyboardEventNotify string `json:"keyboard_event_notify,omitempty"`
}
