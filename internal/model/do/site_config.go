package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	SITE_CONFIG_COLLECTION = "site_config"
)

type SiteConfig struct {
	gmeta.Meta         `collection:"site_config" bson:"-"`
	Domain             string   `bson:"domain,omitempty"`     // 域名
	Title              string   `bson:"title,omitempty"`      // 标题
	Logo               string   `bson:"logo,omitempty"`       // Logo
	Favicon            string   `bson:"favicon,omitempty"`    // Favicon
	Avatar             string   `bson:"avatar"`               // 头像
	BgImg              string   `bson:"bg_img"`               // 背景图
	Copyright          string   `bson:"copyright"`            // 版权信息
	JumpUrl            string   `bson:"jump_url"`             // 跳转URL
	Keywords           string   `bson:"keywords"`             // 关键词
	Description        string   `bson:"description"`          // 描述
	IcpBeian           string   `bson:"icp_beian"`            // ICP备案
	GaBeian            string   `bson:"ga_beian"`             // 公安备案
	RegisterTips       string   `bson:"register_tips"`        // 注册提示
	GrantQuota         int      `bson:"grant_quota"`          // 注册授予额度
	QuotaExpiresAt     int      `bson:"quota_expires_at"`     // 注册授予额度过期时间, 单位: 分钟
	SupportEmailSuffix []string `bson:"support_email_suffix"` // 注册支持的邮箱后缀
	Host               string   `bson:"host"`                 // 发信服务器
	Port               int      `bson:"port"`                 // 发信端口号
	UserName           string   `bson:"user_name"`            // 发信账号
	Password           string   `bson:"password"`             // 发信密码
	FromName           string   `bson:"from_name"`            // 发信人名称
	Remark             string   `bson:"remark"`               // 备注
	Status             int      `bson:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	UserId             int      `bson:"user_id,omitempty"`    // 用户ID
	Creator            string   `bson:"creator,omitempty"`    // 创建人
	Updater            string   `bson:"updater,omitempty"`    // 更新人
	CreatedAt          int64    `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt          int64    `bson:"updated_at,omitempty"` // 更新时间
}
