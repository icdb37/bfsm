package model

import (
	"time"

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

// Commodity 商品
type Commodity struct {
	Xid      uint32           `json:"xid,omitempty" xorm:"pk autoincr 'xid'"`
	ID       string           `json:"id,omitempty" xorm:"varchar(50) unique not null 'id'"`
	Name     string           `json:"name" xorm:"varchar(100) 'name'" validate:"required" cfpx:"name"`
	Desc     string           `json:"desc" xorm:"varchar(200) 'desc'" validate:"required" cfpx:"desc"`
	Spec     string           `json:"spec" xorm:"varchar(100) 'spec'" validate:"required" cfpx:"spec"`
	Size     string           `json:"size" xorm:"varchar(100) 'size'" validate:"required" cfpx:"size"`
	Validity uint16           `json:"validity" xorm:"tinyint 'validity'" validate:"required"`
	Price    uint32           `json:"price" xorm:"int 'price'" validate:"required"`
	Attrs    []*CommodityAttr `json:"attrs" xorm:"json 'attrs'"`
}

func (c *Commodity) Normalize() {
	utils.PstrTrims(&c.Name, &c.Desc, &c.Spec, &c.Size)
	for _, attr := range c.Attrs {
		utils.PstrTrims(&attr.Name, &attr.Value)
	}
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
}
