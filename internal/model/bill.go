package model

import (
	"github.com/icdb37/bfsm/internal/constx/enum"
)

// RefDeal 交易基本信息
type RefDeal struct {
	// DealDesc 校验描述
	DealDesc string `json:"deal_desc" xorm:"varchar(200) 'deal_desc'" cfpx:"desc"`
	// AmountStatus 结算状态
	AmountStatus enum.AmountStatus `json:"amount_status" xorm:"tinyint 'amount_status'"`
	// AmountTotal 交易金额
	AmountTotal int32 `json:"amount_total" xorm:"int 'amount_total'"`
	// AmountClear 已结算金额
	AmountClear int32 `json:"amount_clear" xorm:"int 'amount_clear'"`
	// AmountLeft 未结算金额
	AmountLeft int32 `json:"amount_left" xorm:"int 'amount_left'"`
	// AmountDesc 结算描述
	AmountDesc string `json:"amount_desc" xorm:"varchar(200) 'amount_desc'" cfpx:"desc"`
	// RefCompany 企业基本信息
	RefCompany `json:"ref_company" xorm:"extends"`
}

// Normalize -
func (r *RefDeal) Normalize() {
	r.AmountLeft = r.AmountTotal - r.AmountClear
	if r.AmountLeft <= 0 {
		r.AmountLeft = 0
		r.AmountStatus = enum.AmountStatusPaid
	}
	if r.AmountClear > 0 {
		r.AmountStatus = enum.AmountStatusPaying
	} else {
		r.AmountStatus = enum.AmountStatusUnpaid
	}
	r.RefCompany.Normalize()
}

// BatchDeal 批次交易
type BatchDeal struct {
	// Category 交易类型：支出/收入
	Category enum.DealCategory `json:"category" xorm:"tinyint 'category'"`
	// RefBatch 批次信息
	RefBatch `json:",inline" xorm:"extends"`
	// Datas 商品列表
	Datas []*RefDeal `json:"datas" xorm:"json 'datas'"`
}

// Normalize -
func (b *BatchDeal) Normalize() {
	b.RefBatch.Normalize()
	for _, data := range b.Datas {
		data.Normalize()
	}
}
