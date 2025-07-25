package do

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

const (
	NOTICE_TEMPLATE_COLLECTION = "notice_template"
)

type NoticeTemplate struct {
	gmeta.Meta `collection:"notice_template" bson:"-"`
	Name       string   `bson:"name,omitempty"`       // 名称
	Scenes     []string `bson:"scenes,omitempty"`     // 使用场景[code:验证码, login:登录通知, register:注册通知, forget_password:找回密码, change_password:修改密码, change_email:修改邮箱, change_phone:修改手机号, quota_recharge:充值额度, quota_warning:额度不足提醒, quota_exhaustion:额度耗尽通知, quota_expire_warning:额度过期提醒, quota_expire:额度过期通知, notice:消息通知]
	Title      string   `bson:"title,omitempty"`      // 标题
	Content    string   `bson:"content,omitempty"`    // 内容
	Channels   []string `bson:"channels,omitempty"`   // 适用渠道[web:站内信, email:邮件]
	IsPopup    bool     `bson:"is_popup"`             // 是否弹窗
	IsPublic   bool     `bson:"is_public"`            // 是否公开
	Remark     string   `bson:"remark"`               // 备注
	Status     int      `bson:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Variables  []string `bson:"variables"`            // 变量
	UserId     int      `bson:"user_id,omitempty"`    // 用户ID
	Rid        int      `bson:"rid,omitempty"`        // 代理商ID
	Creator    string   `bson:"creator,omitempty"`    // 创建人
	Updater    string   `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64    `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64    `bson:"updated_at,omitempty"` // 更新时间
}
