package model

import "github.com/iimeta/fastapi-admin/internal/model/common"

// 更新配置接口请求参数
type SysConfigUpdateReq struct {
	Id         string             `json:"id,omitempty"`         // ID
	Action     string             `json:"action,omitempty"`     // 动作
	Core       *common.Core       `json:"core,omitempty"`       // 核心
	Http       *common.Http       `json:"http,omitempty"`       // HTTP
	Email      *common.Email      `json:"email,omitempty"`      // 邮箱
	Statistics *common.Statistics `json:"statistics,omitempty"` // 统计
	Api        *common.Api        `json:"api,omitempty"`        // API
	Midjourney *common.Midjourney `json:"midjourney,omitempty"` // Midjourney
	Gcp        *common.Gcp        `json:"gcp,omitempty"`        // GCP
	Log        *common.Log        `json:"log,omitempty"`        // 日志
	Error      *common.Error      `json:"error,omitempty"`      // 错误配置
	Debug      *common.Debug      `json:"debug,omitempty"`      // 调试
}

// 更改配置状态接口请求参数
type SysConfigChangeStatusReq struct {
	Id     string `json:"id,omitempty"`     // ID
	Action string `json:"action,omitempty"` // 动作
	Status int    `json:"status,omitempty"` // 状态
	Open   bool   `json:"open,omitempty"`   // 开关
}

// 重置配置接口请求参数
type SysConfigResetReq struct {
	Id     string `json:"id,omitempty"`     // ID
	Action string `json:"action,omitempty"` // 动作
}

// 配置详情接口响应参数
type SysConfigDetailRes struct {
	*SysConfig
}

type SysConfig struct {
	Id         string             `json:"id,omitempty"`         // ID
	Core       *common.Core       `json:"core,omitempty"`       // 核心
	Http       *common.Http       `json:"http,omitempty"`       // HTTP
	Email      *common.Email      `json:"email,omitempty"`      // 邮箱
	Statistics *common.Statistics `json:"statistics,omitempty"` // 统计
	Api        *common.Api        `json:"api,omitempty"`        // API
	Midjourney *common.Midjourney `json:"midjourney,omitempty"` // Midjourney
	Gcp        *common.Gcp        `json:"gcp,omitempty"`        // GCP
	Log        *common.Log        `json:"log,omitempty"`        // 日志
	Error      *common.Error      `json:"error,omitempty"`      // 错误配置
	Debug      *common.Debug      `json:"debug,omitempty"`      // 调试
	Creator    string             `json:"creator,omitempty"`    // 创建人
	Updater    string             `json:"updater,omitempty"`    // 更新人
	CreatedAt  string             `json:"created_at,omitempty"` // 创建时间
	UpdatedAt  string             `json:"updated_at,omitempty"` // 更新时间
}
