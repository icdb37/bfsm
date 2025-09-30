package store

import "context"

// BuildFilter 构建过滤条件
type BuildFilter func() Filter
type UnmarshalWhere func(v any) Filter
type BuildTable func(v TableNamer) (Tabler, error)

// NewFilter 创建过滤条件
var NewFilter BuildFilter
var Unmarshal UnmarshalWhere
var NewTable BuildTable

// Filter 过滤条件接口
type Filter interface {
	Eq(key string, val any) Filter
	Ne(key string, val any) Filter
	In(key string, val any) Filter
	Nin(key string, val any) Filter
	Gt(key string, val any) Filter
	Gte(key string, val any) Filter
	Lt(key string, val any) Filter
	Lte(key string, val any) Filter
	Between(key string, vMin, vMax any) Filter
	Regex(key string, val any) Filter
	Or(v any) Filter
}

type Pager interface {
	// GetPageIndex 分页索引
	GetPageIndex() int
	// GetPageSize 分页大小
	GetPageSize() int
	// GetSorts 排序字段，-updated_at按updated_at降序，updated_at按updated_at升序
	GetSorts() []string
}

// Tabler 表接口
type Tabler interface {
	// Total 统计数量
	Total(ctx context.Context, f Filter) (int64, error)
	// Query 查询数据
	Search(ctx context.Context, f Filter, p Pager, v any) (int64, error)
	// Query 查询数据
	Query(ctx context.Context, f Filter, v any) error
	// Insert 插入数据
	Insert(ctx context.Context, vs ...any) error
	// Upsert 更新数据
	Upsert(ctx context.Context, w Filter, v any) error
	// Delete 删除数据
	Delete(ctx context.Context, w Filter) error
	// CreateIndex 创建索引
	CreateIndex(ctx context.Context, keys ...string) error
}

// Namer 表名
type TableNamer interface {
	// TableName 表名
	TableName() string
}

// Init 初始化
func Init(f BuildFilter, w UnmarshalWhere, t BuildTable) {
	NewFilter = f
	Unmarshal = w
	NewTable = t
}

// PageFilter 分页过滤条件
type PageFilter struct {
	PageIndex int      `json:"page_index"`
	PageSize  int      `json:"page_size"`
	Sorts     []string `json:"sorts"`
}

func (q *PageFilter) GetPageIndex() int {
	return q.PageIndex
}
func (q *PageFilter) GetPageSize() int {
	return q.PageSize
}
func (q *PageFilter) GetSorts() []string {
	return q.Sorts
}

// NewPageFilter 创建分页筛选器
func NewPageFilter() *PageFilter {
	return &PageFilter{}
}
