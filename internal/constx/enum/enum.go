// Package enum 枚举
package enum

// StatusCode 状态码
type StatusCode uint16

const (
	// StatusCodeUndefined 未定义
	StatusCodeUndefined StatusCode = 0
	// StatusCodeCompleted 已完成
	StatusCodeCompleted StatusCode = 10000
)

// SourceCode 来源
type SourceCode int8

const (
	// SourceCodeConsume 销售
	SourceCodeConsume SourceCode = -1
	// SourceCodeUndefined 未定义
	SourceCodeUndefined SourceCode = 0
	// SourceCodeProduce 采购
	SourceCodeProduce SourceCode = 1
)
