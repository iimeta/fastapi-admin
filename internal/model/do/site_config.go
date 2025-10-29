package do

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

type SiteConfig struct {
	gmeta.Meta          `collection:"site_config" bson:"-"`
	Domain              string                `bson:"domain,omitempty"`      // 域名
	Title               string                `bson:"title,omitempty"`       // 标题
	Logo                string                `bson:"logo,omitempty"`        // Logo
	Favicon             string                `bson:"favicon,omitempty"`     // Favicon
	Avatar              string                `bson:"avatar"`                // 头像
	BgImg               string                `bson:"bg_img"`                // 背景图
	Copyright           string                `bson:"copyright"`             // 版权信息
	JumpUrl             string                `bson:"jump_url"`              // 跳转URL
	Keywords            string                `bson:"keywords"`              // 关键词
	Description         string                `bson:"description"`           // 描述
	IcpBeian            string                `bson:"icp_beian"`             // ICP备案
	GaBeian             string                `bson:"ga_beian"`              // 公安备案
	RegisterTips        string                `bson:"register_tips"`         // 注册提示
	GrantQuota          int                   `bson:"grant_quota"`           // 注册授予额度
	QuotaExpiresAt      int                   `bson:"quota_expires_at"`      // 注册授予额度过期时间, 单位: 分钟
	SupportEmailSuffix  []string              `bson:"support_email_suffix"`  // 注册支持的邮箱后缀
	DefaultLanguage     string                `bson:"default_language"`      // 默认语言[zh-CN:简体中文, zh-TW:繁體中文, en-US:English]
	CurrencySymbol      string                `bson:"currency_symbol"`       // 货币符号
	RegisterWelcome     string                `bson:"register_welcome"`      // 注册欢迎语
	Host                string                `bson:"host"`                  // 发信服务器
	Port                int                   `bson:"port"`                  // 发信端口号
	UserName            string                `bson:"user_name"`             // 发信账号
	Password            string                `bson:"password"`              // 发信密码
	FromName            string                `bson:"from_name"`             // 发信人名称
	Carousel1Title      string                `bson:"carousel1_title"`       // 轮播图1标题
	Carousels1          []common.Carousel     `bson:"carousels1"`            // 轮播图1
	Carousel2Title      string                `bson:"carousel2_title"`       // 轮播图2标题
	Carousels2          []common.Carousel     `bson:"carousels2"`            // 轮播图2
	AnnouncementTitle   string                `bson:"announcement_title"`    // 公告标题
	AnnouncementMoreUrl string                `bson:"announcement_more_url"` // 公告更多URL
	Announcements       []common.Announcement `bson:"announcements"`         // 公告
	DocumentTitle       string                `bson:"document_title"`        // 文档标题
	DocumentMoreUrl     string                `bson:"document_more_url"`     // 文档更多URL
	Documents           []common.Document     `bson:"documents"`             // 文档
	RechargeTips        string                `bson:"recharge_tips"`         // 充值提示
	Remark              string                `bson:"remark"`                // 备注
	Status              int                   `bson:"status,omitempty"`      // 状态[1:正常, 2:禁用, -1:删除]
	UserId              int                   `bson:"user_id,omitempty"`     // 用户ID
	Rid                 int                   `bson:"rid,omitempty"`         // 代理商ID
	Creator             string                `bson:"creator,omitempty"`     // 创建人
	Updater             string                `bson:"updater,omitempty"`     // 更新人
	CreatedAt           int64                 `bson:"created_at,omitempty"`  // 创建时间
	UpdatedAt           int64                 `bson:"updated_at,omitempty"`  // 更新时间
}
