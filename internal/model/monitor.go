package model

// 监控中心实时性能维度分析接口请求参数
type MonitorPerfBreakdownReq struct {
	Dimension string `json:"dimension"` // 维度: model/model_agent/user/app/app_key/provider/key
	Limit     int    `json:"limit"`     // 数量限制, 默认10
	Window    int    `json:"window"`    // 数据窗口(秒), 默认10
}

// 性能维度项
type MonitorPerfBreakdownItem struct {
	Name            string  `json:"name"`               // 维度名称
	RPS             int     `json:"rps"`                // 每秒请求数
	TPS             int     `json:"tps"`                // 每秒令牌数
	RPM             int     `json:"rpm"`                // 每分钟请求数
	TPM             int     `json:"tpm"`                // 每分钟令牌数
	AvgTotalTime    int64   `json:"avg_total_time"`     // 平均总耗时(ms)
	AvgConnTime     int64   `json:"avg_conn_time"`      // 平均连接耗时(ms)
	AvgDuration     int64   `json:"avg_duration"`       // 平均持续时间(ms)
	AvgInternalTime int64   `json:"avg_internal_time"`  // 平均内耗时间(ms)
	SuccessRate     float64 `json:"success_rate"`       // 成功率(%)
	ErrorCount      int     `json:"error_count"`        // 错误数
	InputTokens     int     `json:"input_tokens"`       // 输入Token数
	OutputTokens    int     `json:"output_tokens"`      // 输出Token数
	AvgTokensPerReq int     `json:"avg_tokens_per_req"` // 平均每请求Token数
}

// 监控中心实时性能维度分析接口响应参数
type MonitorPerfBreakdownRes struct {
	Items []*MonitorPerfBreakdownItem `json:"items"`
}

// 监控中心历史性能数据接口请求参数
type MonitorPerfHistoryReq struct {
	Dimension string `json:"dimension"` // 维度
	Range     string `json:"range"`     // 时间范围: 1h/3h/6h/12h/1d/2d/3d
	Metric    string `json:"metric"`    // 指标
	Limit     int    `json:"limit"`     // 数量限制, 默认10
}

// 监控中心历史性能数据接口响应参数
type MonitorPerfHistoryRes struct {
	Dates  []string             `json:"dates"`  // 时间标签列表
	Series map[string][]float64 `json:"series"` // 维度名 -> 各时间点的指标值
}

// 监控中心全局实时指标接口请求参数
type MonitorGlobalReq struct{}

// 监控中心全局实时指标接口响应参数
type MonitorGlobalRes struct {
	RPS int `json:"rps"` // 每秒请求数
	TPS int `json:"tps"` // 每秒令牌数
	RPM int `json:"rpm"` // 每分钟请求数
	TPM int `json:"tpm"` // 每分钟令牌数
}
