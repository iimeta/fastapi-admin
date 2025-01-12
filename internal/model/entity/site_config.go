package entity

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

type SiteConfig struct {
	gmeta.Meta          `role:"*" bson:"-"`
	Id                  string                `bson:"_id,omitempty"`                   // ID
	Domain              string                `bson:"domain,omitempty"`                // 域名
	Title               string                `bson:"title,omitempty"`                 // 标题
	Logo                string                `bson:"logo,omitempty"`                  // Logo
	Favicon             string                `bson:"favicon,omitempty"`               // Favicon
	Avatar              string                `bson:"avatar,omitempty"`                // 头像
	BgImg               string                `bson:"bg_img,omitempty"`                // 背景图
	Copyright           string                `bson:"copyright,omitempty"`             // 版权信息
	JumpUrl             string                `bson:"jump_url,omitempty"`              // 跳转URL
	Keywords            string                `bson:"keywords,omitempty"`              // 关键词
	Description         string                `bson:"description,omitempty"`           // 描述
	IcpBeian            string                `bson:"icp_beian,omitempty"`             // ICP备案
	GaBeian             string                `bson:"ga_beian,omitempty"`              // 公安备案
	RegisterTips        string                `bson:"register_tips,omitempty"`         // 注册提示
	GrantQuota          int                   `bson:"grant_quota,omitempty"`           // 注册授予额度
	QuotaExpiresAt      int                   `bson:"quota_expires_at,omitempty"`      // 注册授予额度过期时间, 单位: 分钟
	SupportEmailSuffix  []string              `bson:"support_email_suffix,omitempty"`  // 注册支持的邮箱后缀
	Host                string                `bson:"host,omitempty"`                  // 发信服务器
	Port                int                   `bson:"port,omitempty"`                  // 发信端口号
	UserName            string                `bson:"user_name,omitempty"`             // 发信账号
	Password            string                `bson:"password,omitempty"`              // 发信密码
	FromName            string                `bson:"from_name,omitempty"`             // 发信人名称
	Carousel1Title      string                `bson:"carousel1_title,omitempty"`       // 轮播图1标题
	Carousels1          []common.Carousel     `bson:"carousels1,omitempty"`            // 轮播图1
	Carousel2Title      string                `bson:"carousel2_title,omitempty"`       // 轮播图2标题
	Carousels2          []common.Carousel     `bson:"carousels2,omitempty"`            // 轮播图2
	AnnouncementTitle   string                `bson:"announcement_title,omitempty"`    // 公告标题
	AnnouncementMoreUrl string                `bson:"announcement_more_url,omitempty"` // 公告更多URL
	Announcements       []common.Announcement `bson:"announcements,omitempty"`         // 公告
	DocumentTitle       string                `bson:"document_title,omitempty"`        // 文档标题
	DocumentMoreUrl     string                `bson:"document_more_url,omitempty"`     // 文档更多URL
	Documents           []common.Document     `bson:"documents,omitempty"`             // 文档
	Remark              string                `bson:"remark,omitempty"`                // 备注
	Status              int                   `bson:"status,omitempty"`                // 状态[1:正常, 2:禁用, -1:删除]
	UserId              int                   `bson:"user_id,omitempty"`               // 用户ID
	Creator             string                `bson:"creator,omitempty"`               // 创建人
	Updater             string                `bson:"updater,omitempty"`               // 更新人
	CreatedAt           int64                 `bson:"created_at,omitempty"`            // 创建时间
	UpdatedAt           int64                 `bson:"updated_at,omitempty"`            // 更新时间
}
