package model

import "time"

// 用户信息接口响应参数
type UserInfoRes struct {
	UserId    int    `json:"user_id"`
	Account   string `json:"account"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"` // 注册时间
}

// 用户修改密码接口请求参数
type UserChangePasswordReq struct {
	OldPassword string `json:"old_password,omitempty" v:"required"`
	NewPassword string `json:"new_password,omitempty" v:"required|min-length:6"`
}

// 用户修改邮箱接口请求参数
type UserChangeEmailReq struct {
	Email    string `json:"email,omitempty" v:"required"`
	Code     string `json:"code,omitempty" v:"required|length:0,6"`
	Password string `json:"password,omitempty" v:"required"`
}

// 用户更新信息接口请求参数
type UserUpdateInfoReq struct {
	Name string `json:"name,omitempty" v:"required"`
}

// 新建用户接口请求参数
type UserCreateReq struct {
	Name           string `json:"name,omitempty" v:"required"`                               // 姓名
	Account        string `json:"account,omitempty" v:"required"`                            // 账号
	Password       string `json:"password,omitempty" v:"required|min-length:6"`              // 密码
	Terminal       string `json:"terminal,omitempty" v:"required|in:web,h5,ios,windows,mac"` // 终端
	Quota          int    `json:"quota,omitempty"`                                           // 额度
	QuotaExpiresAt string `json:"quota_expires_at,omitempty"`                                // 额度过期时间
	Remark         string `json:"remark,omitempty"`                                          // 备注
}

// 更新用户接口请求参数
type UserUpdateReq struct {
	Id             string   `json:"id,omitempty"`                        // ID
	Name           string   `json:"name,omitempty" v:"required"`         // 姓名
	Account        string   `json:"account,omitempty" v:"required"`      // 账号
	Password       string   `json:"password,omitempty" v:"min-length:6"` // 密码
	Type           int      `json:"type,omitempty"`                      // 用户类型
	Models         []string `json:"models,omitempty" d:"[]"`             // 模型权限
	IsLimitQuota   bool     `json:"is_limit_quota,omitempty"`            // 是否限制额度
	Quota          int      `json:"quota,omitempty"`                     // 额度
	QuotaExpiresAt string   `json:"quota_expires_at,omitempty"`          // 额度过期时间
	IpWhitelist    string   `json:"ip_whitelist,omitempty"`              // IP白名单
	IpBlacklist    string   `json:"ip_blacklist,omitempty"`              // IP黑名单
	Remark         string   `json:"remark,omitempty"`                    // 备注
	Status         int      `json:"status,omitempty" d:"1"`              // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改用户额度过期时间接口请求参数
type UserChangeQuotaExpireReq struct {
	Id             string `json:"id,omitempty"`               // ID
	QuotaExpiresAt string `json:"quota_expires_at,omitempty"` // 额度过期时间
}

// 更改用户状态接口请求参数
type UserChangeStatusReq struct {
	Id     string `json:"id,omitempty"`           // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 用户详情接口响应参数
type UserDetailRes struct {
	*User
}

// 用户分页列表接口请求参数
type UserPageReq struct {
	Paging
	UserId         int      `json:"user_id,omitempty"`          // 用户ID
	Name           string   `json:"name,omitempty"`             // 姓名
	Phone          string   `json:"phone,omitempty"`            // 手机号
	Email          string   `json:"email,omitempty"`            // 邮箱
	Account        string   `json:"account,omitempty"`          // 账号
	Quota          int      `json:"quota,omitempty"`            // 额度
	QuotaExpiresAt []string `json:"quota_expires_at,omitempty"` // 额度过期时间
	Remark         string   `json:"remark,omitempty"`           // 备注
	Status         int      `json:"status,omitempty"`           // 状态[1:正常, 2:禁用, -1:删除]
	CreatedAt      []string `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt      []string `json:"updated_at,omitempty"`       // 更新时间
}

// 用户分页列表接口响应参数
type UserPageRes struct {
	Items  []*User `json:"items"`
	Paging *Paging `json:"paging"`
}

// 用户列表接口请求参数
type UserListReq struct {
	UserId      int      `json:"user_id,omitempty"`      // 用户ID
	Name        string   `json:"name,omitempty"`         // 姓名
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

// 授予用户额度接口请求参数
type UserGrantQuotaReq struct {
	UserId         int    `json:"user_id,omitempty"`            // 用户ID
	Quota          int    `json:"quota,omitempty" v:"required"` // 额度
	QuotaExpiresAt string `json:"quota_expires_at,omitempty"`   // 额度过期时间
}

// 用户模型权限接口请求参数
type UserModelsReq struct {
	UserId int      `json:"user_id,omitempty"`       // 用户ID
	Models []string `json:"models,omitempty" d:"[]"` // 模型权限
}

type User struct {
	Id                     string        `json:"id,omitempty"`                // ID
	UserId                 int           `json:"user_id,omitempty"`           // 用户ID
	Name                   string        `json:"name,omitempty"`              // 姓名
	Avatar                 string        `json:"avatar,omitempty"`            // 头像
	Email                  string        `json:"email,omitempty"`             // 邮箱
	Phone                  string        `json:"phone,omitempty"`             // 手机号
	Quota                  int           `json:"quota"`                       // 剩余额度
	UsedQuota              int           `json:"used_quota"`                  // 已用额度
	QuotaExpiresAt         string        `json:"quota_expires_at"`            // 额度过期时间
	Models                 []string      `json:"models,omitempty"`            // 模型权限
	ModelNames             []string      `json:"model_names,omitempty"`       // 模型名称
	Account                string        `json:"account,omitempty"`           // 账号
	QuotaWarning           bool          `json:"quota_warning,omitempty"`     // 额度预警
	WarningThreshold       int           `json:"warning_threshold,omitempty"` // 预警阈值
	ExpireWarningThreshold time.Duration `json:"expire_warning_threshold"`    // 过期预警阈值, 单位: 天
	WarningNotice          bool          `json:"warning_notice,omitempty"`    // 预警通知
	ExhaustionNotice       bool          `json:"exhaustion_notice,omitempty"` // 耗尽通知
	Remark                 string        `json:"remark,omitempty"`            // 备注
	Status                 int           `json:"status,omitempty"`            // 状态[1:正常, 2:禁用, -1:删除]
	LoginIP                string        `json:"login_ip,omitempty"`          // 登录IP
	LoginTime              string        `json:"login_time,omitempty"`        // 登录时间
	LoginDomain            string        `json:"login_domain,omitempty"`      // 登录域名
	CreatedAt              string        `json:"created_at,omitempty"`        // 创建时间
	UpdatedAt              string        `json:"updated_at,omitempty"`        // 更新时间
}
