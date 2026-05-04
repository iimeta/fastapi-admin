package model

// 创建工单接口请求参数
type TicketCreateReq struct {
	Title    string `json:"title" v:"required"`       // 工单标题
	Content  string `json:"content" v:"required"`     // 工单内容
	Category string `json:"category" v:"required"`    // 分类[account:账户问题, billing:计费问题, technical:技术问题, feature:功能建议, other:其他]
	Priority int    `json:"priority,omitempty" d:"2"` // 优先级[1:低, 2:中, 3:高, 4:紧急]
}

// 回复工单接口请求参数
type TicketReplyReq struct {
	TicketId string `json:"ticket_id" v:"required"` // 工单ID
	Content  string `json:"content" v:"required"`   // 回复内容
}

// 更新工单状态接口请求参数
type TicketUpdateStatusReq struct {
	Id     string `json:"id" v:"required"`     // 工单ID
	Status int    `json:"status" v:"required"` // 状态[1:待回复, 2:待处理, 3:处理中, 4:已回复, 5:已解决, 6:已关闭]
}

// 关闭工单接口请求参数
type TicketCloseReq struct {
	Id string `json:"id" v:"required"` // 工单ID
}

// 分配工单接口请求参数
type TicketAssignReq struct {
	Id         string `json:"id" v:"required"`          // 工单ID
	AssigneeId int    `json:"assignee_id" v:"required"` // 处理人ID
}

// 工单分页列表接口请求参数
type TicketPageReq struct {
	Paging
	TicketNo       string `json:"ticket_no,omitempty"`        // 工单编号
	Title          string `json:"title,omitempty"`            // 工单标题
	UserName       string `json:"user_name,omitempty"`        // 提交人
	Category       string `json:"category,omitempty"`         // 分类[account:账户问题, billing:计费问题, technical:技术问题, feature:功能建议, other:其他]
	Priority       int    `json:"priority,omitempty"`         // 优先级[1:低, 2:中, 3:高, 4:紧急]
	Status         int    `json:"status,omitempty"`           // 状态[1:待回复, 2:待处理, 3:处理中, 4:已回复, 5:已解决, 6:已关闭]
	Scope          string `json:"scope,omitempty"`            // 范围[my:我的工单]
	CreatedAtStart string `json:"created_at_start,omitempty"` // 创建时间开始
	CreatedAtEnd   string `json:"created_at_end,omitempty"`   // 创建时间结束
}

// 工单分页列表接口响应参数
type TicketPageRes struct {
	Items  []*Ticket `json:"items"`
	Paging *Paging   `json:"paging"`
}

// 工单详情接口响应参数
type TicketDetailRes struct {
	*Ticket
	Replies []*TicketReply `json:"replies,omitempty"`
}

// 批量删除工单接口请求参数
type TicketBatchDeleteReq struct {
	Ids []string `json:"ids" v:"required"` // 工单ID列表
}

type Ticket struct {
	Id           string `json:"id,omitempty"`            // ID
	TicketNo     string `json:"ticket_no,omitempty"`     // 工单编号
	Title        string `json:"title,omitempty"`         // 工单标题
	Content      string `json:"content,omitempty"`       // 工单内容
	Category     string `json:"category,omitempty"`      // 分类[account:账户问题, billing:计费问题, technical:技术问题, feature:功能建议, other:其他]
	Priority     int    `json:"priority,omitempty"`      // 优先级[1:低, 2:中, 3:高, 4:紧急]
	Status       int    `json:"status,omitempty"`        // 状态[1:待回复, 2:待处理, 3:处理中, 4:已回复, 5:已解决, 6:已关闭]
	UserId       int    `json:"user_id,omitempty"`       // 提交用户ID
	UserName     string `json:"user_name,omitempty"`     // 提交用户名
	UserRole     string `json:"user_role,omitempty"`     // 提交用户角色[user:用户, reseller:代理商]
	AssigneeId   int    `json:"assignee_id,omitempty"`   // 处理人ID
	AssigneeRole string `json:"assignee_role,omitempty"` // 处理人角色[reseller:代理商, admin:管理员]
	ReplyCount   int    `json:"reply_count,omitempty"`   // 回复数量
	LastReplyAt  string `json:"last_reply_at,omitempty"` // 最后回复时间
	Rid          int    `json:"rid,omitempty"`           // 代理商ID
	Creator      string `json:"creator,omitempty"`       // 创建人
	Updater      string `json:"updater,omitempty"`       // 更新人
	CreatedAt    string `json:"created_at,omitempty"`    // 创建时间
	UpdatedAt    string `json:"updated_at,omitempty"`    // 更新时间
}

type TicketReply struct {
	Id        string `json:"id,omitempty"`         // ID
	TicketId  string `json:"ticket_id,omitempty"`  // 工单ID
	Content   string `json:"content,omitempty"`    // 回复内容
	UserId    int    `json:"user_id,omitempty"`    // 回复者用户ID
	UserName  string `json:"user_name,omitempty"`  // 回复者用户名
	Role      string `json:"role,omitempty"`       // 回复者角色[user:用户, reseller:代理商, admin:管理员]
	CreatedAt string `json:"created_at,omitempty"` // 创建时间
}
