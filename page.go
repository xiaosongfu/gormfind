package gormfind

type Page struct {
	// 第几页
	Page int
	// 每页几条记录
	Size int
	// 排序字段
	SortField *string
	// 排序顺序(asc desc)
	Order *string
}
