package do

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

const (
	NOTICE_TEMPLATE_COLLECTION = "notice_template"
)

type NoticeTemplate struct {
	gmeta.Meta `collection:"notice_template" bson:"-"`
	Name       string `bson:"name,omitempty"`       // 名称
	Action     string `bson:"action,omitempty"`     // 动作
	Content    string `bson:"content,omitempty"`    // 内容
	Category   int    `bson:"category,omitempty"`   // 分类[1:系统公告, 2:活动通知, 3:维护通知]
	IsPublic   bool   `bson:"is_public,omitempty"`  // 是否公开
	Remark     string `bson:"remark"`               // 备注
	Status     int    `bson:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	UserId     int    `bson:"user_id,omitempty"`    // 用户ID
	Rid        int    `bson:"rid,omitempty"`        // 代理商ID
	Creator    string `bson:"creator,omitempty"`    // 创建人
	Updater    string `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}
