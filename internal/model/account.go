package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
)

// 账单

type Account struct {
	// Xid 主键
	Xid uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	// ID 商品ID
	ID string `json:"id" xorm:"char(36) unique not null 'id'"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" xorm:"created 'created_at'"`
	// UpdatedAt 更新时间
	UpdatedAt time.Time `json:"updated_at" xorm:"updated 'updated_at'"`
	// ClearedAt 结算时间
	ClearedAt time.Time `json:"cleared_at" xorm:"datetime 'cleared_at'"`
	// Desc 描述
	Desc string `json:"desc" xorm:"varchar(200) 'desc'"`
	// DealID 交易ID
	DealID string `json:"deal_id" xorm:"char(36) 'deal_id'"`
	// DealCategory 交易类型：支出/收入
	DealCategory enum.AccountDeal `json:"deal_category" xorm:"tinyint 'deal_category'"`
	// CompanyID 企业ID
	CompanyID string `json:"company_id" xorm:"char(36) 'company_id'"`
	// TotalAmount 总金额
	TotalAmount int32 `json:"total_amount" xorm:"int 'total_amount'"`
	// ClearAmount 已结算金额
	ClearAmount int32 `json:"clear_amount" xorm:"int 'clear_amount'"`
	// Status 状态
	Status enum.StatusCode `json:"status" xorm:"tinyint 'status'"`
	// AmountStatus 结算状态
	AmountStatus enum.AmountStatus `json:"amount_status" xorm:"tinyint 'amount_status'"`
	// AmountDesc 结算描述
	AmountDesc string `json:"amount_desc" xorm:"varchar(200) 'amount_desc'" cfpx:"desc"`
}
