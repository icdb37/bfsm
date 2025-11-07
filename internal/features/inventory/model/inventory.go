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
	coModel.QueryCommodity `json:"-,inline" where:",,omitempty"`
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
	q.QueryCommodity.Normalize()
}

// LastCommodity 最新记录,采购/销售合并之后数据
type LastCommodity struct {
	// Xid 主键
	Xid uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	// LastID 标识
	LastID string `json:"last_id" xorm:"char(36) unique not null 'last_id'"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" xorm:"created 'created_at'"`
	// UpdatedAt 修改时间
	UpdatedAt time.Time `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	// UsedCount 已用数量
	UsedCount int32 `json:"used_count" xorm:"int 'used_count'" cfpx:"count"`
	// LeftCount 剩余数量
	LeftCount int32 `json:"left_count" xorm:"int 'left_count'" cfpx:"count"`
	// Count 总量
	Count int32 `json:"count" xorm:"int 'count'" cfpx:"count"`
	// Commodity 商品信息
	coModel.Commodity `json:",inline" xorm:"extends" cfpx:"commodity"`
}

// TableName 数据库表名
func (*LastCommodity) TableName() string {
	return featc.GetTableName(featc.InventoryGoodsLast)
}

// Normalize -
func (l *LastCommodity) Normalize() {
	utils.PstrTrims(&l.LastID)
	l.Commodity.Normalize()
}

// GetFeature 特征
func (*LastCommodity) GetFeature() string {
	return featc.InventoryInventory
}

// ProcessLastCommodity 处理最新商品
func ProcessLastCommodity(_ context.Context, l *LastCommodity) error {
	l.Hash = l.GetHash()
	return nil
}

// QueryFullGoods 商品查询参数
type QueryFullGoods struct {
	coModel.QueryCommodity `json:"-,inline" where:",,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}

// FullGoods 全量商品
type FullGoods struct {
	Xid              uint32    `json:"xid" xorm:"pk autoincr 'xid'"`
	CreatedAt        time.Time `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt        time.Time `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	Storage          string    `json:"storage" xorm:"varchar(100) 'storage'" cfpx:"storage"` // 存储位置
	UsedCount        int32     `json:"used_count" xorm:"int 'used_count'" cfpx:"count"`      // 已用数量
	LeftCount        int32     `json:"left_count" xorm:"int 'left_count'" cfpx:"count"`      // 剩余数量
	coModel.RefBatch `json:",inline" xorm:"extends"`
	coModel.RefGoods `json:",inline" xorm:"extends"`
}

func (l *FullGoods) Normalize() {
	utils.PstrTrims(&l.Storage)
	l.RefGoods.Normalize()
}

// TableName 数据库表名
func (*FullGoods) TableName() string {
	return featc.GetTableName(featc.InventoryGoodsFull)
}

// GetFeature 特征
func (*FullGoods) GetFeature() string {
	return featc.InventoryInventory
}
