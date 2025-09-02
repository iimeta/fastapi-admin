package model

// 新建消息通知接口请求参数
type NoticeCreateReq struct {
	Title         string   `json:"title,omitempty"`           // 标题
	Content       string   `json:"content,omitempty"`         // 内容
	Category      string   `json:"category,omitempty"`        // 分类[service:服务消息, activity:活动消息, safety:安全消息, maint:维护消息, product:产品消息, fault:故障消息]
	Scope         int      `json:"scope,omitempty"`           // 通知范围[1:全部, 2:全部用户, 3:全部代理商, 4:指定用户, 5:指定代理商, 6:指定用户和代理商]
	Users         []int    `json:"users,omitempty"`           // 通知用户
	Resellers     []int    `json:"resellers,omitempty"`       // 通知代理商
	Channels      []string `json:"channels,omitempty"`        // 发送渠道[web:站内信, email:邮件]
	IsPopup       bool     `json:"is_popup,omitempty"`        // 是否弹窗
	Priority      int      `json:"priority,omitempty" d:"20"` // 优先级
	ExpiresAt     string   `json:"expires_at,omitempty"`      // 过期时间
	ScheduledTime string   `json:"scheduled_time,omitempty"`  // 定时发布时间
	Remark        string   `json:"remark,omitempty"`          // 备注
	Status        int      `json:"status,omitempty"`          // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
}

// 更新消息通知接口请求参数
type NoticeUpdateReq struct {
	Id            string   `json:"id" v:"required"`          // ID
	Title         string   `json:"title,omitempty"`          // 标题
	Content       string   `json:"content,omitempty"`        // 内容
	Category      string   `json:"category,omitempty"`       // 分类[service:服务消息, activity:活动消息, safety:安全消息, maint:维护消息, product:产品消息, fault:故障消息]
	Scope         int      `json:"scope,omitempty"`          // 通知范围[1:全部, 2:全部用户, 3:全部代理商, 4:指定用户, 5:指定代理商, 6:指定用户和代理商]
	Users         []int    `json:"users,omitempty"`          // 通知用户
	Resellers     []int    `json:"resellers,omitempty"`      // 通知代理商
	Channels      []string `json:"channels,omitempty"`       // 发送渠道[web:站内信, email:邮件]
	IsPopup       bool     `json:"is_popup,omitempty"`       // 是否弹窗
	Priority      int      `json:"priority,omitempty"`       // 优先级
	ExpiresAt     string   `json:"expires_at,omitempty"`     // 过期时间
	ScheduledTime string   `json:"scheduled_time,omitempty"` // 定时发布时间
	Remark        string   `json:"remark,omitempty"`         // 备注
	Status        int      `json:"status,omitempty"`         // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
}

// 消息通知详情接口响应参数
type NoticeDetailRes struct {
	*Notice
}

// 消息通知分页列表接口请求参数
type NoticePageReq struct {
	Paging
	Title       string   `json:"title,omitempty"`        // 标题
	Content     string   `json:"content,omitempty"`      // 内容
	Category    string   `json:"category,omitempty"`     // 分类[service:服务消息, activity:活动消息, safety:安全消息, maint:维护消息, product:产品消息, fault:故障消息]
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
	PublishTime []string `json:"publish_time,omitempty"` // 发布时间
}

// 消息通知分页列表接口响应参数
type NoticePageRes struct {
	Items  []*Notice `json:"items"`
	Paging *Paging   `json:"paging"`
}

// 消息通知列表接口请求参数
type NoticeListReq struct {
	Title string `json:"title,omitempty"` // 标题
}

// 消息通知列表接口响应参数
type NoticeListRes struct {
	Items []*Notice `json:"items"`
}

// 消息通知批量操作接口请求参数
type NoticeBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

type Notice struct {
	Id            string   `json:"id,omitempty"`             // ID
	Title         string   `json:"title,omitempty"`          // 标题
	Content       string   `json:"content,omitempty"`        // 内容
	Category      string   `json:"category,omitempty"`       // 分类[service:服务消息, activity:活动消息, safety:安全消息, maint:维护消息, product:产品消息, fault:故障消息]
	Scope         int      `json:"scope,omitempty"`          // 通知范围[1:全部, 2:全部用户, 3:全部代理商, 4:指定用户, 5:指定代理商, 6:指定用户和代理商]
	Users         []int    `json:"users,omitempty"`          // 通知用户
	Resellers     []int    `json:"resellers,omitempty"`      // 通知代理商
	Channels      []string `json:"channels,omitempty"`       // 发送渠道[web:站内信, email:邮件]
	IsPopup       bool     `json:"is_popup,omitempty"`       // 是否弹窗
	Priority      int      `json:"priority,omitempty"`       // 优先级
	ExpiresAt     string   `json:"expires_at,omitempty"`     // 过期时间
	ScheduledTime string   `json:"scheduled_time,omitempty"` // 定时发布时间
	Remark        string   `json:"remark,omitempty"`         // 备注
	Status        int      `json:"status,omitempty"`         // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
	Variables     []string `json:"variables,omitempty"`      // 变量
	Publisher     int      `json:"publisher,omitempty"`      // 发布人
	PublishTime   string   `json:"publish_time,omitempty"`   // 发布时间
	Rid           int      `json:"rid,omitempty"`            // 代理商ID
	Creator       string   `json:"creator,omitempty"`        // 创建人
	Updater       string   `json:"updater,omitempty"`        // 更新人
	CreatedAt     string   `json:"created_at,omitempty"`     // 创建时间
	UpdatedAt     string   `json:"updated_at,omitempty"`     // 更新时间
}
