package entity

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

type Notice struct {
	gmeta.Meta    `role:"*" bson:"-"`
	Id            string        `bson:"_id,omitempty"`            // ID
	Title         string        `bson:"title,omitempty"`          // 标题
	Content       string        `bson:"content,omitempty"`        // 内容
	Category      int           `bson:"category,omitempty"`       // 分类[1:系统公告, 2:活动通知, 3:维护通知]
	Scope         int           `bson:"scope,omitempty"`          // 通知范围[1:全部, 2:全部用户, 3:全部代理商, 4:指定用户, 5:指定代理商, 6:指定用户和代理商]
	Users         []int         `bson:"users,omitempty"`          // 通知用户
	Resellers     []int         `bson:"resellers,omitempty"`      // 通知代理商
	Channels      []string      `bson:"channels,omitempty"`       // 发送渠道[web:站内信, email:邮件]
	IsPopup       bool          `bson:"is_popup,omitempty"`       // 是否弹窗
	Priority      int           `bson:"priority,omitempty"`       // 优先级
	ExpiresAt     int64         `bson:"expires_at,omitempty"`     // 过期时间
	ScheduledTime int64         `bson:"scheduled_time,omitempty"` // 定时发布时间
	Remark        string        `bson:"remark,omitempty"`         // 备注
	Status        int           `bson:"status,omitempty"`         // 状态[1:发布, 2:草稿, 3:定时, 4:过期, -1:删除]
	Variables     []string      `bson:"variables,omitempty"`      // 变量
	Reads         []common.Read `bson:"reads,omitempty"`          // 已读
	UserId        int           `bson:"user_id,omitempty"`        // 用户ID
	PublishTime   int64         `bson:"publish_time,omitempty"`   // 发布时间
	Rid           int           `bson:"rid,omitempty"`            // 代理商ID
	Creator       string        `bson:"creator,omitempty"`        // 创建人
	Updater       string        `bson:"updater,omitempty"`        // 更新人
	CreatedAt     int64         `bson:"created_at,omitempty"`     // 创建时间
	UpdatedAt     int64         `bson:"updated_at,omitempty"`     // 更新时间
}
