package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/utils"
)

// QueryProduceCommodity 商品查询参数
type QueryProduceCommodity struct {
	RefQueryCommodity `json:",inline" where:",,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}

// Normalize 归一化查询商品参数
func (q *QueryProduceCommodity) Normalize() {
	utils.PstrTrims(&q.Desc)
	q.RefQueryCommodity.Normalize()
}

// ProduceCommodity 库存商品
type ProduceCommodity struct {
	Xid          uint32                                                   `json:"xid" xorm:"pk autoincr 'xid'"`
	ID           string                                                   `json:"id" xorm:"char(36) unique not null 'id'"`
	SourceCode   enum.SourceCode                                          `json:"source_code" xorm:"tinyint 'source_code'"`
	CreatedAt    time.Time                                                `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt    time.Time                                                `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	BatchDesc    string                                                   `json:"batch_desc" xorm:"varchar(500) 'batch_desc'" cfpx:"batch_desc"` // 批次描述
	BatchID      string                                                   `json:"batch_id" xorm:"varchar(50) 'batch_id'"`                        // 批次标识
	Storage      string                                                   `json:"storage" xorm:"varchar(100) 'storage'" cfpx:"storage"`          // 存储位置
	UsedCount    int32                                                    `json:"used_count" xorm:"int 'used_count'" cfpx:"count"`               // 已用数量
	LeftCount    int32                                                    `json:"left_count" xorm:"int 'left_count'" cfpx:"count"`               // 剩余数量
	RefCommodity `json:",inline" xorm:"extends" cfpx:"commodity"`         // 商品
	RefCompany   `json:",inline" xorm:"extends 'company'" cfpx:"company"` // 企业
}

// TableName 数据库表名
func (*ProduceCommodity) TableName() string {
	return featc.GetTableName(featc.InventoryProduce)
}

// GetFeature 特征
func (*ProduceCommodity) GetFeature() string {
	return featc.InventoryProduce
}

func (f *ProduceCommodity) Normalize() {
	f.RefCommodity.Normalize()
	f.RefCompany.Normalize()
	utils.PstrTrims(&f.ID, &f.BatchDesc, &f.Storage)
}

// ProduceBatch -
type ProduceBatch struct {
	ID         string              `json:"id" xorm:"varchar(50) 'id'"`
	Desc       string              `json:"desc" xorm:"varchar(200) 'desc'" validate:"required"`
	Storage    string              `json:"storage" xorm:"varchar(100) 'storage'" cfpx:"storage"` // 存储位置
	Commodity  []*ProduceCommodity `json:"commodity" xorm:"json 'commodity'" cfpx:"commodity"`   //商品费用
	CreatedAt  time.Time           `json:"createdAt" xorm:"created 'created_at'"`
	UpdatedAt  time.Time           `json:"updatedAt" xorm:"updated 'updated_at'"`
	SourceCode enum.SourceCode     `json:"source_code" xorm:"tinyint 'source_code'"`
}

// Normalize -
func (b *ProduceBatch) Normalize() {
	utils.PstrTrims(&b.ID, &b.Desc, &b.Storage)
	for _, c := range b.Commodity {
		c.Normalize()
	}
}

// QueryProduceBatch -
type QueryProduceBatch struct {
	// ID 批次ID
	ID string `json:"id,omitempty" where:"regex,id,omitempty"`
	// Desc 备注
	Desc string `json:"desc,omitempty" where:"regex,desc,omitempty"`
	// Storage 存储位置
	Storage string `json:"storage,omitempty" where:"regex,storage,omitempty"`
	// CompanyName 公司名称
	CompanyName string `json:"company_name,omitempty" where:"regex,company_name,omitempty"`
	// CommodityName 商品名称
	CommodityName string `json:"commodity_name,omitempty" where:"regex,commodity,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}

// ConsumeCommodity -
type ConsumeCommodity struct {
	RefFullID      string `json:"ref_full_id" xorm:"varchar(36) 'ref_full_id'"`              // 引用生产批次ID
	RefBatchID     string `json:"ref_batch_id" xorm:"varchar(50) 'ref_batch_id'"`            // 引用生产批次标识
	CommodityCount int32  `json:"commodity_count" xorm:"int 'commodity_count'" cfpx:"count"` // 商品总量
	CommodityHash  string `json:"commodity_hash" xorm:"char(32) 'commodity_hash'"`
}

// TableName -
func (*ConsumeCommodity) TableName() string {
	return featc.GetTableName(featc.InventoryConsume)
}

// GetFeature -
func (*ConsumeCommodity) GetFeature() string {
	return featc.InventoryConsume
}

// Normalize -
func (b *ConsumeCommodity) Normalize() {

	utils.PstrTrims(&b.RefFullID, &b.RefBatchID)
}

// ConsumeBatch -
type ConsumeBatch struct {
	ID         string              `json:"id" xorm:"varchar(50) 'id'"`
	Desc       string              `json:"desc" xorm:"varchar(200) 'desc'" validate:"required"`
	Commodity  []*ConsumeCommodity `json:"commodity" xorm:"json 'commodity'" cfpx:"commodity"` //商品费用
	CreatedAt  time.Time           `json:"createdAt" xorm:"created 'created_at'"`
	UpdatedAt  time.Time           `json:"updatedAt" xorm:"updated 'updated_at'"`
	SourceCode enum.SourceCode     `json:"source_code" xorm:"tinyint 'source_code'"`
}

// Normalize -
func (b *ConsumeBatch) Normalize() {
	utils.PstrTrims(&b.ID, &b.Desc)
	for _, c := range b.Commodity {
		c.Normalize()
	}
}

// TableName 数据库表名
func (e *ConsumeBatch) TableName() string {
	return featc.GetTableName(featc.InventoryConsume)
}

// GetFeature 特征
func (e *ConsumeBatch) GetFeature() string {
	return featc.InventoryProduce
}

// UpdateCommodityCount 修改商品总量
type UpdateCommodityCount struct {
	UpdatedAt      time.Time `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	CommodityCount int32     `json:"commodity_count" xorm:"int 'commodity_count'" cfpx:"count"`
	UsedCount      int32     `json:"used_count" xorm:"int 'used_count'" cfpx:"count"` // 已用数量
	LeftCount      int32     `json:"left_count" xorm:"int 'left_count'" cfpx:"count"` // 剩余数量
}
