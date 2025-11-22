package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// PurchaseGoods 采购商品
type PurchaseGoods struct {
	// Xid 主键
	Xid uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" xorm:"created 'created_at'"`
	// PurchaseID 采购标识
	PurchaseID string `json:"purchase_id" xorm:"char(36) not null 'purchase_id'"`
	// PurchaseName 采购名称
	PurchaseName string `json:"purchase_name" xorm:"varchar(200) 'purchase_name'" cfpx:"name"`
	// Goods 商品信息
	coModel.RefGoods `json:",inline" xorm:"extends 'goods'" cfpx:"goods"`
}

// TableName 表名
func (*PurchaseGoods) TableName() string {
	return featc.GetTableName(featc.PurchaseGoods)
}

// GetFeature 特征
func (*PurchaseGoods) GetFeature() string {
	return featc.PurchaseGoods
}

// QueryPurchaseGoods 商品查询参数
type QueryPurchaseGoods struct {
	coModel.QueryRefGoods `json:",inline" where:",,omitempty"`
	// PurchaseName 采购名称，模糊匹配
	PurchaseName string `json:"purchase_name" where:"regex,purchase_name,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt *coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt *coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}
