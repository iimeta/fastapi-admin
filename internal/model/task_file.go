package model

import (
	serrors "github.com/iimeta/fastapi-sdk/errors"
)

// 文件任务详情接口响应参数
type TaskFileDetailRes struct {
	*TaskFile
}

// 文件任务分页列表接口请求参数
type TaskFilePageReq struct {
	Paging
	UserId    int      `json:"user_id,omitempty"`    // 用户ID
	AppId     int      `json:"app_id,omitempty"`     // 应用ID
	TraceId   string   `json:"trace_id,omitempty"`   // 日志ID
	FileId    string   `json:"file_id,omitempty"`    // 文件ID
	FileUrl   string   `json:"file_url,omitempty"`   // 文件地址
	Status    string   `json:"status,omitempty"`     // 状态[queued:排队中, in_progress:进行中, completed:已完成, failed:已失败, expired:已过期, deleted:已删除]
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
}

// 文件任务分页列表接口响应参数
type TaskFilePageRes struct {
	Items  []*TaskFile `json:"items"`
	Paging *Paging     `json:"paging"`
}

// 文件任务详情复制字段值接口请求参数
type TaskFileCopyFieldReq struct {
	Id    string `json:"id"`
	Field string `json:"field"`
}

// 文件任务详情复制字段值接口响应参数
type TaskFileCopyFieldRes struct {
	Value string `json:"value"`
}

type TaskFile struct {
	Id           string            `json:"id,omitempty"`             // ID
	TraceId      string            `json:"trace_id,omitempty"`       // 日志ID
	UserId       int               `json:"user_id,omitempty"`        // 用户ID
	AppId        int               `json:"app_id,omitempty"`         // 应用ID
	Model        string            `json:"model,omitempty"`          // 模型
	Purpose      string            `json:"purpose,omitempty"`        // 用途[assistants, assistants_output, batch, batch_output, fine-tune, fine-tune-results, vision, user_data]
	FileId       string            `json:"file_id,omitempty"`        // 文件ID
	FileName     string            `json:"file_name,omitempty"`      // 文件名
	Bytes        int               `json:"bytes,omitempty"`          // 文件大小
	Status       string            `json:"status,omitempty"`         // 状态[uploaded:已上传, processed:已处理, error:已失败, expired:已过期, deleted:已删除]
	ExpiresAt    string            `json:"expires_at,omitempty"`     // 过期时间
	FileUrl      string            `json:"file_url,omitempty"`       // 文件地址
	FilePath     string            `json:"file_path,omitempty"`      // 文件路径
	Error        *serrors.ApiError `json:"error,omitempty"`          // 错误信息
	BatchTraceId string            `json:"batch_trace_id,omitempty"` // 批处理日志ID
	Creator      string            `json:"creator,omitempty"`        // 创建人
	Updater      string            `json:"updater,omitempty"`        // 更新人
	CreatedAt    string            `json:"created_at,omitempty"`     // 创建时间
	UpdatedAt    string            `json:"updated_at,omitempty"`     // 更新时间
}
