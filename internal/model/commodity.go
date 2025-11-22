package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/utils"
)

// CommodityAttr 商品属性
type CommodityAttr struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Normalize -
func (a *CommodityAttr) Normalize() {
	utils.PstrTrims(&a.Name, &a.Value)
}

// Commodity 商品
type Commodity struct {
	Hash     string           `json:"hash" xorm:"char(32) 'hash'"`
	Name     string           `json:"name" xorm:"varchar(100) 'name'" cfpx:"name"`
	Desc     string           `json:"desc" xorm:"varchar(200) 'desc'" cfpx:"desc"`
	Spec     string           `json:"spec" xorm:"varchar(100) 'spec'" cfpx:"spec"`
	Size     string           `json:"size" xorm:"varchar(100) 'size'" cfpx:"size"`
	Attrs    []*CommodityAttr `json:"attrs" xorm:"json 'attrs'"`
	Validity int32            `json:"validity" xorm:"tinyint 'validity'" cfpx:"validity"`
	Price    int32            `json:"price" xorm:"int 'price'" cfpx:"price"`
}

func (c *Commodity) GetHash() string {
	return utils.Hash(c.Name, c.Spec, c.Size)
}

func (c *Commodity) Normalize() {
	utils.PstrTrims(&c.Name, &c.Desc, &c.Spec, &c.Size)
	for _, attr := range c.Attrs {
		utils.PstrTrims(&attr.Name, &attr.Value)
	}
	c.Hash = c.GetHash()
}

// QueryCommodity - 查询商品
type QueryCommodity struct {
	// Name 姓名
	Name string `json:"name,omitempty" where:"regex,name,omitempty"`
	// Desc 备注
	Desc string `json:"desc,omitempty" where:"regex,desc,omitempty"`
	// Spec 规格
	Spec string `json:"spec,omitempty" where:"regex,spec,omitempty"`
	// Size 尺寸
	Size string `json:"size,omitempty" where:"regex,size,omitempty"`
	// Count 数量
	Count *RangeX[uint32] `json:"count,omitempty" where:"range,count,omitempty"`
	// Hash 商品哈希值
	Hash string `json:"hash,omitempty" where:"eq,hash,omitempty"`
}

// Normalize 归一化查询商品参数
func (q *QueryCommodity) Normalize() {
	utils.PstrTrims(&q.Name, &q.Desc, &q.Spec, &q.Size, &q.Hash)
}

// Goods 货物
type Goods struct {
	// Commodity 商品信息
	Commodity `json:",inline" xorm:"extends" cfpx:"commodity"`
	// Count 数量
	Count int32 `json:"count" xorm:"int 'count'" cfpx:"count"`
	// Amount 商品金额，分
	Amount int32 `json:"amount" xorm:"int 'amount'" cfpx:"amount"`
	// ProducedAt 生产时间
	ProducedAt time.Time `json:"produced_at" xorm:"datetime 'produced_at'" cfpx:"produced_at"`
	// ExpiredAt 过期时间
	ExpiredAt time.Time `json:"expired_at" xorm:"datetime 'expired_at'" cfpx:"expired_at"`
	// Storeage 存储位置
	Storeage string `json:"storeage" xorm:"varchar(100) 'storeage'" cfpx:"storeage"`
}

// QueryRefGoods 查询商品
type QueryRefGoods struct {
	// QueryCommodity 查询商品
	QueryCommodity `json:",inline" where:",,omitempty"`
	// RefCompany 引用企业
	QueryRefCompany `json:",inline" where:",,omitempty"`
	// Count 数量
	Count *RangeX[int32] `json:"count,omitempty" where:"range,count,omitempty"`
	// Amount 商品金额，分
	Amount *RangeX[int64] `json:"amount,omitempty" where:"range,amount,omitempty"`
	// ProducedAt 生产时间
	ProducedAt *RangeX[time.Time] `json:"produced_at,omitempty" where:"range,produced_at,omitempty"`
	// ExpiredAt 过期时间
	ExpiredAt *RangeX[time.Time] `json:"expired_at,omitempty" where:"range,expired_at,omitempty"`
}

func (q *QueryRefGoods) Normalize() {
	q.QueryCommodity.Normalize()
	q.QueryRefCompany.Normalize()
}

// RefGoods 引用商品信息
type RefGoods struct {
	// ID 商品ID
	ID string `json:"id" xorm:"char(36) unique not null 'id'"`
	// Goods 商品信息
	Goods `json:",inline" xorm:"extends"`
	// RefCompany 引用公司信息
	RefCompany `json:",inline" xorm:"extends"`
}

func (r *RefGoods) Normalize() {
	utils.PstrTrims(&r.ID)
	r.Goods.Normalize()
	r.RefCompany.Normalize()
}

// BatchGoods 批次商品
type BatchGoods struct {
	// Datas 商品列表
	Datas []*RefGoods `json:"datas" xorm:"json 'datas'"`
	// RefBatch 批次信息
	RefBatch `json:",inline" xorm:"extends"`
}

func (l *BatchGoods) Normalize() {
	l.RefBatch.Normalize()
	for _, data := range l.Datas {
		data.Normalize()
	}
}

// GoodsCount 商品总量
type GoodsCount struct {
	UpdatedAt time.Time `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	Count     int32     `json:"count" xorm:"int 'count'" cfpx:"count"`
	UsedCount int32     `json:"used_count" xorm:"int 'used_count'" cfpx:"count"` // 已用数量
	LeftCount int32     `json:"left_count" xorm:"int 'left_count'" cfpx:"count"` // 剩余数量
}

// QueryRefBatch 查询批次
type QueryRefBatch struct {
	// BatchID 批次ID
	BatchID string `json:"batch_id,omitempty" where:"eq,batch_id,omitempty"`
	// BatchName 批次名称
	BatchName string `json:"batch_name,omitempty" where:"regex,batch_name,omitempty"`
	// BatchDesc 批次描述
	BatchDesc string `json:"batch_desc,omitempty" where:"regex,batch_desc,omitempty"`
	// SourceCode 来源
	SourceCode enum.SourceCode `json:"source_code,omitempty" where:"eq,source_code,omitempty"`
}

// Normalize -
func (q *QueryRefBatch) Normalize() {
	utils.PstrTrims(&q.BatchID, &q.BatchName, &q.BatchDesc)
}

// RefBatch 引用批次
type RefBatch struct {
	// BatchID 批次ID
	BatchID string `json:"batch_id" xorm:"char(36) 'batch_id'"`
	// BatchName 批次名称
	BatchName string `json:"batch_name" xorm:"varchar(100) 'batch_name'"`
	// BatchDesc 批次描述
	BatchDesc string `json:"batch_desc" xorm:"varchar(200) 'batch_desc'"`
	// SourceCode 来源
	SourceCode enum.SourceCode `json:"source_code" xorm:"tinyint 'source_code'" cfpx:"source_code"`
}

func (l *RefBatch) Normalize() {
	utils.PstrTrims(&l.BatchID, &l.BatchName, &l.BatchDesc)
}
