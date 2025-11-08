// Package model 账单模型
package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// QueryBillBatch 账单批次查询条件
type QueryBillBatch struct {
	// ID 采购标识
	ID string `json:"id" where:"eq,id,omitempty"`
	// Name 名称
	Name string `json:"name" where:"regex,name,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// Status 状态码，精确匹配
	Status enum.StatusCode `json:"status" where:"eq,status,omitempty"`
	// Business 业务类型
	Business string `json:"business" where:"regex,business,omitempty"`
	// Category 交易类型：支出/收入
	Category enum.DealCategory `json:"category" where:"eq,category,omitempty"`
	// AmountStatus 结算状态
	AmountStatus enum.AmountStatus `json:"amount_status" where:"eq,amount_status,omitempty"`
	// AmountTotal 订单金额，分
	AmountTotal *coModel.RangeX[int32] `json:"amount_total" where:"range,amount_total,omitempty"`
	// AmountLeft 订单金额，分
	AmountLeft *coModel.RangeX[int32] `json:"amount_left" where:"range,amount_left,omitempty"`
	// AmountClear 已结算金额，分
	AmountClear *coModel.RangeX[int32] `json:"amount_clear" where:"range,amount_clear,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt *coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt *coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
	// ClearedAt 范围搜索
	ClearedAt *coModel.RangeX[time.Time] `json:"cleared_at" where:"range,cleared_at,omitempty"`
	// RefBatch 批次基本信息
	coModel.QueryRefBatch `json:",inline" where:",,omitempty"`
}

// SimpleBillBatch 账单批次基本信息
type SimpleBillBatch struct {
	// Xid 主键
	Xid uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	// ID 账单批次ID
	ID string `json:"id" xorm:"varchar(36) 'id'" cfpx:"id"`
	// Desc 账单描述
	Desc string `json:"desc" xorm:"varchar(200) 'desc'" cfpx:"desc"`
	// Business 业务类型
	Business string `json:"business" xorm:"varchar(100) 'business'"`
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
}

// BillBatch 账单
type BillBatch struct {
	// SimpleBillBatch 账单批次基本信息
	SimpleBillBatch `json:",inline" xorm:"extends"`
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
