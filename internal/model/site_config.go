package model

import "github.com/iimeta/fastapi-admin/v2/internal/model/common"

// 新建站点配置接口请求参数
type SiteConfigCreateReq struct {
	Domain              string                `json:"domain,omitempty"`                // 域名
	Title               string                `json:"title,omitempty"`                 // 标题
	Logo                string                `json:"logo,omitempty"`                  // Logo
	Favicon             string                `json:"favicon,omitempty"`               // Favicon
	Avatar              string                `json:"avatar,omitempty"`                // 头像
	BgImg               string                `json:"bg_img,omitempty"`                // 背景图
	Copyright           string                `json:"copyright,omitempty"`             // 版权信息
	JumpUrl             string                `json:"jump_url,omitempty"`              // 跳转URL
	Keywords            string                `json:"keywords,omitempty"`              // 关键词
	Description         string                `json:"description,omitempty"`           // 描述
	IcpBeian            string                `json:"icp_beian,omitempty"`             // ICP备案
	GaBeian             string                `json:"ga_beian,omitempty"`              // 公安备案
	RegisterTips        string                `json:"register_tips,omitempty"`         // 注册提示
	GrantQuota          float64               `json:"grant_quota,omitempty"`           // 注册授予额度
	QuotaExpiresAt      int                   `json:"quota_expires_at,omitempty"`      // 注册授予额度过期时间, 单位: 分钟
	SupportEmailSuffix  []string              `json:"support_email_suffix,omitempty"`  // 注册支持的邮箱后缀
	RegisterWelcome     string                `json:"register_welcome,omitempty"`      // 注册欢迎语
	DefaultLanguage     string                `json:"default_language,omitempty"`      // 默认语言[zh-CN:简体中文, zh-TW:繁體中文, en-US:English]
	CurrencySymbol      string                `json:"currency_symbol,omitempty"`       // 货币符号
	Host                string                `json:"host,omitempty"`                  // 发信服务器
	Port                int                   `json:"port,omitempty"`                  // 发信端口号
	UserName            string                `json:"user_name,omitempty"`             // 发信账号
	Password            string                `json:"password,omitempty"`              // 发信密码
	FromName            string                `json:"from_name,omitempty"`             // 发信人名称
	Carousel1Title      string                `json:"carousel1_title,omitempty"`       // 轮播图1标题
	Carousels1          []common.Carousel     `json:"carousels1,omitempty"`            // 轮播图1
	Carousel2Title      string                `json:"carousel2_title,omitempty"`       // 轮播图2标题
	Carousels2          []common.Carousel     `json:"carousels2,omitempty"`            // 轮播图2
	AnnouncementTitle   string                `json:"announcement_title,omitempty"`    // 公告标题
	AnnouncementMoreUrl string                `json:"announcement_more_url,omitempty"` // 公告更多URL
	Announcements       []common.Announcement `json:"announcements,omitempty"`         // 公告
	DocumentTitle       string                `json:"document_title,omitempty"`        // 文档标题
	DocumentMoreUrl     string                `json:"document_more_url,omitempty"`     // 文档更多URL
	Documents           []common.Document     `json:"documents,omitempty"`             // 文档
	RechargeTips        string                `json:"recharge_tips,omitempty"`         // 充值提示
	Remark              string                `json:"remark,omitempty"`                // 备注
	Status              int                   `json:"status,omitempty" d:"1"`          // 状态[1:正常, 2:禁用, -1:删除]
}

// 更新站点配置接口请求参数
type SiteConfigUpdateReq struct {
	Id                  string                `json:"id,omitempty"`                    // ID
	Domain              string                `json:"domain,omitempty"`                // 域名
	Title               string                `json:"title,omitempty"`                 // 标题
	Logo                string                `json:"logo,omitempty"`                  // Logo
	Favicon             string                `json:"favicon,omitempty"`               // Favicon
	Avatar              string                `json:"avatar,omitempty"`                // 头像
	BgImg               string                `json:"bg_img,omitempty"`                // 背景图
	Copyright           string                `json:"copyright,omitempty"`             // 版权信息
	JumpUrl             string                `json:"jump_url,omitempty"`              // 跳转URL
	Keywords            string                `json:"keywords,omitempty"`              // 关键词
	Description         string                `json:"description,omitempty"`           // 描述
	IcpBeian            string                `json:"icp_beian,omitempty"`             // ICP备案
	GaBeian             string                `json:"ga_beian,omitempty"`              // 公安备案
	RegisterTips        string                `json:"register_tips,omitempty"`         // 注册提示
	GrantQuota          float64               `json:"grant_quota,omitempty"`           // 注册授予额度
	QuotaExpiresAt      int                   `json:"quota_expires_at,omitempty"`      // 注册授予额度过期时间, 单位: 分钟
	SupportEmailSuffix  []string              `json:"support_email_suffix,omitempty"`  // 注册支持的邮箱后缀
	RegisterWelcome     string                `json:"register_welcome,omitempty"`      // 注册欢迎语
	DefaultLanguage     string                `json:"default_language,omitempty"`      // 默认语言[zh-CN:简体中文, zh-TW:繁體中文, en-US:English]
	CurrencySymbol      string                `json:"currency_symbol,omitempty"`       // 货币符号
	Host                string                `json:"host,omitempty"`                  // 发信服务器
	Port                int                   `json:"port,omitempty"`                  // 发信端口号
	UserName            string                `json:"user_name,omitempty"`             // 发信账号
	Password            string                `json:"password,omitempty"`              // 发信密码
	FromName            string                `json:"from_name,omitempty"`             // 发信人名称
	Carousel1Title      string                `json:"carousel1_title,omitempty"`       // 轮播图1标题
	Carousels1          []common.Carousel     `json:"carousels1,omitempty"`            // 轮播图1
	Carousel2Title      string                `json:"carousel2_title,omitempty"`       // 轮播图2标题
	Carousels2          []common.Carousel     `json:"carousels2,omitempty"`            // 轮播图2
	AnnouncementTitle   string                `json:"announcement_title,omitempty"`    // 公告标题
	AnnouncementMoreUrl string                `json:"announcement_more_url,omitempty"` // 公告更多URL
	Announcements       []common.Announcement `json:"announcements,omitempty"`         // 公告
	DocumentTitle       string                `json:"document_title,omitempty"`        // 文档标题
	DocumentMoreUrl     string                `json:"document_more_url,omitempty"`     // 文档更多URL
	Documents           []common.Document     `json:"documents,omitempty"`             // 文档
	RechargeTips        string                `json:"recharge_tips,omitempty"`         // 充值提示
	Remark              string                `json:"remark,omitempty"`                // 备注
	Status              int                   `json:"status,omitempty" d:"1"`          // 状态[1:正常, 2:禁用, -1:删除]
}

// 更改站点配置状态接口请求参数
type SiteConfigChangeStatusReq struct {
	Id     string `json:"id,omitempty"`           // ID
	Status int    `json:"status,omitempty" d:"1"` // 状态[1:正常, 2:禁用, -1:删除]
}

// 站点配置详情接口请求参数
type SiteConfigDetailReq struct {
	Id     string `json:"id,omitempty"`     // ID
	Domain string `json:"domain,omitempty"` // 域名
}

// 站点配置详情接口响应参数
type SiteConfigDetailRes struct {
	*SiteConfig
}

// 站点配置分页列表接口请求参数
type SiteConfigPageReq struct {
	Paging
	UserId       int    `json:"user_id,omitempty"`       // 用户ID
	RegisterTips string `json:"register_tips,omitempty"` // 注册提示
	Logo         string `json:"logo,omitempty"`          // Logo
	Domain       string `json:"domain,omitempty"`        // 域名
	Title        string `json:"title,omitempty"`         // 标题
	Status       int    `json:"status,omitempty"`        // 状态[1:正常, 2:禁用, -1:删除]
}

// 站点配置分页列表接口响应参数
type SiteConfigPageRes struct {
	Items  []*SiteConfig `json:"items"`
	Paging *Paging       `json:"paging"`
}

// 站点配置批量操作接口请求参数
type SiteConfigBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

type SiteConfig struct {
	Id                  string                `json:"id,omitempty"`                    // ID
	Domain              string                `json:"domain,omitempty"`                // 域名
	Title               string                `json:"title,omitempty"`                 // 标题
	Logo                string                `json:"logo,omitempty"`                  // Logo
	Favicon             string                `json:"favicon,omitempty"`               // Favicon
	Avatar              string                `json:"avatar,omitempty"`                // 头像
	BgImg               string                `json:"bg_img,omitempty"`                // 背景图
	Copyright           string                `json:"copyright,omitempty"`             // 版权信息
	JumpUrl             string                `json:"jump_url,omitempty"`              // 跳转URL
	Keywords            string                `json:"keywords,omitempty"`              // 关键词
	Description         string                `json:"description,omitempty"`           // 描述
	IcpBeian            string                `json:"icp_beian,omitempty"`             // ICP备案
	GaBeian             string                `json:"ga_beian,omitempty"`              // 公安备案
	RegisterTips        string                `json:"register_tips,omitempty"`         // 注册提示
	GrantQuota          float64               `json:"grant_quota,omitempty"`           // 注册授予额度
	QuotaExpiresAt      int                   `json:"quota_expires_at,omitempty"`      // 注册授予额度过期时间, 单位: 分钟
	SupportEmailSuffix  []string              `json:"support_email_suffix,omitempty"`  // 注册支持的邮箱后缀
	RegisterWelcome     string                `json:"register_welcome,omitempty"`      // 注册欢迎语
	DefaultLanguage     string                `json:"default_language,omitempty"`      // 默认语言[zh-CN:简体中文, zh-TW:繁體中文, en-US:English]
	CurrencySymbol      string                `json:"currency_symbol,omitempty"`       // 货币符号
	Host                string                `json:"host,omitempty"`                  // 发信服务器
	Port                int                   `json:"port,omitempty"`                  // 发信端口号
	UserName            string                `json:"user_name,omitempty"`             // 发信账号
	Password            string                `json:"password,omitempty"`              // 发信密码
	FromName            string                `json:"from_name,omitempty"`             // 发信人名称
	Carousel1Title      string                `json:"carousel1_title,omitempty"`       // 轮播图1标题
	Carousels1          []common.Carousel     `json:"carousels1,omitempty"`            // 轮播图1
	Carousel2Title      string                `json:"carousel2_title,omitempty"`       // 轮播图2标题
	Carousels2          []common.Carousel     `json:"carousels2,omitempty"`            // 轮播图2
	AnnouncementTitle   string                `json:"announcement_title,omitempty"`    // 公告标题
	AnnouncementMoreUrl string                `json:"announcement_more_url,omitempty"` // 公告更多URL
	Announcements       []common.Announcement `json:"announcements,omitempty"`         // 公告
	DocumentTitle       string                `json:"document_title,omitempty"`        // 文档标题
	DocumentMoreUrl     string                `json:"document_more_url,omitempty"`     // 文档更多URL
	Documents           []common.Document     `json:"documents,omitempty"`             // 文档
	RechargeTips        string                `json:"recharge_tips,omitempty"`         // 充值提示
	Remark              string                `json:"remark,omitempty"`                // 备注
	Status              int                   `json:"status,omitempty"`                // 状态[1:正常, 2:禁用, -1:删除]
	UserId              int                   `json:"user_id,omitempty"`               // 用户ID
	Creator             string                `json:"creator,omitempty"`               // 创建人
	Updater             string                `json:"updater,omitempty"`               // 更新人
	CreatedAt           string                `json:"created_at,omitempty"`            // 创建时间
	UpdatedAt           string                `json:"updated_at,omitempty"`            // 更新时间
}
