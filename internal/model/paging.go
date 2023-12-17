package model

type Paging struct {
	Page     int64 `json:"current,omitempty"`  // 当前页
	PageSize int64 `json:"pageSize,omitempty"` // 每页条数
	Total    int64 `json:"total,omitempty"`    // 总条数
}
