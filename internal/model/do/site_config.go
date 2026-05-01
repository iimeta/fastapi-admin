package do

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/v2/internal/model/common"
)

type SiteConfig struct {
	gmeta.Meta                        `collection:"site_config" bson:"-"`
	Domain                            string                `bson:"domain,omitempty"`                      // 域名
	Title                             string                `bson:"title,omitempty"`                       // 标题
	Logo                              string                `bson:"logo,omitempty"`                        // Logo
	Favicon                           string                `bson:"favicon,omitempty"`                     // Favicon
	Avatar                            string                `bson:"avatar"`                                // 头像
	BgImg                             string                `bson:"bg_img"`                                // 背景图
	Copyright                         string                `bson:"copyright"`                             // 版权信息
	JumpUrl                           string                `bson:"jump_url"`                              // 跳转URL
	Keywords                          string                `bson:"keywords"`                              // 关键词
	Description                       string                `bson:"description"`                           // 描述
	IcpBeian                          string                `bson:"icp_beian"`                             // ICP备案
	GaBeian                           string                `bson:"ga_beian"`                              // 公安备案
	RegisterTips                      string                `bson:"register_tips"`                         // 注册提示
	GrantQuota                        int                   `bson:"grant_quota"`                           // 注册授予额度
	InviteEnabled                     bool                  `bson:"invite_enabled"`                        // 是否开启邀请注册
	InviteCodeRequired                bool                  `bson:"invite_code_required"`                  // 注册时邀请码是否必填
	InviteRewardQuota                 int                   `bson:"invite_reward_quota"`                   // 邀请人收益额度
	InviteeGrantQuota                 int                   `bson:"invitee_grant_quota"`                   // 被邀请人额外赠送额度
	InviteMinApplyQuota               int                   `bson:"invite_min_apply_quota"`                // 最低申请入账额度
	InviteDailyLimit                  int                   `bson:"invite_daily_limit"`                    // 单日邀请收益次数上限
	InviteTotalLimit                  int                   `bson:"invite_total_limit"`                    // 累计邀请收益次数上限
	InviteIpDailyLimit                int                   `bson:"invite_ip_daily_limit"`                 // 同IP每日邀请注册上限
	InviteIpTotalLimit                int                   `bson:"invite_ip_total_limit"`                 // 同IP累计邀请注册上限
	InviteIpPerInviterLimit           int                   `bson:"invite_ip_per_inviter_limit"`           // 同IP+同邀请人注册上限
	InviteIpLimitAction               string                `bson:"invite_ip_limit_action"`                // IP限制触发动作[block:拒绝注册, silent:允许注册但不发放额度]
	InviteRuleText                    string                `bson:"invite_rule_text"`                      // 邀请规则说明
	InviteInvalidCodeAction           string                `bson:"invite_invalid_code_action"`            // 无效邀请码处理方式
	InviteRechargeRebateEnabled       bool                  `bson:"invite_recharge_rebate_enabled"`        // 是否开启邀请充值返利
	InviteRechargeRebateFirstEnabled  bool                  `bson:"invite_recharge_rebate_first_enabled"`  // 首次充值返利是否开启
	InviteRechargeRebateFirstType     string                `bson:"invite_recharge_rebate_first_type"`     // 首次充值返利类型[percent:百分比, fixed:固定额度]
	InviteRechargeRebateFirstRate     float64               `bson:"invite_recharge_rebate_first_rate"`     // 首次充值返利比例
	InviteRechargeRebateFirstQuota    int                   `bson:"invite_recharge_rebate_first_quota"`    // 首次充值固定返利额度，按系统内部整数额度存储
	InviteRechargeRebateSecondEnabled bool                  `bson:"invite_recharge_rebate_second_enabled"` // 后续充值返利是否开启
	InviteRechargeRebateSecondType    string                `bson:"invite_recharge_rebate_second_type"`    // 后续充值返利类型[percent:百分比, fixed:固定额度]
	InviteRechargeRebateSecondRate    float64               `bson:"invite_recharge_rebate_second_rate"`    // 后续充值返利比例
	InviteRechargeRebateSecondQuota   int                   `bson:"invite_recharge_rebate_second_quota"`   // 后续充值固定返利额度，按系统内部整数额度存储
	QuotaExpiresAt                    int                   `bson:"quota_expires_at"`                      // 注册授予额度过期时间, 单位: 分钟
	SupportEmailSuffix                []string              `bson:"support_email_suffix"`                  // 注册支持的邮箱后缀
	DefaultLanguage                   string                `bson:"default_language"`                      // 默认语言[zh-CN:简体中文, zh-TW:繁體中文, en-US:English]
	CurrencySymbol                    string                `bson:"currency_symbol"`                       // 货币符号
	RegisterWelcome                   string                `bson:"register_welcome"`                      // 注册欢迎语
	Host                              string                `bson:"host"`                                  // 发信服务器
	Port                              int                   `bson:"port"`                                  // 发信端口号
	UserName                          string                `bson:"user_name"`                             // 发信账号
	Password                          string                `bson:"password"`                              // 发信密码
	FromName                          string                `bson:"from_name"`                             // 发信人名称
	Carousel1Title                    string                `bson:"carousel1_title"`                       // 轮播图1标题
	Carousels1                        []common.Carousel     `bson:"carousels1"`                            // 轮播图1
	Carousel2Title                    string                `bson:"carousel2_title"`                       // 轮播图2标题
	Carousels2                        []common.Carousel     `bson:"carousels2"`                            // 轮播图2
	AnnouncementTitle                 string                `bson:"announcement_title"`                    // 公告标题
	AnnouncementMoreUrl               string                `bson:"announcement_more_url"`                 // 公告更多URL
	Announcements                     []common.Announcement `bson:"announcements"`                         // 公告
	DocumentTitle                     string                `bson:"document_title"`                        // 文档标题
	DocumentMoreUrl                   string                `bson:"document_more_url"`                     // 文档更多URL
	Documents                         []common.Document     `bson:"documents"`                             // 文档
	RechargeTips                      string                `bson:"recharge_tips"`                         // 充值提示
	Remark                            string                `bson:"remark"`                                // 备注
	Status                            int                   `bson:"status,omitempty"`                      // 状态[1:正常, 2:禁用, -1:删除]
	UserId                            int                   `bson:"user_id,omitempty"`                     // 用户ID
	Rid                               int                   `bson:"rid,omitempty"`                         // 代理商ID
	Creator                           string                `bson:"creator,omitempty"`                     // 创建人
	Updater                           string                `bson:"updater,omitempty"`                     // 更新人
	CreatedAt                         int64                 `bson:"created_at,omitempty"`                  // 创建时间
	UpdatedAt                         int64                 `bson:"updated_at,omitempty"`                  // 更新时间
}
