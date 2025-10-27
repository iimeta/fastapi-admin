package model

import "time"

// 仪表盘基础数据接口响应参数
type DashboardBaseDataRes struct {
	Dashboard
}

// 仪表盘调用数据接口请求参数
type DashboardCallDataReq struct {
	Days int `json:"days"` // 天数
}

// 仪表盘调用数据接口响应参数
type DashboardCallDataRes struct {
	Items []*CallData `json:"items"`
}

// 仪表盘费用接口响应参数
type DashboardExpenseRes struct {
	*Expense
}

// 仪表盘数据TOP接口请求参数
type DashboardDataTopReq struct {
	Days     int    `json:"days"`      // 天数
	DataType string `json:"data_type"` // 数据类型
}

// 仪表盘数据TOP接口响应参数
type DashboardDataTopRes struct {
	Items []*DataTop `json:"items"`
}

// 仪表盘模型占比接口请求参数
type DashboardModelPercentReq struct {
	Days int `json:"days"` // 天数
}

// 仪表盘模型占比接口响应参数
type DashboardModelPercentRes struct {
	Models []string        `json:"models"`
	Items  []*ModelPercent `json:"items"`
}

// 每分钟数据接口请求参数
type DashboardPerMinuteReq struct {
	TraceId     string   `json:"trace_id,omitempty"`     // 日志ID
	UserId      int      `json:"user_id,omitempty"`      // 用户ID
	AppId       int      `json:"app_id,omitempty"`       // 应用ID
	Key         string   `json:"key,omitempty"`          // 密钥
	Models      []string `json:"models,omitempty"`       // 模型
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
	TotalTime   int64    `json:"total_time,omitempty"`   // 总时间
	Status      int      `json:"status,omitempty"`       // 状态[1:成功, -1:失败]
	ReqTime     []string `json:"req_time,omitempty"`     // 请求时间
}

// 每分钟数据接口响应参数
type DashboardPerMinuteRes struct {
	RPM int `json:"rpm,omitempty"` // 每分钟请求数
	TPM int `json:"tpm,omitempty"` // 每分钟令牌数
}

// 每秒钟数据接口请求参数
type DashboardPerSecondReq struct {
	TraceId     string   `json:"trace_id,omitempty"`     // 日志ID
	UserId      int      `json:"user_id,omitempty"`      // 用户ID
	AppId       int      `json:"app_id,omitempty"`       // 应用ID
	Key         string   `json:"key,omitempty"`          // 密钥
	Models      []string `json:"models,omitempty"`       // 模型
	ModelAgents []string `json:"model_agents,omitempty"` // 模型代理
	TotalTime   int64    `json:"total_time,omitempty"`   // 总时间
	Status      int      `json:"status,omitempty"`       // 状态[1:成功, -1:失败]
	ReqTime     []string `json:"req_time,omitempty"`     // 请求时间
}

// 每秒钟数据接口响应参数
type DashboardPerSecondRes struct {
	RPS int `json:"rps,omitempty"` // 每秒钟请求数
	TPS int `json:"tps,omitempty"` // 每秒钟令牌数
}

// 额度预警接口请求参数
type DashboardQuotaWarningReq struct {
	QuotaWarning           bool          `json:"quota_warning,omitempty"`            // 额度预警开关
	WarningThreshold       int           `json:"warning_threshold,omitempty"`        // 预警阈值, 单位: $
	ExpireWarningThreshold time.Duration `json:"expire_warning_threshold,omitempty"` // 过期预警阈值, 单位: 天
}

// 基础数据
type Dashboard struct {
	App       int64 `json:"app,omitempty"`        // 应用数
	TodayApp  int64 `json:"today_app,omitempty"`  // 今日新增应用数
	Model     int64 `json:"model,omitempty"`      // 模型数
	AppKey    int64 `json:"app_key,omitempty"`    // 应用密钥数
	ModelKey  int64 `json:"model_key,omitempty"`  // 模型密钥数
	User      int64 `json:"user,omitempty"`       // 用户数
	TodayUser int64 `json:"today_user,omitempty"` // 今日新增用户数
	Call      int   `json:"call,omitempty"`       // 调用数
	Group     int   `json:"group,omitempty"`      // 分组数
}

// 调用数据
type CallData struct {
	Date     string  `json:"date,omitempty"`     // 日期
	Spend    float64 `json:"spend,omitempty"`    // 花费($)
	Call     int     `json:"call,omitempty"`     // 调用数
	Tokens   int     `json:"tokens,omitempty"`   // 令牌数
	User     int     `json:"user,omitempty"`     // 用户数
	App      int     `json:"app,omitempty"`      // 应用数
	Abnormal int     `json:"abnormal,omitempty"` // 异常数
}

// 费用
type Expense struct {
	Quota                  float64       `json:"quota,omitempty"`                    // 剩余额度
	UsedQuota              float64       `json:"used_quota,omitempty"`               // 已用额度
	AllocatedQuota         float64       `json:"allocated_quota,omitempty"`          // 已分配额度
	ToBeAllocatedQuota     float64       `json:"to_be_allocated_quota,omitempty"`    // 待分配额度
	QuotaExpiresAt         string        `json:"quota_expires_at,omitempty"`         // 额度过期时间
	QuotaWarning           bool          `json:"quota_warning,omitempty"`            // 额度预警开关
	WarningThreshold       int           `json:"warning_threshold,omitempty"`        // 预警阈值, 单位: $
	ExpireWarningThreshold time.Duration `json:"expire_warning_threshold,omitempty"` // 过期预警阈值, 单位: 天
}

// 数据TOP
type DataTop struct {
	UserId int    `json:"user_id,omitempty"` // 用户ID
	AppId  int    `json:"app_id,omitempty"`  // 应用ID
	AppKey string `json:"app_key,omitempty"` // 应用密钥
	Model  string `json:"model,omitempty"`   // 模型
	Call   int    `json:"call,omitempty"`    // 调用数
	Models int    `json:"models,omitempty"`  // 模型数
	Tokens int    `json:"tokens,omitempty"`  // 令牌数
	User   int    `json:"user,omitempty"`    // 用户数
	App    int    `json:"app,omitempty"`     // 应用数
}

// 模型占比
type ModelPercent struct {
	Name  string `json:"name,omitempty"`  // 模型
	Value int    `json:"value,omitempty"` // 调用数
}
