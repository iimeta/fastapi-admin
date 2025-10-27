package model

import "github.com/iimeta/fastapi-admin/internal/model/common"

// 新建分组接口请求参数
type GroupCreateReq struct {
	Name               string                `json:"name,omitempty"`                  // 分组名称
	Discount           float64               `json:"discount,omitempty"`              // 分组折扣
	Models             []string              `json:"models,omitempty" d:"[]"`         // 模型权限
	IsEnableModelAgent bool                  `json:"is_enable_model_agent,omitempty"` // 是否启用模型代理
	LbStrategy         int                   `json:"lb_strategy,omitempty" d:"1"`     // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents        []string              `json:"model_agents,omitempty" d:"[]"`   // 模型代理
	IsDefault          bool                  `json:"is_default,omitempty"`            // 是否默认分组
	IsLimitQuota       bool                  `json:"is_limit_quota,omitempty"`        // 是否限制额度
	Quota              float64               `json:"quota,omitempty"`                 // 额度
	IsEnableForward    bool                  `json:"is_enable_forward,omitempty"`     // 是否启用模型转发
	ForwardConfig      *common.ForwardConfig `json:"forward_config,omitempty"`        // 模型转发配置
	IsPublic           bool                  `json:"is_public,omitempty"`             // 是否公开
	Weight             int                   `json:"weight,omitempty"`                // 权重
	ExpiresAt          string                `json:"expires_at,omitempty"`            // 过期时间
	Remark             string                `json:"remark,omitempty"`                // 备注
	Status             int                   `json:"status,omitempty" d:"1"`          // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新分组接口请求参数
type GroupUpdateReq struct {
	Id                 string                `json:"id,omitempty"`                    // ID
	Name               string                `json:"name,omitempty"`                  // 分组名称
	Discount           float64               `json:"discount,omitempty"`              // 分组折扣
	Models             []string              `json:"models,omitempty" d:"[]"`         // 模型权限
	IsEnableModelAgent bool                  `json:"is_enable_model_agent,omitempty"` // 是否启用模型代理
	LbStrategy         int                   `json:"lb_strategy,omitempty" d:"1"`     // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents        []string              `json:"model_agents,omitempty" d:"[]"`   // 模型代理
	IsDefault          bool                  `json:"is_default,omitempty"`            // 是否默认分组
	IsLimitQuota       bool                  `json:"is_limit_quota,omitempty"`        // 是否限制额度
	Quota              float64               `json:"quota,omitempty"`                 // 额度
	IsEnableForward    bool                  `json:"is_enable_forward,omitempty"`     // 是否启用模型转发
	ForwardConfig      *common.ForwardConfig `json:"forward_config,omitempty"`        // 模型转发配置
	IsPublic           bool                  `json:"is_public,omitempty"`             // 是否公开
	Weight             int                   `json:"weight,omitempty"`                // 权重
	ExpiresAt          string                `json:"expires_at,omitempty"`            // 过期时间
	Remark             string                `json:"remark,omitempty"`                // 备注
	Status             int                   `json:"status,omitempty" d:"1"`          // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改过期时间接口请求参数
type GroupChangeExpireReq struct {
	Id        string `json:"id,omitempty"`         // ID
	ExpiresAt string `json:"expires_at,omitempty"` // 过期时间
}

// 更改分组公开状态接口请求参数
type GroupChangePublicReq struct {
	Id       string `json:"id" v:"required"`     // ID
	IsPublic bool   `json:"is_public,omitempty"` // 是否公开
}

// 更改分组状态接口请求参数
type GroupChangeStatusReq struct {
	Id     string `json:"id,omitempty"`           // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 分组详情接口响应参数
type GroupDetailRes struct {
	*Group
}

// 分组分页列表接口请求参数
type GroupPageReq struct {
	Paging
	Name        string   `json:"name,omitempty"`         // 分组名称
	Models      []string `json:"models,omitempty"`       // 模型
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常, 2:禁用, -1:删除]
	ExpiresAt   []string `json:"created_at,omitempty"`   // 过期时间
}

// 分组分页列表接口响应参数
type GroupPageRes struct {
	Items  []*Group `json:"items"`
	Paging *Paging  `json:"paging"`
}

// 分组列表接口请求参数
type GroupListReq struct {
	Name string `json:"name,omitempty"` // 分组名称
}

// 分组列表接口响应参数
type GroupListRes struct {
	Items []*Group `json:"items"`
}

// 分组批量操作接口请求参数
type GroupBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

type Group struct {
	Id                 string                `json:"id,omitempty"`                    // ID
	Name               string                `json:"name,omitempty"`                  // 分组名称
	Discount           float64               `json:"discount,omitempty"`              // 分组折扣
	Models             []string              `json:"models,omitempty"`                // 模型权限
	ModelNames         []string              `json:"model_names,omitempty"`           // 模型名称
	IsEnableModelAgent bool                  `json:"is_enable_model_agent,omitempty"` // 是否启用模型代理
	LbStrategy         int                   `json:"lb_strategy,omitempty"`           // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents        []string              `json:"model_agents,omitempty"`          // 模型代理
	ModelAgentNames    []string              `json:"model_agent_names,omitempty"`     // 模型代理名称
	IsDefault          bool                  `json:"is_default,omitempty"`            // 是否默认分组
	IsLimitQuota       bool                  `json:"is_limit_quota,omitempty"`        // 是否限制额度
	Quota              float64               `json:"quota,omitempty"`                 // 剩余额度
	UsedQuota          float64               `json:"used_quota,omitempty"`            // 已用额度
	IsEnableForward    bool                  `json:"is_enable_forward,omitempty"`     // 是否启用模型转发
	ForwardConfig      *common.ForwardConfig `json:"forward_config,omitempty"`        // 模型转发配置
	IsPublic           bool                  `json:"is_public,omitempty"`             // 是否公开
	Weight             int                   `json:"weight,omitempty"`                // 权重
	ExpiresAt          string                `json:"expires_at,omitempty"`            // 过期时间
	Remark             string                `json:"remark,omitempty"`                // 备注
	Status             int                   `json:"status,omitempty"`                // 状态[1:正常, 2:禁用, -1:删除]
	Creator            string                `json:"creator,omitempty"`               // 创建人
	Updater            string                `json:"updater,omitempty"`               // 更新人
	CreatedAt          string                `json:"created_at,omitempty"`            // 创建时间
	UpdatedAt          string                `json:"updated_at,omitempty"`            // 更新时间
}
