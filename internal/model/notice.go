package model

import "github.com/iimeta/fastapi-admin/internal/model/common"

// 新建通知公告接口请求参数
type NoticeCreateReq struct {
	Title         string `json:"title,omitempty"`          // 标题
	Content       string `json:"content,omitempty"`        // 内容
	Category      int    `json:"category,omitempty"`       // 分类[1:系统公告, 2:活动通知, 3:维护通知]
	Scope         int    `json:"scope,omitempty"`          // 通知范围[1:全部, 2:指定用户, 3:指定代理商, 4:指定用户和代理商]
	Users         []int  `json:"users,omitempty"`          // 通知用户
	Resellers     []int  `json:"resellers,omitempty"`      // 通知代理商
	Methods       []int  `json:"methods,omitempty"`        // 通知方式[1:站内信, 2:邮件]
	Priority      int    `json:"priority,omitempty"`       // 优先级
	ExpiresAt     string `json:"expires_at,omitempty"`     // 过期时间
	ScheduledTime string `json:"scheduled_time,omitempty"` // 定时发布时间
	Remark        string `json:"remark,omitempty"`         // 备注
	Status        int    `json:"status,omitempty"`         // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
}

// 更新通知公告接口请求参数
type NoticeUpdateReq struct {
	Id            string `json:"id" v:"required"`          // ID
	Title         string `json:"title,omitempty"`          // 标题
	Content       string `json:"content,omitempty"`        // 内容
	Category      int    `json:"category,omitempty"`       // 分类[1:系统公告, 2:活动通知, 3:维护通知]
	Scope         int    `json:"scope,omitempty"`          // 通知范围[1:全部, 2:指定用户, 3:指定代理商, 4:指定用户和代理商]
	Users         []int  `json:"users,omitempty"`          // 通知用户
	Resellers     []int  `json:"resellers,omitempty"`      // 通知代理商
	Methods       []int  `json:"methods,omitempty"`        // 通知方式[1:站内信, 2:邮件]
	Priority      int    `json:"priority,omitempty"`       // 优先级
	ExpiresAt     string `json:"expires_at,omitempty"`     // 过期时间
	ScheduledTime string `json:"scheduled_time,omitempty"` // 定时发布时间
	Remark        string `json:"remark,omitempty"`         // 备注
	Status        int    `json:"status,omitempty"`         // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
}

// 通知公告详情接口响应参数
type NoticeDetailRes struct {
	*Notice
}

// 通知公告分页列表接口请求参数
type NoticePageReq struct {
	Paging
	Title     string   `json:"title,omitempty"`      // 标题
	Content   string   `json:"content,omitempty"`    // 内容
	Remark    string   `json:"remark,omitempty"`     // 备注
	Status    int      `json:"status,omitempty"`     // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
	UpdatedAt []string `json:"updated_at,omitempty"` // 更新时间
}

// 通知公告分页列表接口响应参数
type NoticePageRes struct {
	Items  []*Notice `json:"items"`
	Paging *Paging   `json:"paging"`
}

// 通知公告列表接口请求参数
type NoticeListReq struct {
	Title string `json:"title,omitempty"` // 标题
}

// 通知公告列表接口响应参数
type NoticeListRes struct {
	Items []*Notice `json:"items"`
}

// 通知公告批量操作接口请求参数
type NoticeBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

type Notice struct {
	Id            string        `json:"id,omitempty"`             // ID
	Title         string        `json:"title,omitempty"`          // 标题
	Content       string        `json:"content,omitempty"`        // 内容
	Category      int           `json:"category,omitempty"`       // 分类[1:系统公告, 2:活动通知, 3:维护通知]
	Scope         int           `json:"scope,omitempty"`          // 通知范围[1:全部, 2:指定用户, 3:指定代理商, 4:指定用户和代理商]
	Users         []int         `json:"users,omitempty"`          // 通知用户
	Resellers     []int         `json:"resellers,omitempty"`      // 通知代理商
	Methods       []int         `json:"methods,omitempty"`        // 通知方式[1:站内信, 2:邮件]
	Priority      int           `json:"priority,omitempty"`       // 优先级
	ExpiresAt     string        `json:"expires_at,omitempty"`     // 过期时间
	ScheduledTime string        `json:"scheduled_time,omitempty"` // 定时发布时间
	Remark        string        `json:"remark,omitempty"`         // 备注
	Status        int           `json:"status,omitempty"`         // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
	Reads         []common.Read `json:"reads,omitempty"`          // 已读
	Rid           int           `json:"rid,omitempty"`            // 代理商ID
	Creator       string        `json:"creator,omitempty"`        // 创建人
	Updater       string        `json:"updater,omitempty"`        // 更新人
	CreatedAt     string        `json:"created_at,omitempty"`     // 创建时间
	UpdatedAt     string        `json:"updated_at,omitempty"`     // 更新时间
}
