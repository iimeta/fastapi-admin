package entity

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

type StatisticsUser struct {
	gmeta.Meta     `role:"*" bson:"-"`
	Id             string              `bson:"_id,omitempty"`             // ID
	UserId         int                 `bson:"user_id,omitempty"`         // 用户ID
	StatDate       string              `bson:"stat_date,omitempty"`       // 统计日期
	StatTime       int64               `bson:"stat_time,omitempty"`       // 统计时间
	Total          int                 `bson:"total,omitempty"`           // 总数
	Tokens         int                 `bson:"tokens,omitempty"`          // 令牌数
	Abnormal       int                 `bson:"abnormal,omitempty"`        // 异常数
	AbnormalTokens int                 `bson:"abnormal_tokens,omitempty"` // 异常令牌数
	ModelStats     []*common.ModelStat `bson:"model_stats,omitempty"`     // 模型统计数据
	Creator        string              `bson:"creator,omitempty"`         // 创建人
	Updater        string              `bson:"updater,omitempty"`         // 更新人
	CreatedAt      int64               `bson:"created_at,omitempty"`      // 创建时间
	UpdatedAt      int64               `bson:"updated_at,omitempty"`      // 更新时间
}

type StatisticsApp struct {
	gmeta.Meta     `role:"*" bson:"-"`
	Id             string              `bson:"_id,omitempty"`             // ID
	UserId         int                 `bson:"user_id,omitempty"`         // 用户ID
	AppId          int                 `bson:"app_id,omitempty"`          // 应用ID
	StatDate       string              `bson:"stat_date,omitempty"`       // 统计日期
	StatTime       int64               `bson:"stat_time,omitempty"`       // 统计时间
	Total          int                 `bson:"total,omitempty"`           // 总数
	Tokens         int                 `bson:"tokens,omitempty"`          // 令牌数
	Abnormal       int                 `bson:"abnormal,omitempty"`        // 异常数
	AbnormalTokens int                 `bson:"abnormal_tokens,omitempty"` // 异常令牌数
	ModelStats     []*common.ModelStat `bson:"model_stats,omitempty"`     // 模型统计数据
	Creator        string              `bson:"creator,omitempty"`         // 创建人
	Updater        string              `bson:"updater,omitempty"`         // 更新人
	CreatedAt      int64               `bson:"created_at,omitempty"`      // 创建时间
	UpdatedAt      int64               `bson:"updated_at,omitempty"`      // 更新时间
}

type StatisticsAppKey struct {
	gmeta.Meta     `role:"*" bson:"-"`
	Id             string              `bson:"_id,omitempty"`             // ID
	UserId         int                 `bson:"user_id,omitempty"`         // 用户ID
	AppId          int                 `bson:"app_id,omitempty"`          // 应用ID
	AppKey         string              `bson:"app_key,omitempty"`         // 应用密钥
	StatDate       string              `bson:"stat_date,omitempty"`       // 统计日期
	StatTime       int64               `bson:"stat_time,omitempty"`       // 统计时间
	Total          int                 `bson:"total,omitempty"`           // 总数
	Tokens         int                 `bson:"tokens,omitempty"`          // 令牌数
	Abnormal       int                 `bson:"abnormal,omitempty"`        // 异常数
	AbnormalTokens int                 `bson:"abnormal_tokens,omitempty"` // 异常令牌数
	ModelStats     []*common.ModelStat `bson:"model_stats,omitempty"`     // 模型统计数据
	Creator        string              `bson:"creator,omitempty"`         // 创建人
	Updater        string              `bson:"updater,omitempty"`         // 更新人
	CreatedAt      int64               `bson:"created_at,omitempty"`      // 创建时间
	UpdatedAt      int64               `bson:"updated_at,omitempty"`      // 更新时间
}
