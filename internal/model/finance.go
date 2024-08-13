package model

// 账单明细分页列表接口请求参数
type FinanceBillPageReq struct {
	Paging
	UserId   int      `json:"user_id,omitempty"`   // 用户ID
	StatDate []string `json:"stat_date,omitempty"` // 统计时间
}

// 账单明细分页列表接口响应参数
type FinanceBillPageRes struct {
	Items  []*Bill `json:"items"`
	Paging *Paging `json:"paging"`
}

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

// 账单明细
type Bill struct {
	Id       string `json:"id"`        // ID
	UserId   int    `json:"user_id"`   // 用户ID
	Total    int    `json:"total"`     // 总数
	Tokens   int    `json:"tokens"`    // 令牌数
	Models   int    `json:"models"`    // 模型数
	StatDate string `json:"stat_date"` // 统计日期
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