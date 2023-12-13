package model

type Paging struct {
	Page      int64 `json:"page,omitempty"`       // 当前页
	PageSize  int64 `json:"page_size,omitempty"`  // 每页条数
	Total     int64 `json:"total,omitempty"`      // 总条数
	PageCount int64 `json:"page_count,omitempty"` // 总页数
}
