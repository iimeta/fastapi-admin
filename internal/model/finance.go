package model

// 交易记录分页列表接口请求参数
type FinanceDealRecordPageReq struct {
	Paging
	UserId    int      `json:"user_id,omitempty"`    // 用户ID
	Remark    string   `json:"remark,omitempty"`     // 备注
	Status    int      `json:"status,omitempty"`     // 状态[1:正常, 2:退款, -1:删除]
	CreatedAt []string `json:"created_at,omitempty"` // 创建时间
}

// 交易记录分页列表接口响应参数
type FinanceDealRecordPageRes struct {
	Items  []*DealRecord `json:"items"`
	Paging *Paging       `json:"paging"`
}

// 交易记录
type DealRecord struct {
	Id        string `json:"id,omitempty"`         // ID
	UserId    int    `json:"user_id,omitempty"`    // 用户ID
	Quota     int    `json:"quota"`                // 充值额度
	Remark    string `json:"remark,omitempty"`     // 备注
	Status    int    `json:"status,omitempty"`     // 状态[1:正常, 2:退款, -1:删除]
	Creator   string `json:"creator,omitempty"`    // 创建人
	Updater   string `json:"updater,omitempty"`    // 更新人
	CreatedAt string `json:"created_at,omitempty"` // 创建时间
	UpdatedAt string `json:"updated_at,omitempty"` // 更新时间
}
