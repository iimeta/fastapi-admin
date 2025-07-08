package model

// 新建通知模板接口请求参数
type NoticeTemplateCreateReq struct {
	Name     string   `json:"name,omitempty"`         // 名称
	Scenes   []string `json:"scenes,omitempty"`       // 使用场景[code:验证码, login:登录通知, register:注册通知, forget_password:找回密码, change_password:修改密码, change_email:修改邮箱, quota_warning:额度不足提醒, quota_exhaustion:额度耗尽通知, quota_expire_warning:额度过期提醒, quota_expire:额度过期通知, notice:通知公告]
	Title    string   `json:"title,omitempty"`        // 标题
	Content  string   `json:"content,omitempty"`      // 内容
	Channels []string `json:"channels,omitempty"`     // 适用渠道[web:站内信, email:邮件]
	IsPopup  bool     `json:"is_popup,omitempty"`     // 是否弹窗
	IsPublic bool     `json:"is_public,omitempty"`    // 是否公开
	Remark   string   `json:"remark,omitempty"`       // 备注
	Status   int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新通知模板接口请求参数
type NoticeTemplateUpdateReq struct {
	Id       string   `json:"id" v:"required"`        // ID
	Name     string   `json:"name,omitempty"`         // 名称
	Scenes   []string `json:"scenes,omitempty"`       // 使用场景[code:验证码, login:登录通知, register:注册通知, forget_password:找回密码, change_password:修改密码, change_email:修改邮箱, quota_warning:额度不足提醒, quota_exhaustion:额度耗尽通知, quota_expire_warning:额度过期提醒, quota_expire:额度过期通知, notice:通知公告]
	Title    string   `json:"title,omitempty"`        // 标题
	Content  string   `json:"content,omitempty"`      // 内容
	Channels []string `json:"channels,omitempty"`     // 适用渠道[web:站内信, email:邮件]
	IsPopup  bool     `json:"is_popup,omitempty"`     // 是否弹窗
	IsPublic bool     `json:"is_public,omitempty"`    // 是否公开
	Remark   string   `json:"remark,omitempty"`       // 备注
	Status   int      `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改通知模板公开状态接口请求参数
type NoticeTemplateChangePublicReq struct {
	Id       string `json:"id" v:"required"`     // ID
	IsPublic bool   `json:"is_public,omitempty"` // 是否公开
}

// 更改通知模板状态接口请求参数
type NoticeTemplateChangeStatusReq struct {
	Id     string `json:"id,omitempty"`           // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 通知模板详情接口响应参数
type NoticeTemplateDetailRes struct {
	*NoticeTemplate
}

// 通知模板分页列表接口请求参数
type NoticeTemplatePageReq struct {
	Paging
	Name     string   `json:"name,omitempty"`     // 名称
	Scenes   []string `json:"scenes,omitempty"`   // 使用场景[code:验证码, login:登录通知, register:注册通知, forget_password:找回密码, change_password:修改密码, change_email:修改邮箱, quota_warning:额度不足提醒, quota_exhaustion:额度耗尽通知, quota_expire_warning:额度过期提醒, quota_expire:额度过期通知, notice:通知公告]
	Title    string   `json:"title,omitempty"`    // 标题
	Content  string   `json:"content,omitempty"`  // 内容
	Channels []string `json:"channels,omitempty"` // 适用渠道[web:站内信, email:邮件]
	Remark   string   `json:"remark,omitempty"`   // 备注
}

// 通知模板分页列表接口响应参数
type NoticeTemplatePageRes struct {
	Items  []*NoticeTemplate `json:"items"`
	Paging *Paging           `json:"paging"`
}

// 通知模板列表接口请求参数
type NoticeTemplateListReq struct {
	Title string `json:"title,omitempty"` // 标题
}

// 通知模板列表接口响应参数
type NoticeTemplateListRes struct {
	Items []*NoticeTemplate `json:"items"`
}

// 通知模板批量操作接口请求参数
type NoticeTemplateBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

type NoticeTemplate struct {
	Id        string   `json:"id,omitempty"`         // ID
	Name      string   `json:"name,omitempty"`       // 名称
	Scenes    []string `json:"scenes,omitempty"`     // 使用场景[code:验证码, login:登录通知, register:注册通知, forget_password:找回密码, change_password:修改密码, change_email:修改邮箱, quota_warning:额度不足提醒, quota_exhaustion:额度耗尽通知, quota_expire_warning:额度过期提醒, quota_expire:额度过期通知, notice:通知公告]
	Title     string   `json:"title,omitempty"`      // 标题
	Content   string   `json:"content,omitempty"`    // 内容
	Channels  []string `json:"channels,omitempty"`   // 适用渠道[web:站内信, email:邮件]
	IsPopup   bool     `json:"is_popup,omitempty"`   // 是否弹窗
	IsPublic  bool     `json:"is_public,omitempty"`  // 是否公开
	Remark    string   `json:"remark,omitempty"`     // 备注
	Status    int      `json:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Variables []string `json:"variables,omitempty"`  // 变量
	UserId    int      `json:"user_id,omitempty"`    // 用户ID
	Rid       int      `json:"rid,omitempty"`        // 代理商ID
	Creator   string   `json:"creator,omitempty"`    // 创建人
	Updater   string   `json:"updater,omitempty"`    // 更新人
	CreatedAt string   `json:"created_at,omitempty"` // 创建时间
	UpdatedAt string   `json:"updated_at,omitempty"` // 更新时间
}
