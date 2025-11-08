package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// BillDeal 账单交易详情
type BillDeal struct {
	// Xid 主键
	Xid uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	// ID 账单批次ID
	ID string `json:"id" xorm:"varchar(36) 'id'" cfpx:"id"`
	// Desc 账单描述
	Desc string `json:"desc" xorm:"varchar(200) 'desc'" cfpx:"desc"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" xorm:"created 'created_at'"`
	// UpdatedAt 更新时间
	UpdatedAt time.Time `json:"updated_at" xorm:"updated 'updated_at'"`
	// ClearedAt 结算时间
	ClearedAt time.Time `json:"cleared_at" xorm:"datetime 'cleared_at'"`
	// Category 交易类型：支出/收入
	Category enum.DealCategory `json:"category" xorm:"tinyint 'category'"`
	// RefBatch 引用账单批次
	coModel.RefBatch `json:",inline" xorm:"extends" cfpx:"ref_batch"`
	// RefDeal 引用账单交易详情
	coModel.RefDeal `json:",inline" xorm:"extends" cfpx:"ref_deal"`
}

// TableName 采购订单表名
func (*BillDeal) TableName() string {
	return featc.GetTableName(featc.BillDeal)
}

// GetFeature 获取功能名称
func (*BillDeal) GetFeature() string {
	return featc.BillDeal
}
