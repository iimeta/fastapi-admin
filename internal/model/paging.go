package model

type Paging struct {
	Page      int64 // 当前页
	PageSize  int64 // 每页条数
	Total     int64 // 总条数
	PageCount int64 // 总页数
	StartNums int64 // 起始条数
	EndNums   int64 // 结束条数
}
