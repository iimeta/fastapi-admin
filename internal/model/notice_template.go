package model

// 新建通知公告模板接口请求参数
type NoticeTemplateCreateReq struct {
	Name     string `json:"name,omitempty"`      // 名称
	Action   string `json:"action,omitempty"`    // 动作
	Content  string `json:"content,omitempty"`   // 内容
	Category int    `json:"category,omitempty"`  // 分类[1:系统公告, 2:活动通知, 3:维护通知]
	IsPublic bool   `json:"is_public,omitempty"` // 是否公开
	Remark   string `json:"remark,omitempty"`    // 备注
	Status   int    `json:"status,omitempty"`    // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新通知公告模板接口请求参数
type NoticeTemplateUpdateReq struct {
	Id       string `json:"id" v:"required"`     // ID
	Name     string `json:"name,omitempty"`      // 名称
	Action   string `json:"action,omitempty"`    // 动作
	Content  string `json:"content,omitempty"`   // 内容
	Category int    `json:"category,omitempty"`  // 分类[1:系统公告, 2:活动通知, 3:维护通知]
	IsPublic bool   `json:"is_public,omitempty"` // 是否公开
	Remark   string `json:"remark,omitempty"`    // 备注
	Status   int    `json:"status,omitempty"`    // 状态[1:正常, 2:禁用, -1:删除]
}

// 通知公告模板详情接口响应参数
type NoticeTemplateDetailRes struct {
	*NoticeTemplate
}

// 通知公告模板分页列表接口请求参数
type NoticeTemplatePageReq struct {
	Paging
	Name        string   `json:"name,omitempty"`         // 名称
	Content     string   `json:"content,omitempty"`      // 内容
	Category    int      `json:"category,omitempty"`     // 分类[1:系统公告, 2:活动通知, 3:维护通知]
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
	PublishTime []string `json:"publish_time,omitempty"` // 发布时间
}

// 通知公告模板分页列表接口响应参数
type NoticeTemplatePageRes struct {
	Items  []*NoticeTemplate `json:"items"`
	Paging *Paging           `json:"paging"`
}

// 通知公告模板列表接口请求参数
type NoticeTemplateListReq struct {
	Title string `json:"title,omitempty"` // 标题
}

// 通知公告模板列表接口响应参数
type NoticeTemplateListRes struct {
	Items []*NoticeTemplate `json:"items"`
}

// 通知公告模板批量操作接口请求参数
type NoticeTemplateBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

type NoticeTemplate struct {
	Id        string `json:"id,omitempty"`         // ID
	Name      string `json:"name,omitempty"`       // 名称
	Action    string `json:"action,omitempty"`     // 动作
	Content   string `json:"content,omitempty"`    // 内容
	Category  int    `json:"category,omitempty"`   // 分类[1:系统公告, 2:活动通知, 3:维护通知]
	IsPublic  bool   `json:"is_public,omitempty"`  // 是否公开
	Remark    string `json:"remark,omitempty"`     // 备注
	Status    int    `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	UserId    int    `json:"user_id,omitempty"`    // 用户ID
	Rid       int    `json:"rid,omitempty"`        // 代理商ID
	Creator   string `json:"creator,omitempty"`    // 创建人
	Updater   string `json:"updater,omitempty"`    // 更新人
	CreatedAt string `json:"created_at,omitempty"` // 创建时间
	UpdatedAt string `json:"updated_at,omitempty"` // 更新时间
}
