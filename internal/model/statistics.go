package model

import (
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

// 统计数据接口请求参数
type StatisticsDataReq struct {
	UserId        int    `json:"user_id"`         // 用户ID
	AppId         int    `json:"app_id"`          // 应用ID
	AppKey        string `json:"app_key"`         // 应用密钥
	StatStartTime int64  `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64  `json:"stat_end_time"`   // 统计结束时间
}

// 统计数据接口响应参数
type StatisticsDataRes struct {
	Total    int     `json:"total,omitempty"`    // 总数
	Tokens   float64 `json:"tokens,omitempty"`   // 令牌数
	Abnormal int     `json:"abnormal,omitempty"` // 异常数
}

type StatisticsUser struct {
	Id             string              `json:"id,omitempty"`              // ID
	UserId         int                 `json:"user_id,omitempty"`         // 用户ID
	StatDate       string              `json:"stat_date,omitempty"`       // 统计日期
	StatTime       int64               `json:"stat_time,omitempty"`       // 统计时间
	Total          int                 `json:"total,omitempty"`           // 总数
	Tokens         float64             `json:"tokens,omitempty"`          // 令牌数
	Abnormal       int                 `json:"abnormal,omitempty"`        // 异常数
	AbnormalTokens float64             `json:"abnormal_tokens,omitempty"` // 异常令牌数
	ModelStats     []*common.ModelStat `json:"model_stats,omitempty"`     // 模型统计数据
	Creator        string              `json:"creator,omitempty"`         // 创建人
	Updater        string              `json:"updater,omitempty"`         // 更新人
	CreatedAt      string              `json:"created_at,omitempty"`      // 创建时间
	UpdatedAt      string              `json:"updated_at,omitempty"`      // 更新时间
}
