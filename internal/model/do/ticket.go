package do

import "github.com/gogf/gf/v2/util/gmeta"

type Ticket struct {
	gmeta.Meta     `collection:"ticket" bson:"-"`
	TicketNo       string `bson:"ticket_no,omitempty"`     // 工单编号
	Title          string `bson:"title,omitempty"`         // 工单标题
	Content        string `bson:"content,omitempty"`       // 工单内容
	Category       string `bson:"category,omitempty"`      // 分类[account:账户问题, billing:计费问题, technical:技术问题, feature:功能建议, other:其他]
	Priority       int    `bson:"priority,omitempty"`      // 优先级[1:低, 2:中, 3:高, 4:紧急]
	Status         int    `bson:"status,omitempty"`        // 状态[1:待回复, 2:待处理, 3:处理中, 4:已回复, 5:已解决, 6:已关闭]
	UserId         int    `bson:"user_id,omitempty"`       // 提交用户ID
	UserName       string `bson:"user_name,omitempty"`     // 提交用户名
	UserRole       string `bson:"user_role,omitempty"`     // 提交用户角色[user:用户, reseller:代理商]
	AssigneeId     int    `bson:"assignee_id"`             // 处理人ID
	AssigneeRole   string `bson:"assignee_role,omitempty"` // 处理人角色[reseller:代理商, admin:管理员]
	ReplyCount     int    `bson:"reply_count"`             // 回复数量
	LastReplyAt    int64  `bson:"last_reply_at"`           // 最后回复时间
	UserNotice     bool   `bson:"user_notice"`             // 提交方提醒
	AssigneeNotice bool   `bson:"assignee_notice"`         // 处理方提醒
	Rid            int    `bson:"rid,omitempty"`           // 代理商ID
	Creator        string `bson:"creator,omitempty"`       // 创建人
	Updater        string `bson:"updater,omitempty"`       // 更新人
	CreatedAt      int64  `bson:"created_at,omitempty"`    // 创建时间
	UpdatedAt      int64  `bson:"updated_at,omitempty"`    // 更新时间
}

type TicketReply struct {
	gmeta.Meta `collection:"ticket_reply" bson:"-"`
	TicketId   string `bson:"ticket_id,omitempty"`  // 工单ID
	Content    string `bson:"content,omitempty"`    // 回复内容
	UserId     int    `bson:"user_id,omitempty"`    // 回复者用户ID
	UserName   string `bson:"user_name,omitempty"`  // 回复者用户名
	Role       string `bson:"role,omitempty"`       // 回复者角色[user:用户, reseller:代理商, admin:管理员]
	Rid        int    `bson:"rid,omitempty"`        // 代理商ID
	Creator    string `bson:"creator,omitempty"`    // 创建人
	Updater    string `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}
