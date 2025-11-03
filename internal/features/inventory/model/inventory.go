package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// QueryPurchase 商品查询参数
type QueryCommodity struct {
	coModel.QueryCommodity `json:"-,inline" where:",,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}

// FullCommodity 全量记录
type FullCommodity struct {
	Xid        uint32              `json:"xid" xorm:"pk autoincr 'xid'"`
	ID         string              `json:"id" xorm:"char(36) unique not null 'id'"`
	SourceCode enum.SourceCode     `json:"source_code" xorm:"tinyint 'source_code'"`
	CreatedAt  time.Time           `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt  time.Time           `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	BatchDesc  string              `json:"batch_desc" xorm:"varchar(500) 'batch_desc'" cfpx:"batch_desc"` // 批次描述
	BatchID    string              `json:"batch_id" xorm:"varchar(50) 'batch_id'"`                        // 批次标识
	Commodity  *coModel.Commodity  `json:"commodity" xorm:"extends 'commodity'" cfpx:"commodity"`         // 商品
	Company    *coModel.RefCompany `json:"company" xorm:"extends 'company'" cfpx:"company"`               // 企业
	Storage    string              `json:"storage" xorm:"varchar(100) 'storage'" cfpx:"storage"`          // 存储位置
}

// TableName 数据库表名
func (*FullCommodity) TableName() string {
	return featc.GetTableName(featc.InventoryInventory + "_full")
}

// GetFeature 特征
func (*FullCommodity) GetFeature() string {
	return featc.InventoryInventory
}

// LastCommodity 最新记录,采购/销售合并之后数据
type LastCommodity struct {
	Xid       uint32    `json:"xid" xorm:"pk autoincr 'xid'"`
	ID        string    `json:"id" xorm:"char(36) unique not null 'id'"`
	CreatedAt time.Time `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	// 商品
	coModel.Commodity `json:"commodity" xorm:"extends" cfpx:"commodity"`
}

// TableName 数据库表名
func (*LastCommodity) TableName() string {
	return featc.GetTableName(featc.InventoryInventory + "_last")
}

// GetFeature 特征
func (*LastCommodity) GetFeature() string {
	return featc.InventoryInventory
}
