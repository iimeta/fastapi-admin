package model

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/common"
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

// 数据看板汇总接口请求参数
type StatisticsSummaryReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	Rid           int      `json:"rid"`             // 代理商ID
	UserId        int      `json:"user_id"`         // 用户ID
	AppId         int      `json:"app_id"`          // 应用ID
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用)
	Models        []string `json:"models"`          // 模型
	Provider      string   `json:"provider"`        // 提供商ID
}

// 数据看板汇总接口响应参数
type StatisticsSummaryRes struct {
	Total            int     `json:"total"`              // 总调用数
	Tokens           float64 `json:"tokens"`             // 总令牌数
	Abnormal         int     `json:"abnormal"`           // 异常数
	AbnormalRate     float64 `json:"abnormal_rate"`      // 异常率
	ActiveUsers      int     `json:"active_users"`       // 活跃用户数
	ActiveApps       int     `json:"active_apps"`        // 活跃应用数
	PrevTotal        int     `json:"prev_total"`         // 上一周期总调用数
	PrevTokens       float64 `json:"prev_tokens"`        // 上一周期总令牌数
	PrevAbnormal     int     `json:"prev_abnormal"`      // 上一周期异常数
	PrevAbnormalRate float64 `json:"prev_abnormal_rate"` // 上一周期异常率
}

// 数据看板趋势接口请求参数
type StatisticsTrendReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	Rid           int      `json:"rid"`             // 代理商ID
	UserId        int      `json:"user_id"`         // 用户ID
	AppId         int      `json:"app_id"`          // 应用ID
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用)
	Models        []string `json:"models"`          // 模型
	Provider      string   `json:"provider"`        // 提供商ID
}

// 数据看板趋势接口响应参数
type StatisticsTrendRes struct {
	Items []*StatisticsTrendItem `json:"items"`
}

// 趋势数据项
type StatisticsTrendItem struct {
	Date         string  `json:"date"`          // 日期
	Total        int     `json:"total"`         // 总调用数
	Tokens       float64 `json:"tokens"`        // 总令牌数
	Abnormal     int     `json:"abnormal"`      // 异常数
	AbnormalRate float64 `json:"abnormal_rate"` // 异常率
	ActiveUsers  int     `json:"active_users"`  // 活跃用户数
	ActiveApps   int     `json:"active_apps"`   // 活跃应用数
}

// 数据看板模型分布接口请求参数
type StatisticsModelPercentReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	Rid           int      `json:"rid"`             // 代理商ID
	UserId        int      `json:"user_id"`         // 用户ID
	AppId         int      `json:"app_id"`          // 应用ID
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用)
	Models        []string `json:"models"`          // 模型列表
	Provider      string   `json:"provider"`        // 提供商名称
	DataType      string   `json:"data_type"`       // 数据类型: calls(默认)/tokens
}

// 数据看板模型分布接口响应参数
type StatisticsModelPercentRes struct {
	Models []string        `json:"models"` // 模型列表
	Items  []*ModelPercent `json:"items"`  // 模型占比数据
}

// 数据看板排行接口请求参数
type StatisticsTopReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	DataType      string   `json:"data_type"`       // 数据类型: user/app/app_key/model/provider
	Rid           int      `json:"rid"`             // 代理商ID
	UserId        int      `json:"user_id"`         // 用户ID
	AppId         int      `json:"app_id"`          // 应用ID
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用)
	Models        []string `json:"models"`          // 模型列表
	Provider      string   `json:"provider"`        // 提供商名称
	Limit         int      `json:"limit"`           // 数量限制
}

// 数据看板排行接口响应参数
type StatisticsTopRes struct {
	Items []*DataTop `json:"items"`
}

// 数据看板明细接口请求参数
type StatisticsDetailReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	DataType      string   `json:"data_type"`       // 数据类型: user/app/app_key/model
	Rid           int      `json:"rid"`             // 代理商ID
	UserId        int      `json:"user_id"`         // 用户ID
	AppId         int      `json:"app_id"`          // 应用ID
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用)
	Models        []string `json:"models"`          // 模型列表
	Provider      string   `json:"provider"`        // 提供商名称
	ModelId       string   `json:"model_id"`        // 模型ID
	Paging
}

// 数据看板明细数据项
type StatisticsDetailItem struct {
	UserId       int                 `json:"user_id,omitempty"`     // 用户ID
	AppId        int                 `json:"app_id,omitempty"`      // 应用ID
	AppKey       string              `json:"app_key,omitempty"`     // 应用密钥
	Model        string              `json:"model,omitempty"`       // 模型名
	StatDate     string              `json:"stat_date"`             // 统计日期
	Total        int                 `json:"total"`                 // 总调用数
	Tokens       float64             `json:"tokens"`                // 令牌数
	Abnormal     int                 `json:"abnormal"`              // 异常数
	AbnormalRate float64             `json:"abnormal_rate"`         // 异常率
	ModelStats   []*common.ModelStat `json:"model_stats,omitempty"` // 模型统计数据
}

// 数据看板明细接口响应参数
type StatisticsDetailRes struct {
	Items  []*StatisticsDetailItem `json:"items"`
	Paging *Paging                 `json:"paging"`
}

// 数据看板全局总览接口请求参数
type StatisticsOverviewReq struct {
	Rid    int    `json:"rid"`     // 代理商ID
	UserId int    `json:"user_id"` // 用户ID
	AppId  int    `json:"app_id"`  // 应用ID
	AppKey string `json:"app_key"` // 应用密钥
}

// 数据看板全局总览接口响应参数
type StatisticsOverviewRes struct {
	TotalCalls      int     `json:"total_calls"`       // 历史总调用数
	TotalTokens     float64 `json:"total_tokens"`      // 历史总花费
	TotalAbnormal   int     `json:"total_abnormal"`    // 历史总异常数
	AbnormalRate    float64 `json:"abnormal_rate"`     // 历史异常率
	TotalUsers      int64   `json:"total_users"`       // 总用户数
	TotalApps       int64   `json:"total_apps"`        // 总应用数
	TotalAppKeys    int64   `json:"total_app_keys"`    // 总密钥数
	TotalModels     int64   `json:"total_models"`      // 总模型数
	TotalModelKeys  int64   `json:"total_model_keys"`  // 总模型密钥数(仅admin)
	TotalAgents     int64   `json:"total_agents"`      // 总模型代理数(仅admin)
	TotalProviders  int64   `json:"total_providers"`   // 总提供商数(仅admin)
	TotalGroups     int     `json:"total_groups"`      // 总分组数
	TodayCalls      int     `json:"today_calls"`       // 今日调用数
	TodayTokens     float64 `json:"today_tokens"`      // 今日花费
	TodayAbnormal   int     `json:"today_abnormal"`    // 今日异常数
	TodayUsers      int64   `json:"today_users"`       // 今日新增用户
	TodayApps       int64   `json:"today_apps"`        // 今日新增应用
	TotalBatchTasks int64   `json:"total_batch_tasks"` // 总批处理任务数
	TotalFileTasks  int64   `json:"total_file_tasks"`  // 总文件任务数
	TotalImageTasks int64   `json:"total_image_tasks"` // 总绘图任务数
	TotalVideoTasks int64   `json:"total_video_tasks"` // 总视频任务数
}

// 数据看板模型趋势接口请求参数
type StatisticsModelTrendReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	Rid           int      `json:"rid"`             // 代理商ID
	UserId        int      `json:"user_id"`         // 用户ID
	AppId         int      `json:"app_id"`          // 应用ID
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用)
	Models        []string `json:"models"`          // 指定模型列表(为空则取top5)
	Provider      string   `json:"provider"`        // 提供商ID
}

// 数据看板模型趋势接口响应参数
type StatisticsModelTrendRes struct {
	Models []string                     `json:"models"` // 模型名列表
	Dates  []string                     `json:"dates"`  // 日期列表
	Series map[string]*ModelTrendSeries `json:"series"` // 按模型名索引
}

// 模型趋势数据系列
type ModelTrendSeries struct {
	Calls    []int     `json:"calls"`    // 每日调用数
	Tokens   []float64 `json:"tokens"`   // 每日花费
	Abnormal []int     `json:"abnormal"` // 每日异常数
}

// 数据看板密钥状态接口请求参数(仅admin)
type StatisticsKeyStatusReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用,匹配key.key或key.creator)
	Models        []string `json:"models"`          // 模型列表
	Provider      string   `json:"provider"`        // 提供商ID
}

// 数据看板密钥状态接口响应参数
type StatisticsKeyStatusRes struct {
	Total        int64      `json:"total"`         // 总密钥数
	Active       int64      `json:"active"`        // 正常数
	Disabled     int64      `json:"disabled"`      // 禁用数
	AutoDisabled int64      `json:"auto_disabled"` // 自动禁用数
	ByKey        []*KeyStat `json:"by_key"`        // 按密钥统计
}

// 按密钥统计
type KeyStat struct {
	Key       string  `json:"key"`        // 密钥(后5位)
	Status    int     `json:"status"`     // 状态[1:正常, 2:禁用]
	UsedQuota float64 `json:"used_quota"` // 已用额度
}

// 数据看板代理状态接口请求参数(仅admin)
type StatisticsAgentStatusReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用,匹配key.key或key.creator)
	Models        []string `json:"models"`          // 模型列表
	Provider      string   `json:"provider"`        // 提供商ID
}

// 数据看板代理状态接口响应参数
type StatisticsAgentStatusRes struct {
	Total        int64        `json:"total"`         // 总代理数
	Active       int64        `json:"active"`        // 正常数
	Disabled     int64        `json:"disabled"`      // 禁用数
	AutoDisabled int64        `json:"auto_disabled"` // 自动禁用数
	ByAgent      []*AgentStat `json:"by_agent"`      // 按代理统计
}

// 按代理统计
type AgentStat struct {
	Name      string  `json:"name"`       // 代理名称
	Status    int     `json:"status"`     // 状态[1:正常, 2:禁用]
	UsedQuota float64 `json:"used_quota"` // 已用额度
}

// 数据看板任务状态分布接口请求参数
type StatisticsTaskStatusReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	Rid           int      `json:"rid"`             // 代理商ID
	UserId        int      `json:"user_id"`         // 用户ID
	AppId         int      `json:"app_id"`          // 应用ID
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用,匹配creator)
	Models        []string `json:"models"`          // 模型列表
	Provider      string   `json:"provider"`        // 提供商ID
}

// 任务状态项
type TaskStatusItem struct {
	Status string `json:"status"` // 状态
	Count  int64  `json:"count"`  // 数量
}

// 数据看板任务状态分布接口响应参数
type StatisticsTaskStatusRes struct {
	Batch       []*TaskStatusItem `json:"batch"`        // 批处理任务状态
	File        []*TaskStatusItem `json:"file"`         // 文件任务状态
	Image       []*TaskStatusItem `json:"image"`        // 绘图任务状态
	Video       []*TaskStatusItem `json:"video"`        // 视频任务状态
	ActiveBatch int64             `json:"active_batch"` // 进行中批处理数
	ActiveImage int64             `json:"active_image"` // 进行中绘图任务数
	ActiveVideo int64             `json:"active_video"` // 进行中视频任务数
	QueuedImage int64             `json:"queued_image"` // 排队中绘图任务数
	QueuedVideo int64             `json:"queued_video"` // 排队中视频任务数
	QueuedBatch int64             `json:"queued_batch"` // 排队中批处理数
}

// 数据看板响应耗时趋势接口请求参数
type StatisticsLatencyTrendReq struct {
	StatStartTime int64    `json:"stat_start_time"` // 统计开始时间
	StatEndTime   int64    `json:"stat_end_time"`   // 统计结束时间
	Rid           int      `json:"rid"`             // 代理商ID
	UserId        int      `json:"user_id"`         // 用户ID
	AppId         int      `json:"app_id"`          // 应用ID
	AppKey        string   `json:"app_key"`         // 应用密钥
	Key           string   `json:"key"`             // 密钥筛选(管理员专用)
	Models        []string `json:"models"`          // 模型列表
	Provider      string   `json:"provider"`        // 提供商名称
}

// 数据看板响应耗时趋势接口响应参数
type StatisticsLatencyTrendRes struct {
	Models []string                       `json:"models"` // 模型名列表
	Dates  []string                       `json:"dates"`  // 日期列表
	Series map[string]*LatencyTrendSeries `json:"series"` // 按模型名索引
}

// 耗时趋势数据系列
type LatencyTrendSeries struct {
	AvgTotalTime []int64 `json:"avg_total_time"` // 每日平均总耗时(毫秒)
}
