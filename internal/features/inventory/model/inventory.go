package model

import (
	"context"
	"time"

	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

// QueryLastCommodity 商品查询参数
type QueryLastCommodity struct {
	coModel.RefQueryCommodity `json:"-,inline" where:",,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}

// Normalize 归一化查询商品参数
func (q *QueryLastCommodity) Normalize() {
	utils.PstrTrims(&q.Desc)
	q.RefQueryCommodity.Normalize()
}

// LastCommodity 最新记录,采购/销售合并之后数据
type LastCommodity struct {
	Xid                  uint32    `json:"xid" xorm:"pk autoincr 'xid'"`
	ID                   string    `json:"id" xorm:"char(36) unique not null 'id'"  cfpx:"id"`
	CreatedAt            time.Time `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt            time.Time `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	UsedCount            int32     `json:"used_count" xorm:"int 'used_count'" cfpx:"count"` // 已用数量
	LeftCount            int32     `json:"left_count" xorm:"int 'left_count'" cfpx:"count"` // 剩余数量
	coModel.RefCommodity `json:",inline" xorm:"extends" cfpx:"commodity"`
}

// TableName 数据库表名
func (*LastCommodity) TableName() string {
	return featc.GetTableName(featc.InventoryInventoryLast)
}

// Normalize -
func (l *LastCommodity) Normalize() {
	utils.PstrTrims(&l.ID)
	l.RefCommodity.Normalize()
	l.CommodityHash = utils.Hash(l.CommodityName, l.CommoditySpec, l.CommoditySize)
}

// GetFeature 特征
func (*LastCommodity) GetFeature() string {
	return featc.InventoryInventory
}

// ProcessLastCommodity 处理最新商品
func ProcessLastCommodity(_ context.Context, l *LastCommodity) error {
	l.CommodityHash = utils.Hash(l.CommodityName, l.CommoditySpec, l.CommoditySize)
	return nil
}

// FullCommodity 全量商品
type FullCommodity coModel.ProduceCommodity

func (l *FullCommodity) Normalize() {
	l.RefCommodity.Normalize()
	l.RefCompany.Normalize()
}

// TableName 数据库表名
func (*FullCommodity) TableName() string {
	return featc.GetTableName(featc.InventoryInventoryFull)
}

// GetFeature 特征
func (*FullCommodity) GetFeature() string {
	return featc.InventoryInventory
}
