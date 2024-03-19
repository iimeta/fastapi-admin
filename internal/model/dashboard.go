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

type CallData struct {
	Date   string `json:"date"`   // 日期
	Count  int    `json:"count"`  // 调用数
	Tokens int    `json:"tokens"` // 令牌数
	User   int    `json:"user"`   // 用户数
	App    int    `json:"app"`    // 应用数
}

type Expense struct {
	Quota int `json:"quota"` // 额度
}
