package do

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

type Group struct {
	gmeta.Meta         `collection:"group" bson:"-"`
	Name               string                `bson:"name,omitempty"`        // 分组名称
	Discount           float64               `bson:"discount,omitempty"`    // 分组折扣
	Models             []string              `bson:"models"`                // 模型权限
	IsEnableModelAgent bool                  `bson:"is_enable_model_agent"` // 是否启用模型代理
	LbStrategy         int                   `bson:"lb_strategy,omitempty"` // 代理负载均衡策略[1:轮询, 2:权重]
	ModelAgents        []string              `bson:"model_agents"`          // 模型代理
	IsDefault          bool                  `bson:"is_default"`            // 是否默认分组
	IsLimitQuota       bool                  `bson:"is_limit_quota"`        // 是否限制额度
	Quota              int                   `bson:"quota"`                 // 剩余额度
	UsedQuota          int                   `bson:"used_quota,omitempty"`  // 已用额度
	IsEnableForward    bool                  `bson:"is_enable_forward"`     // 是否启用模型转发
	ForwardConfig      *common.ForwardConfig `bson:"forward_config"`        // 模型转发配置
	IsPublic           bool                  `bson:"is_public"`             // 是否公开
	Weight             int                   `bson:"weight"`                // 权重
	ExpiresAt          int64                 `bson:"expires_at"`            // 过期时间
	Remark             string                `bson:"remark"`                // 备注
	Status             int                   `bson:"status,omitempty"`      // 状态[1:正常, 2:禁用, -1:删除]
	Creator            string                `bson:"creator,omitempty"`     // 创建人
	Updater            string                `bson:"updater,omitempty"`     // 更新人
	CreatedAt          int64                 `bson:"created_at,omitempty"`  // 创建时间
	UpdatedAt          int64                 `bson:"updated_at,omitempty"`  // 更新时间
}
