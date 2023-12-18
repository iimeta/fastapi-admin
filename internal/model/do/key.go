package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	KEY_COLLECTION = "key"
)

type Key struct {
	gmeta.Meta `collection:"key" bson:"-"`
	Corp       string   `bson:"corp,omitempty"`       // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Key        string   `bson:"key,omitempty"`        // 密钥
	Models     []string `bson:"models,omitempty"`     // 模型
	Remark     string   `bson:"remark,omitempty"`     // 备注
	Status     int      `bson:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Creator    string   `bson:"creator,omitempty"`    // 创建人
	Updater    string   `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64    `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64    `bson:"updated_at,omitempty"` // 更新时间
}
