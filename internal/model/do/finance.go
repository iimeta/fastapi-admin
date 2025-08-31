package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	DEAL_RECORD_COLLECTION = "deal_record"
)

type DealRecord struct {
	gmeta.Meta `collection:"deal_record" bson:"-"`
	UserId     int    `bson:"user_id,omitempty"`    // 用户ID
	Quota      int    `bson:"quota,omitempty"`      // 额度
	Type       int    `bson:"type,omitempty"`       // 交易类型[1:充值, 2:扣除, 3:赠送, 4:过期]
	Remark     string `bson:"remark,omitempty"`     // 备注
	Status     int    `bson:"status,omitempty"`     // 状态[1:成功, 2:退款, 3:失败, -1:删除]
	Rid        int    `bson:"rid,omitempty"`        // 代理商ID
	Creator    string `bson:"creator,omitempty"`    // 创建人
	Updater    string `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}
