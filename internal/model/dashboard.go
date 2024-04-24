package model

// 仪表盘基础数据接口响应参数
type DashboardBaseDataRes struct {
	*Dashboard
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

// 基础数据
type Dashboard struct {
	App       int64 `json:"app"`        // 应用数
	TodayApp  int64 `json:"today_app"`  // 今日新增应用数
	Model     int64 `json:"model"`      // 模型数
	AppKey    int64 `json:"app_key"`    // 应用密钥数
	ModelKey  int64 `json:"model_key"`  // 模型密钥数
	User      int64 `json:"user"`       // 用户数
	TodayUser int64 `json:"today_user"` // 今日新增用户数
	Call      int   `json:"call"`       // 调用数
}

// 调用数据
type CallData struct {
	Date   string `json:"date"`   // 日期
	Call   int    `json:"call"`   // 调用数
	Tokens int    `json:"tokens"` // 令牌数
	User   int    `json:"user"`   // 用户数
	App    int    `json:"app"`    // 应用数
}

// 费用
type Expense struct {
	Quota        int     `json:"quota"`          // 剩余额度
	QuotaUSD     float64 `json:"quota_usd"`      // 剩余额度美元单位
	UsedQuota    int     `json:"used_quota"`     // 已用额度
	UsedQuotaUSD float64 `json:"used_quota_usd"` // 已用额度美元单位
}

// 数据TOP
type DataTop struct {
	UserId int    `json:"user_id,omitempty"` // 用户ID
	AppId  int    `json:"app_id,omitempty"`  // 应用ID
	AppKey string `json:"app_key,omitempty"` // 应用密钥
	Model  string `json:"model,omitempty"`   // 模型
	Call   int    `json:"call"`              // 调用数
	Models int    `json:"models"`            // 模型数
	Tokens int    `json:"tokens"`            // 令牌数
	User   int    `json:"user"`              // 用户数
	App    int    `json:"app"`               // 应用数
}

// 模型占比
type ModelPercent struct {
	Name  string `json:"name"`  // 模型
	Value int    `json:"value"` // 调用数
}
