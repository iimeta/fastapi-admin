package model

import smodel "github.com/iimeta/fastapi-sdk/model"

// 视频任务详情接口响应参数
type TaskVideoDetailRes struct {
	*TaskVideo
}

// 视频任务分页列表接口请求参数
type TaskVideoPageReq struct {
	Paging
	UserId    int      `json:"user_id,omitempty"`    // 用户ID
	AppId     int      `json:"app_id,omitempty"`     // 应用ID
	TraceId   string   `json:"trace_id,omitempty"`   // 日志ID
	VideoId   string   `json:"video_id,omitempty"`   // 视频ID
	VideoUrl  string   `json:"video_url,omitempty"`  // 视频地址
	Status    string   `json:"status,omitempty"`     // 状态[queued:排队中, in_progress:进行中, completed:已完成, failed:已失败, expired:已过期]
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
}

// 视频任务分页列表接口响应参数
type TaskVideoPageRes struct {
	Items  []*TaskVideo `json:"items"`
	Paging *Paging      `json:"paging"`
}

type TaskVideo struct {
	Id                 string             `json:"id,omitempty"`                    // ID
	TraceId            string             `json:"trace_id,omitempty"`              // 日志ID
	UserId             int                `json:"user_id,omitempty"`               // 用户ID
	AppId              int                `json:"app_id,omitempty"`                // 应用ID
	Model              string             `json:"model,omitempty"`                 // 模型
	VideoId            string             `json:"video_id,omitempty"`              // 视频ID
	Width              int                `json:"width,omitempty"`                 // 宽度
	Height             int                `json:"height,omitempty"`                // 高度
	Seconds            int                `json:"seconds,omitempty"`               // 秒数
	Prompt             string             `json:"prompt,omitempty"`                // 提示
	Progress           int                `json:"progress,omitempty"`              // 进度
	RemixedFromVideoId string             `json:"remixed_from_video_id,omitempty"` // 混合ID
	Status             string             `json:"status,omitempty"`                // 状态[queued:排队中, in_progress:进行中, completed:已完成, failed:已失败, expired:已过期]
	CompletedAt        string             `json:"completed_at,omitempty"`          // 完成时间
	ExpiresAt          string             `json:"expires_at,omitempty"`            // 过期时间
	VideoUrl           string             `json:"video_url,omitempty"`             // 视频地址
	FileName           string             `json:"file_name,omitempty"`             // 文件名
	FilePath           string             `json:"file_path,omitempty"`             // 文件路径
	Error              *smodel.VideoError `json:"error,omitempty"`                 // 错误信息
	Creator            string             `json:"creator,omitempty"`               // 创建人
	Updater            string             `json:"updater,omitempty"`               // 更新人
	CreatedAt          string             `json:"created_at,omitempty"`            // 创建时间
	UpdatedAt          string             `json:"updated_at,omitempty"`            // 更新时间
}
