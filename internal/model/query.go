package model

// RangeX 范围查询
type RangeX[T comparable] struct {
	Beg T `json:"beg" where:"gte,,omitempty"`
	End T `json:"end" where:"lte,,omitempty"`
}

// QueryPage 查询分页
type QueryPage struct {
	PageIndex uint32   `json:"index"`
	PageSize  uint32   `json:"size"`
	Sorts     []string `json:"sorts"`
}

// GetPageIndex 分页索引
func (q *QueryPage) GetPageIndex() int {
	return int(q.PageIndex)
}

// GetPageSize 分页大小
func (q *QueryPage) GetPageSize() int {
	return int(q.PageSize)
}

// GetSorts 排序字段，-updated_at按updated_at降序，updated_at按updated_at升序
func (q *QueryPage) GetSorts() []string {
	return q.Sorts
}

// SearchRequest 搜索请求
type SearchRequest[T any] struct {
	QueryPage `json:",inline"`
	Query     *T `json:"query"`
}

func (q *SearchRequest[T]) GetPage() *QueryPage {
	return &q.QueryPage
}

// SearchResponse 搜索应答
type SearchResponse[T any] struct {
	Total int64 `json:"total"`
	Datas []*T  `json:"datas"`
}

// IDResponse ID应答
type IDResponse[T any] struct {
	ID T `json:"id"`
}

// NewIDResponse 创建ID应答
func NewIDResponse[T any](id T) *IDResponse[T] {
	return &IDResponse[T]{ID: id}
}
