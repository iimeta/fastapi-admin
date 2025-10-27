package model

import "time"

// 代理商修改密码接口请求参数
type ResellerChangePasswordReq struct {
	OldPassword string `json:"old_password,omitempty" v:"required"`
	NewPassword string `json:"new_password,omitempty" v:"required|min-length:6"`
}

// 代理商修改邮箱接口请求参数
type ResellerChangeEmailReq struct {
	Email    string `json:"email,omitempty" v:"required"`
	Code     string `json:"code,omitempty" v:"required|length:0,6"`
	Password string `json:"password,omitempty" v:"required"`
}

// 代理商更新信息接口请求参数
type ResellerUpdateInfoReq struct {
	Name string `json:"name,omitempty" v:"required"`
}

// 新建代理商接口请求参数
type ResellerCreateReq struct {
	Name           string   `json:"name,omitempty" v:"required"`                  // 姓名
	Email          string   `json:"email,omitempty" v:"required"`                 // 邮箱
	Account        string   `json:"account,omitempty" v:"required"`               // 账号
	Password       string   `json:"password,omitempty" v:"required|min-length:6"` // 密码
	Quota          float64  `json:"quota,omitempty"`                              // 额度
	QuotaType      int      `json:"quota_type,omitempty"`                         // 额度类型[1:充值, 2:扣除, 3:赠送, 4:过期]
	QuotaExpiresAt string   `json:"quota_expires_at,omitempty"`                   // 额度过期时间
	Groups         []string `json:"groups,omitempty" d:"[]"`                      // 分组权限
	Remark         string   `json:"remark,omitempty"`                             // 备注
}

// 更新代理商接口请求参数
type ResellerUpdateReq struct {
	Id             string   `json:"id,omitempty"`                        // ID
	Name           string   `json:"name,omitempty" v:"required"`         // 姓名
	Email          string   `json:"email,omitempty" v:"required"`        // 邮箱
	Account        string   `json:"account,omitempty" v:"required"`      // 账号
	Password       string   `json:"password,omitempty" v:"min-length:6"` // 密码
	Quota          int      `json:"quota,omitempty"`                     // 额度
	QuotaExpiresAt string   `json:"quota_expires_at,omitempty"`          // 额度过期时间
	Groups         []string `json:"groups,omitempty"`                    // 分组权限
	Remark         string   `json:"remark,omitempty"`                    // 备注
	Status         int      `json:"status,omitempty" d:"1"`              // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改代理商额度过期时间接口请求参数
type ResellerChangeQuotaExpireReq struct {
	Id             string `json:"id,omitempty"`               // ID
	QuotaExpiresAt string `json:"quota_expires_at,omitempty"` // 额度过期时间
}

// 更改代理商状态接口请求参数
type ResellerChangeStatusReq struct {
	Id     string `json:"id,omitempty"`           // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 删除代理商接口请求参数
type ResellerDeleteReq struct {
	Id   string `json:"id,omitempty"`   // ID
	Data []int  `json:"data,omitempty"` // 删除数据[1:用户数据, 2:应用数据, 3:交易记录, 4:账单明细, 5:日志数据]
}

// 代理商详情接口响应参数
type ResellerDetailRes struct {
	*Reseller
}

// 代理商分页列表接口请求参数
type ResellerPageReq struct {
	Paging
	UserId         int      `json:"user_id,omitempty"`          // 代理商ID
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

// 代理商分页列表接口响应参数
type ResellerPageRes struct {
	Items  []*Reseller `json:"items"`
	Paging *Paging     `json:"paging"`
}

// 代理商列表接口请求参数
type ResellerListReq struct {
	UserId int    `json:"user_id,omitempty"` // 代理商ID
	Name   string `json:"name,omitempty"`    // 姓名
	Status int    `json:"status,omitempty"`  // 状态[1:正常, 2:禁用, -1:删除]
}

// 代理商列表接口响应参数
type ResellerListRes struct {
	Items []*Reseller `json:"items"`
}

// 代理商批量操作接口请求参数
type ResellerBatchOperateReq struct {
	Action         string   `json:"action"`                     // 动作
	Ids            []string `json:"ids"`                        // 主键Ids
	Value          any      `json:"value"`                      // 值
	QuotaType      int      `json:"quota_type"`                 // 额度类型[1:充值, 2:扣除, 3:赠送, 4:过期]
	QuotaExpiresAt string   `json:"quota_expires_at,omitempty"` // 额度过期时间
	IsSendNotice   bool     `json:"is_send_notice,omitempty"`   // 是否发送通知
	Data           []int    `json:"data,omitempty"`             // 删除数据[1:用户数据, 2:应用数据, 3:交易记录, 4:账单明细, 5:日志数据]
}

// 代理商充值接口请求参数
type ResellerRechargeReq struct {
	UserId         int     `json:"user_id,omitempty"`            // 代理商ID
	Quota          float64 `json:"quota,omitempty" v:"required"` // 额度
	QuotaType      int     `json:"quota_type,omitempty"`         // 额度类型[1:充值, 2:扣除, 3:赠送, 4:过期]
	QuotaExpiresAt string  `json:"quota_expires_at,omitempty"`   // 额度过期时间
	IsSendNotice   bool    `json:"is_send_notice,omitempty"`     // 是否发送通知
}

type Reseller struct {
	Id                     string        `json:"id,omitempty"`                       // ID
	UserId                 int           `json:"user_id,omitempty"`                  // 代理商ID
	Name                   string        `json:"name,omitempty"`                     // 姓名
	Avatar                 string        `json:"avatar,omitempty"`                   // 头像
	Email                  string        `json:"email,omitempty"`                    // 邮箱
	Phone                  string        `json:"phone,omitempty"`                    // 手机号
	Quota                  float64       `json:"quota,omitempty"`                    // 剩余额度
	UsedQuota              float64       `json:"used_quota,omitempty"`               // 已用额度
	AllocatedQuota         float64       `json:"allocated_quota,omitempty"`          // 已分配额度
	ToBeAllocatedQuota     float64       `json:"to_be_allocated_quota,omitempty"`    // 待分配额度
	QuotaExpiresAt         string        `json:"quota_expires_at,omitempty"`         // 额度过期时间
	Groups                 []string      `json:"groups,omitempty"`                   // 分组权限
	GroupNames             []string      `json:"group_names,omitempty"`              // 分组名称
	Account                string        `json:"account,omitempty"`                  // 账号
	QuotaWarning           bool          `json:"quota_warning,omitempty"`            // 额度预警
	WarningThreshold       int           `json:"warning_threshold,omitempty"`        // 预警阈值, 单位: $
	ExpireWarningThreshold time.Duration `json:"expire_warning_threshold,omitempty"` // 过期预警阈值, 单位: 天
	WarningNotice          bool          `json:"warning_notice,omitempty"`           // 预警通知
	ExhaustionNotice       bool          `json:"exhaustion_notice,omitempty"`        // 耗尽通知
	ExpireWarningNotice    bool          `json:"expire_warning_notice,omitempty"`    // 额度过期预警通知
	ExpireNotice           bool          `json:"expire_notice,omitempty"`            // 额度过期通知
	Remark                 string        `json:"remark,omitempty"`                   // 备注
	Status                 int           `json:"status,omitempty"`                   // 状态[1:正常, 2:禁用, -1:删除]
	LoginIP                string        `json:"login_ip,omitempty"`                 // 登录IP
	LoginTime              string        `json:"login_time,omitempty"`               // 登录时间
	LoginDomain            string        `json:"login_domain,omitempty"`             // 登录域名
	CreatedAt              string        `json:"created_at,omitempty"`               // 创建时间
	UpdatedAt              string        `json:"updated_at,omitempty"`               // 更新时间
}
