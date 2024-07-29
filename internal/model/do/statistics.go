package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	STATISTICS_USER_COLLECTION    = "statistics_user"
	STATISTICS_APP_COLLECTION     = "statistics_app"
	STATISTICS_APP_KEY_COLLECTION = "statistics_app_key"
)

type StatisticsUser struct {
	gmeta.Meta `collection:"statistics_user" bson:"-"`
	UserId     int    `bson:"user_id,omitempty"`    // 用户ID
	StatDate   string `bson:"stat_date,omitempty"`  // 统计日期
	StatTime   int64  `bson:"stat_time,omitempty"`  // 统计时间
	Total      int    `bson:"total,omitempty"`      // 总数
	Tokens     int    `bson:"tokens,omitempty"`     // 令牌数
	Abnormal   int    `bson:"abnormal,omitempty"`   // 异常数
	Creator    string `bson:"creator,omitempty"`    // 创建人
	Updater    string `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}

type StatisticsApp struct {
	gmeta.Meta `collection:"statistics_app" bson:"-"`
	UserId     int    `bson:"user_id,omitempty"`    // 用户ID
	AppId      int    `bson:"app_id,omitempty"`     // 应用ID
	StatDate   string `bson:"stat_date,omitempty"`  // 统计日期
	StatTime   int64  `bson:"stat_time,omitempty"`  // 统计时间
	Total      int    `bson:"total,omitempty"`      // 总数
	Tokens     int    `bson:"tokens,omitempty"`     // 令牌数
	Abnormal   int    `bson:"abnormal,omitempty"`   // 异常数
	Creator    string `bson:"creator,omitempty"`    // 创建人
	Updater    string `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}

type StatisticsAppKey struct {
	gmeta.Meta `collection:"statistics_app_key" bson:"-"`
	UserId     int    `bson:"user_id,omitempty"`    // 用户ID
	AppId      int    `bson:"app_id,omitempty"`     // 应用ID
	AppKey     string `bson:"app_key,omitempty"`    // 应用密钥
	StatDate   string `bson:"stat_date,omitempty"`  // 统计日期
	StatTime   int64  `bson:"stat_time,omitempty"`  // 统计时间
	Total      int    `bson:"total,omitempty"`      // 总数
	Tokens     int    `bson:"tokens,omitempty"`     // 令牌数
	Abnormal   int    `bson:"abnormal,omitempty"`   // 异常数
	Creator    string `bson:"creator,omitempty"`    // 创建人
	Updater    string `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}
