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
	// SourceCodeInventoryUpdateConsume 编辑出库
	SourceCodeInventoryUpdateConsume SourceCode = -3
	// SourceCodeInventoryCreateConsume  手动出库
	SourceCodeInventoryCreateConsume SourceCode = -2
	// SourceCodeConsume 销售
	SourceCodeConsume SourceCode = -1
	// SourceCodeUndefined 未定义
	SourceCodeUndefined SourceCode = 0
	// SourceCodePurchaseProduce 采购入库
	SourceCodePurchaseProduce SourceCode = 1
	// SourceCodeInventoryCreateProduce 手动入库
	SourceCodeInventoryCreateProduce SourceCode = 2
	// SourceCodeInventoryUpdateProduce 编辑入库
	SourceCodeInventoryUpdateProduce SourceCode = 3
)
