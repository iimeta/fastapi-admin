package model

import smodel "github.com/iimeta/fastapi-sdk/v2/model"

// 绘图任务详情接口响应参数
type TaskImageDetailRes struct {
	*TaskImage
}

// 绘图任务分页列表接口请求参数
type TaskImagePageReq struct {
	Paging
	UserId    int      `json:"user_id,omitempty"`    // 用户ID
	AppId     int      `json:"app_id,omitempty"`     // 应用ID
	TraceId   string   `json:"trace_id,omitempty"`   // 日志ID
	ImageId   string   `json:"image_id,omitempty"`   // 图像ID
	ImageUrl  string   `json:"image_url,omitempty"`  // 图像地址
	Status    string   `json:"status,omitempty"`     // 状态[queued:排队中, in_progress:进行中, completed:已完成, failed:已失败, expired:已过期, deleted:已删除]
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
}

// 绘图任务分页列表接口响应参数
type TaskImagePageRes struct {
	Items  []*TaskImage `json:"items"`
	Paging *Paging      `json:"paging"`
}

// 绘图任务批量操作接口请求参数
type TaskImageBatchOperateReq struct {
	Action string   `json:"action"` // 动作
	Ids    []string `json:"ids"`    // 主键Ids
	Value  any      `json:"value"`  // 值
}

// 绘图任务详情复制字段值接口请求参数
type TaskImageCopyFieldReq struct {
	Id    string `json:"id"`
	Field string `json:"field"`
}

// 绘图任务详情复制字段值接口响应参数
type TaskImageCopyFieldRes struct {
	Value string `json:"value"`
}

type TaskImage struct {
	Id             string             `json:"id,omitempty"`               // ID
	TraceId        string             `json:"trace_id,omitempty"`         // 日志ID
	UserId         int                `json:"user_id,omitempty"`          // 用户ID
	AppId          int                `json:"app_id,omitempty"`           // 应用ID
	Model          string             `json:"model,omitempty"`            // 模型
	Action         string             `json:"action,omitempty"`           // 接口
	ImageId        string             `json:"image_id,omitempty"`         // 图像ID
	JobId          string             `json:"job_id,omitempty"`           // 上游任务ID
	Width          int                `json:"width,omitempty"`            // 宽度
	Height         int                `json:"height,omitempty"`           // 高度
	N              int                `json:"n,omitempty"`                // 生成数量
	Quality        string             `json:"quality,omitempty"`          // 质量
	Size           string             `json:"size,omitempty"`             // 尺寸大小
	OutputFormat   string             `json:"output_format,omitempty"`    // 输出格式
	ResponseFormat string             `json:"response_format,omitempty"`  // 响应格式
	Prompt         string             `json:"prompt,omitempty"`           // 提示
	Progress       int                `json:"progress,omitempty"`         // 进度
	Status         string             `json:"status,omitempty"`           // 状态[queued:排队中, in_progress:进行中, completed:已完成, failed:已失败, expired:已过期, deleted:已删除]
	CompletedAt    string             `json:"completed_at,omitempty"`     // 完成时间
	ExpiresAt      string             `json:"expires_at,omitempty"`       // 过期时间
	ImageUrl       string             `json:"image_url,omitempty"`        // 图像地址
	ImageUrls      []string           `json:"image_urls,omitempty"`       // 图像地址列表
	FileName       string             `json:"file_name,omitempty"`        // 文件名
	FileNames      []string           `json:"file_names,omitempty"`       // 文件名列表
	FilePath       string             `json:"file_path,omitempty"`        // 文件路径
	FilePaths      []string           `json:"file_paths,omitempty"`       // 文件路径列表
	InputFilePaths []string           `json:"input_file_paths,omitempty"` // 输入文件路径列表(异步任务base64转储)
	ResponseData   map[string]any     `json:"response_data,omitempty"`    // 响应数据
	Error          *smodel.ImageError `json:"error,omitempty"`            // 错误信息
	Creator        string             `json:"creator,omitempty"`          // 创建人
	Updater        string             `json:"updater,omitempty"`          // 更新人
	CreatedAt      string             `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt      string             `json:"updated_at,omitempty"`       // 更新时间
}
