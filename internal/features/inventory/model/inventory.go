package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// EntireCommodity 库存商品
type EntireCommodity struct {
	Xid         uint32                 `json:"xid" xorm:"pk autoincr 'xid'"`
	CreatedAt   time.Time              `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt   time.Time              `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	ID          string                 `json:"id" xorm:"char(36) unique not null 'id'"`
	Desc        string                 `json:"desc" xorm:"varchar(500) 'desc'" cfpx:"desc"`
	BatchID     string                 `json:"batch_id" xorm:"varchar(50) 'batch_id'"`                 // 批次标识
	Commodities *coModel.Commodity     `json:"commodities" xorm:"json 'commodities'" cfpx:"commodity"` // 商品
	Company     *coModel.SimpleCompany `json:"company" xorm:"extends 'company'" cfpx:"company"`        // 企业
	Storage     string                 `json:"storage" xorm:"varchar(100) 'storage'" cfpx:"storage"`   // 存储位置
}

// TableName 数据库表名
func (e *EntireCommodity) TableName() string {
	return featc.InventoryInventory
}

// GetFeature 特征
func (e *EntireCommodity) GetFeature() string {
	return featc.InventoryInventory
}

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

// AggregateCommodity 商品聚合查询参数
type AggregateCommodity struct {
}

func (a *AggregateCommodity) GetxWhere() (any, []any) {
	where := "select id,name,spec,size,sum(count),max(updated_at) from %s group by id,name,spec,size"
	return where, nil
}
