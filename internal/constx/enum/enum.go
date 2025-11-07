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

// AmountStatus 结算状态
type AmountStatus uint8

const (
	// AmountStatusUndefined 未定义
	AmountStatusUndefined AmountStatus = 0
	// AmountStatusUnpaid 未结算
	AmountStatusUnpaid AmountStatus = 1
	// AmountStatusPaying 结算中
	AmountStatusPaying AmountStatus = 2
	// AmountStatusPaid 已结算
	AmountStatusPaid AmountStatus = 3
)

// AccountDeal 账单交易
type AccountDeal int8

const (
	// AccountDealExpense 支出
	AccountDealExpense AccountDeal = -1
	// AccountDealUndefined 未定义
	AccountDealUndefined AccountDeal = 0
	// AccountDealIncome 收入
	AccountDealIncome AccountDeal = 1
)

// AccountStatus 账单状态
type AccountStatus uint8

const (
	// AccountStatusUndefined 未定义
	AccountStatusUndefined AccountStatus = iota
	// AccountStatusSubmitted 已提交
	AccountStatusSubmitted
	// AccountStatusApproved 已审核
	AccountStatusApproved
	// AccountStatusCompleted 已完成
	AccountStatusCompleted
	// AccountStatusCanceled 已取消
	AccountStatusCanceled
	// AccountStatusClosed 已关闭
	AccountStatusClosed
)
