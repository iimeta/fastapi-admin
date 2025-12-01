package model

// 视频任务分页列表接口请求参数
type TaskVideoPageReq struct {
	Paging
	UserId    int      `json:"user_id,omitempty"`    // 用户ID
	AppId     int      `json:"app_id,omitempty"`     // 应用ID
	TraceId   string   `json:"trace_id,omitempty"`   // 日志ID
	TaskId    string   `json:"task_id,omitempty"`    // 任务ID
	Status    int      `json:"status,omitempty"`     // 状态
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
}

// 视频任务分页列表接口响应参数
type TaskVideoPageRes struct {
	Items  []*TaskVideo `json:"items"`
	Paging *Paging      `json:"paging"`
}

type TaskVideo struct {
	Id        string  `json:"id,omitempty"`         // ID
	TraceId   string  `json:"trace_id,omitempty"`   // 日志ID
	UserId    int     `json:"user_id,omitempty"`    // 用户ID
	AppId     int     `json:"app_id,omitempty"`     // 应用ID
	TaskId    string  `json:"task_id,omitempty"`    // 任务ID
	VideoUrl  string  `json:"video_url,omitempty"`  // 视频地址
	VideoTime float64 `json:"video_time,omitempty"` // 视频时长(秒)
	Status    int     `json:"status,omitempty"`     // 状态
	Creator   string  `json:"creator,omitempty"`    // 创建人
	Updater   string  `json:"updater,omitempty"`    // 更新人
	CreatedAt string  `json:"created_at,omitempty"` // 创建时间
	UpdatedAt string  `json:"updated_at,omitempty"` // 更新时间
}
