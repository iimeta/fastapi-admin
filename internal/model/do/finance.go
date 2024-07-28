package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	DEAL_RECORD_COLLECTION = "deal_record"
)

type DealRecord struct {
	gmeta.Meta `collection:"deal_record" bson:"-"`
	UserId     int    `bson:"user_id,omitempty"`    // 用户ID
	Quota      int    `bson:"quota"`                // 充值额度
	Remark     string `bson:"remark,omitempty"`     // 备注
	Status     int    `bson:"status,omitempty"`     // 状态[1:正常, 2:退款, -1:删除]
	Creator    string `bson:"creator,omitempty"`    // 创建人
	Updater    string `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}
