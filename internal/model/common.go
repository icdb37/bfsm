package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/utils"
)

// IdentityCard 身份证信息
type IdentityCard struct {
	No       string    `json:"no" xorm:"varchar(18) 'no'" validate:"required"`
	Name     string    `json:"name" xorm:"varchar(50) 'name'" validate:"required"`
	Sex      uint8     `json:"sex" xorm:"tinyint 'sex'"`
	Birth    time.Time `json:"birth" xorm:"date 'birth'"`
	Address  string    `json:"address" xorm:"varchar(200) 'address'"`
	ValidBeg time.Time `json:"validBeg" xorm:"date 'valid_beg'"`
	ValidEnd time.Time `json:"validEnd" xorm:"date 'valid_end'"`
}

func (i *IdentityCard) Normalize() {
	utils.PstrTrims(&i.Name, &i.Address)
}

// BankCard 银行卡信息
type BankCard struct {
	ID        int64     `json:"-" xorm:"pk autoincr 'id'"`
	BuilderID int64     `json:"-" xorm:"'builder_id' index"`
	No        string    `json:"no" xorm:"varchar(30) 'no'" validate:"required"`
	Name      string    `json:"name" xorm:"varchar(50) 'name'" validate:"required"`
	Bank      string    `json:"bank" xorm:"varchar(100) 'bank'" validate:"required"`
	ValidBeg  time.Time `json:"validBeg" xorm:"date 'valid_beg'"`
	ValidEnd  time.Time `json:"validEnd" xorm:"date 'valid_end'"`
}

func (b *BankCard) Normalize() {
	utils.PstrTrims(&b.No, &b.Name, &b.Bank)
}

// Contact 联系人
type Contact struct {
	Name  string `json:"name" xorm:"varchar(50) 'name'" validate:"required" cfpx:"name"`
	Phone string `json:"phone" xorm:"varchar(20) 'phone'" validate:"required" cfpx:"phone"`
	Desc  string `json:"desc,omitempty" xorm:"varchar(100) 'desc'"`
}

func (c *Contact) Normalize() {
	utils.PstrTrims(&c.Name, &c.Phone, &c.Desc)
}

// Tag 标签
type Tag struct {
	// Category 标签类别
	Category string `json:"category,omitempty"`
	// Value 标签值
	Value string `json:"value,omitempty"`
	// Color 标签颜色
	Color string `json:"color,omitempty"`
	// Shape 标签形状，例如：空心矩形
	Shape string `json:"shape,omitempty"`
}

func (t *Tag) Normalize() {
	utils.PstrTrims(&t.Category, &t.Value, &t.Color, &t.Shape)
}

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
}

// CloneRef 克隆商品引用
func (c *Commodity) CloneRef() RefCommodity {
	return RefCommodity{
		CommodityHash:     c.Hash,
		CommodityName:     c.Name,
		CommodityDesc:     c.Desc,
		CommoditySpec:     c.Spec,
		CommoditySize:     c.Size,
		CommodityValidity: c.Validity,
		CommodityPrice:    c.Price,
		CommodityAttrs:    c.Attrs,
	}
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
	Count RangeX[uint32] `json:"count,omitempty" where:"range,count,omitempty"`
	// Hash 商品哈希值
	Hash string `json:"hash,omitempty" where:"eq,hash,omitempty"`
}

// Normalize 归一化查询商品参数
func (q *QueryCommodity) Normalize() {
	utils.PstrTrims(&q.Name, &q.Desc, &q.Spec, &q.Size, &q.Hash)
}

// SimpleCompany 简单公司信息
type SimpleCompany struct {
	ID   string `json:"id" xorm:"char(36) unique not null 'id'"`
	Name string `json:"name" xorm:"varchar(100) 'name'"`
}

// Normalize 归一化企业信息
func (c *SimpleCompany) Normalize() {
	utils.PstrTrims(&c.Name, &c.Name)
}

// TableName 表名
func (c *SimpleCompany) TableName() string {
	return featc.GetTableName(featc.CompanyCompany)
}
