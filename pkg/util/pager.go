package util

// InitPageNo 初始化pageNo
func InitPageNo(pageNo int) int {
	if pageNo < 1 {
		pageNo = 1
	}
	return pageNo
}

// InitPageSize 初始化pageSize
func InitPageSize(pageSize int) int {
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 20 {
		pageSize = 20
	}
	return pageSize
}
