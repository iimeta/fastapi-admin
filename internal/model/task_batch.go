package model

// 批处理任务详情接口响应参数
type TaskBatchDetailRes struct {
	*TaskBatch
}

// 批处理任务分页列表接口请求参数
type TaskBatchPageReq struct {
	Paging
	UserId    int      `json:"user_id,omitempty"`    // 用户ID
	AppId     int      `json:"app_id,omitempty"`     // 应用ID
	TraceId   string   `json:"trace_id,omitempty"`   // 日志ID
	BatchId   string   `json:"batch_id,omitempty"`   // 批处理ID
	BatchUrl  string   `json:"batch_url,omitempty"`  // 批处理地址
	Status    string   `json:"status,omitempty"`     // 状态[queued:排队中, in_progress:进行中, completed:已完成, failed:已失败, expired:已过期, deleted:已删除]
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
}

// 批处理任务分页列表接口响应参数
type TaskBatchPageRes struct {
	Items  []*TaskBatch `json:"items"`
	Paging *Paging      `json:"paging"`
}

// 批处理任务详情复制字段值接口请求参数
type TaskBatchCopyFieldReq struct {
	Id    string `json:"id"`
	Field string `json:"field"`
}

// 批处理任务详情复制字段值接口响应参数
type TaskBatchCopyFieldRes struct {
	Value string `json:"value"`
}

type TaskBatch struct {
	Id           string         `json:"id,omitempty"`             // ID
	TraceId      string         `json:"trace_id,omitempty"`       // 日志ID
	UserId       int            `json:"user_id,omitempty"`        // 用户ID
	AppId        int            `json:"app_id,omitempty"`         // 应用ID
	Model        string         `json:"model,omitempty"`          // 模型
	BatchId      string         `json:"batch_id,omitempty"`       // 批处理ID
	InputFileId  string         `json:"input_file_id,omitempty"`  // 输入文件ID
	OutputFileId string         `json:"output_file_id,omitempty"` // 输出文件ID
	Status       string         `json:"status,omitempty"`         // 状态[validating:验证中, in_progress:进行中, finalizing:定稿中, completed:已完成, cancelling:取消中, cancelled:已取消, failed:已失败, expired:已过期, deleted:已删除]
	CompletedAt  string         `json:"completed_at,omitempty"`   // 完成时间
	ExpiresAt    string         `json:"expires_at,omitempty"`     // 过期时间
	ResponseData map[string]any `json:"response_data,omitempty"`  // 响应数据
	Creator      string         `json:"creator,omitempty"`        // 创建人
	Updater      string         `json:"updater,omitempty"`        // 更新人
	CreatedAt    string         `json:"created_at,omitempty"`     // 创建时间
	UpdatedAt    string         `json:"updated_at,omitempty"`     // 更新时间
}
