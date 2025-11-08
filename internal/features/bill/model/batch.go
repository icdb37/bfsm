// Package model 账单模型
package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// BillBatch 账单
type BillBatch struct {
	// Xid 主键
	Xid uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	// ID 账单批次ID
	ID string `json:"id" xorm:"varchar(36) 'id'" cfpx:"id"`
	// Desc 账单描述
	Desc string `json:"desc" xorm:"varchar(200) 'desc'" cfpx:"desc"`
	// Status 状态
	Status enum.StatusCode `json:"status" xorm:"tinyint 'status'"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" xorm:"created 'created_at'"`
	// UpdatedAt 更新时间
	UpdatedAt time.Time `json:"updated_at" xorm:"updated 'updated_at'"`
	// ClearedAt 结算时间
	ClearedAt time.Time `json:"cleared_at" xorm:"datetime 'cleared_at'"`
	// Category 交易类型：支出/收入
	Category enum.DealCategory `json:"category" xorm:"tinyint 'category'"`
	// AmountTotal 总金额
	AmountTotal int32 `json:"amount_total" xorm:"int 'amount_total'"`
	// AmountLeft 未结算金额
	AmountLeft int32 `json:"amount_left" xorm:"int 'amount_left'"`
	// AmountClear 已结算金额
	AmountClear int32 `json:"amount_clear" xorm:"int 'amount_clear'"`
	// AmountStatus 结算状态
	AmountStatus enum.AmountStatus `json:"amount_status" xorm:"tinyint 'amount_status'"`
	// AmountDesc 结算描述
	AmountDesc string `json:"amount_desc" xorm:"varchar(200) 'amount_desc'" cfpx:"desc"`
	// RefBatch 批次基本信息
	coModel.RefBatch `json:",inline" xorm:"extends"`
	// Datas 交易详情
	Datas []*coModel.RefDeal `json:"datas" xorm:"json 'datas'"`
}

// TableName 采购订单表名
func (*BillBatch) TableName() string {
	return featc.GetTableName(featc.BillBatch)
}

// GetFeature 获取功能名称
func (*BillBatch) GetFeature() string {
	return featc.BillBatch
}

// Normalize -
func (a *BillBatch) Normalize() {
	a.RefBatch.Normalize()
	a.AmountLeft = 0
	a.AmountClear = 0
	a.UpdatedAt = time.Now()
	a.CreatedAt = a.UpdatedAt
	a.Status = enum.StatusCodeSubmitted
	for _, d := range a.Datas {
		d.Normalize()
		a.AmountTotal += d.AmountTotal
		a.AmountClear += d.AmountClear
	}
	a.AmountLeft = a.AmountTotal - a.AmountClear
	switch {
	case a.AmountClear == 0:
		a.AmountStatus = enum.AmountStatusUndefined
		a.AmountLeft = a.AmountTotal
		a.Desc = "待付款"
	case a.AmountClear >= a.AmountTotal:
		a.AmountStatus = enum.AmountStatusPaid
		a.ClearedAt = a.UpdatedAt
		a.AmountLeft = 0
		a.Desc = "钱货两清"
		a.Status = enum.StatusCodeCompleted
	default:
		a.AmountStatus = enum.AmountStatusPaying
		a.AmountLeft = a.AmountTotal - a.AmountClear
		a.Desc = "已付定金"
	}
}
