package do

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

const (
	NOTICE_COLLECTION = "notice"
)

type Notice struct {
	gmeta.Meta    `collection:"notice" bson:"-"`
	Title         string        `bson:"title,omitempty"`        // 标题
	Content       string        `bson:"content,omitempty"`      // 内容
	Category      string        `bson:"category,omitempty"`     // 分类[service:服务消息, activity:活动消息, safety:安全消息, maint:维护消息, product:产品消息, fault:故障消息]
	Scope         int           `bson:"scope,omitempty"`        // 通知范围[1:全部, 2:全部用户, 3:全部代理商, 4:指定用户, 5:指定代理商, 6:指定用户和代理商]
	Users         []int         `bson:"users"`                  // 通知用户
	Resellers     []int         `bson:"resellers"`              // 通知代理商
	Channels      []string      `bson:"channels,omitempty"`     // 发送渠道[web:站内信, email:邮件]
	IsPopup       bool          `bson:"is_popup"`               // 是否弹窗
	Priority      int           `bson:"priority,omitempty"`     // 优先级
	ExpiresAt     int64         `bson:"expires_at,omitempty"`   // 过期时间
	ScheduledTime int64         `bson:"scheduled_time"`         // 定时发布时间
	Remark        string        `bson:"remark"`                 // 备注
	Status        int           `bson:"status,omitempty"`       // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
	Reads         []common.Read `bson:"reads,omitempty"`        // 已读
	Variables     []string      `bson:"variables"`              // 变量
	Publisher     int           `bson:"publisher,omitempty"`    // 发布人
	PublishTime   int64         `bson:"publish_time,omitempty"` // 发布时间
	Rid           int           `bson:"rid,omitempty"`          // 代理商ID
	Creator       string        `bson:"creator,omitempty"`      // 创建人
	Updater       string        `bson:"updater,omitempty"`      // 更新人
	CreatedAt     int64         `bson:"created_at,omitempty"`   // 创建时间
	UpdatedAt     int64         `bson:"updated_at,omitempty"`   // 更新时间
}
