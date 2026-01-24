package model

import "github.com/iimeta/fastapi-admin/v2/internal/model/common"

// 更新配置接口请求参数
type SysConfigUpdateReq struct {
	Action                string                        `json:"action,omitempty"`                  // 动作
	Core                  *common.Core                  `json:"core,omitempty"`                    // 核心
	Http                  *common.Http                  `json:"http,omitempty"`                    // HTTP
	Email                 *common.Email                 `json:"email,omitempty"`                   // 邮箱
	Statistics            *common.Statistics            `json:"statistics,omitempty"`              // 统计
	Base                  *common.Base                  `json:"base,omitempty"`                    // 基础
	Midjourney            *common.Midjourney            `json:"midjourney,omitempty"`              // Midjourney
	Log                   *common.Log                   `json:"log,omitempty"`                     // 日志
	UserLoginRegister     *common.UserLoginRegister     `json:"user_login_register,omitempty"`     // 用户登录注册
	UserShieldError       *common.UserShieldError       `json:"user_shield_error,omitempty"`       // 用户屏蔽错误
	ResellerLoginRegister *common.ResellerLoginRegister `json:"reseller_login_register,omitempty"` // 代理商登录注册
	ResellerShieldError   *common.ResellerShieldError   `json:"reseller_shield_error,omitempty"`   // 代理商屏蔽错误
	AdminLogin            *common.AdminLogin            `json:"admin_login,omitempty"`             // 管理员登录
	AutoDisabledError     *common.AutoDisabledError     `json:"auto_disabled_error,omitempty"`     // 自动禁用错误
	AutoEnableError       *common.AutoEnableError       `json:"auto_enable_error,omitempty"`       // 自动启用错误
	NotRetryError         *common.NotRetryError         `json:"not_retry_error,omitempty"`         // 不重试错误
	NotShieldError        *common.NotShieldError        `json:"not_shield_error,omitempty"`        // 不屏蔽错误
	Notice                *common.Notice                `json:"notice,omitempty"`                  // 通知
	Quota                 *common.Quota                 `json:"quota,omitempty"`                   // 额度
	QuotaTask             *common.QuotaTask             `json:"quota_task,omitempty"`              // 额度任务
	VideoTask             *common.VideoTask             `json:"video_task,omitempty"`              // 视频任务
	FileTask              *common.FileTask              `json:"file_task,omitempty"`               // 文件任务
	BatchTask             *common.BatchTask             `json:"batch_task,omitempty"`              // 批处理任务
	ServiceUnavailable    *common.ServiceUnavailable    `json:"service_unavailable,omitempty"`     // 暂停服务
	GeneralApi            *common.GeneralApi            `json:"general_api,omitempty"`             // 通用API
	Test                  *common.Test                  `json:"test,omitempty"`                    // 测试
	Debug                 *common.Debug                 `json:"debug,omitempty"`                   // 调试
}

// 更改配置状态接口请求参数
type SysConfigChangeStatusReq struct {
	Action string `json:"action,omitempty"` // 动作
	Open   bool   `json:"open,omitempty"`   // 开关
}

// 重置配置接口请求参数
type SysConfigResetReq struct {
	Action string `json:"action,omitempty"` // 动作
}

// 刷新配置接口请求参数
type SysConfigRefreshReq struct {
	Action string `json:"action,omitempty"` // 动作
}

// 配置详情接口响应参数
type SysConfigDetailRes struct {
	*SysConfig
}

type SysConfig struct {
	Id                    string                        `json:"id,omitempty"`                      // ID
	Core                  *common.Core                  `json:"core,omitempty"`                    // 核心
	Http                  *common.Http                  `json:"http,omitempty"`                    // HTTP
	Email                 *common.Email                 `json:"email,omitempty"`                   // 邮箱
	Statistics            *common.Statistics            `json:"statistics,omitempty"`              // 统计
	Base                  *common.Base                  `json:"base,omitempty"`                    // 基础
	Midjourney            *common.Midjourney            `json:"midjourney,omitempty"`              // Midjourney
	Log                   *common.Log                   `json:"log,omitempty"`                     // 日志
	UserLoginRegister     *common.UserLoginRegister     `json:"user_login_register,omitempty"`     // 用户登录注册
	UserShieldError       *common.UserShieldError       `json:"user_shield_error,omitempty"`       // 用户屏蔽错误
	ResellerLoginRegister *common.ResellerLoginRegister `json:"reseller_login_register,omitempty"` // 代理商登录注册
	ResellerShieldError   *common.ResellerShieldError   `json:"reseller_shield_error,omitempty"`   // 代理商屏蔽错误
	AdminLogin            *common.AdminLogin            `json:"admin_login,omitempty"`             // 管理员登录
	AutoDisabledError     *common.AutoDisabledError     `json:"auto_disabled_error,omitempty"`     // 自动禁用错误
	AutoEnableError       *common.AutoEnableError       `json:"auto_enable_error,omitempty"`       // 自动启用错误
	NotRetryError         *common.NotRetryError         `json:"not_retry_error,omitempty"`         // 不重试错误
	NotShieldError        *common.NotShieldError        `json:"not_shield_error,omitempty"`        // 不屏蔽错误
	Notice                *common.Notice                `json:"notice,omitempty"`                  // 通知
	Quota                 *common.Quota                 `json:"quota,omitempty"`                   // 额度
	QuotaTask             *common.QuotaTask             `json:"quota_task,omitempty"`              // 额度任务
	VideoTask             *common.VideoTask             `json:"video_task,omitempty"`              // 视频任务
	FileTask              *common.FileTask              `json:"file_task,omitempty"`               // 文件任务
	BatchTask             *common.BatchTask             `json:"batch_task,omitempty"`              // 批处理任务
	ServiceUnavailable    *common.ServiceUnavailable    `json:"service_unavailable,omitempty"`     // 暂停服务
	GeneralApi            *common.GeneralApi            `json:"general_api,omitempty"`             // 通用API
	Test                  *common.Test                  `json:"test,omitempty"`                    // 测试
	Debug                 *common.Debug                 `json:"debug,omitempty"`                   // 调试
	Creator               string                        `json:"creator,omitempty"`                 // 创建人
	Updater               string                        `json:"updater,omitempty"`                 // 更新人
	CreatedAt             string                        `json:"created_at,omitempty"`              // 创建时间
	UpdatedAt             string                        `json:"updated_at,omitempty"`              // 更新时间
}
