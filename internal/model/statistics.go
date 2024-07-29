package model

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
	Total    int `json:"total"`    // 总数
	Tokens   int `json:"tokens"`   // 令牌数
	Abnormal int `json:"abnormal"` // 异常数
}
