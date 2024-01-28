package model

// 仪表盘基础数据接口响应参数
type DashboardBaseDataRes struct {
	*Dashboard
}

// 仪表盘费用接口响应参数
type DashboardExpenseRes struct {
	*Expense
}

type Dashboard struct {
	App       int64 `json:"app,omitempty"`        // 应用数
	TodayApp  int64 `json:"today_app,omitempty"`  // 今日新增应用数
	Model     int64 `json:"model,omitempty"`      // 模型数
	AppKey    int64 `json:"app_key,omitempty"`    // 应用密钥数
	ModelKey  int64 `json:"model_key,omitempty"`  // 模型密钥数
	User      int64 `json:"user,omitempty"`       // 用户数
	TodayUser int64 `json:"today_user,omitempty"` // 今日新增用户数
}

type Expense struct {
	Quota int `json:"quota"` // 额度
}
