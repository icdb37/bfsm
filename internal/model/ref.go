package model

import "github.com/icdb37/bfsm/internal/utils"

// RefCompany 引用企业
type RefCompany struct {
	CompanyID   string `json:"company_id" xorm:"char(36) 'company_id'"`
	CompanyName string `json:"company_name" xorm:"varchar(100) 'company_name'"`
}

// Normalize -
func (r *RefCompany) Normalize() {
	utils.PstrTrims(&r.CompanyID, &r.CompanyName)
}

// RefSimpleCommodity 引用基本商品信息
type RefSimpleCommodity struct {
	Hash string `json:"hash" xorm:"char(32) 'commodity_hash'"`
	Name string `json:"name" xorm:"varchar(100) 'commodity_name'" cfpx:"name"`
	Desc string `json:"desc" xorm:"varchar(200) 'commodity_desc'" cfpx:"desc"`
	Spec string `json:"spec" xorm:"varchar(100) 'commodity_spec'" cfpx:"spec"`
	Size string `json:"size" xorm:"varchar(100) 'commodity_size'" cfpx:"size"`
}

// RefCommodity 引用商品
type RefCommodity struct {
	CommodityCount    int32            `json:"commodity_count" xorm:"int 'commodity_count'" cfpx:"count"`
	CommodityPrice    int32            `json:"commodity_price" xorm:"int 'commodity_price'" cfpx:"price"`
	CommodityValidity int32            `json:"commodity_validity" xorm:"tinyint 'commodity_validity'" cfpx:"validity"`
	CommodityHash     string           `json:"commodity_hash" xorm:"char(32) 'commodity_hash'"`
	CommodityName     string           `json:"commodity_name" xorm:"varchar(100) 'commodity_name'" cfpx:"name"`
	CommodityDesc     string           `json:"commodity_desc" xorm:"varchar(200) 'commodity_desc'" cfpx:"desc"`
	CommoditySpec     string           `json:"commodity_spec" xorm:"varchar(100) 'commodity_spec'" cfpx:"spec"`
	CommoditySize     string           `json:"commodity_size" xorm:"varchar(100) 'commodity_size'" cfpx:"size"`
	CommodityAttrs    []*CommodityAttr `json:"commodity_attrs" xorm:"json 'commodity_attrs'"`
}

// Normalize -
func (r *RefCommodity) Normalize() {
	utils.PstrTrims(&r.CommodityName, &r.CommodityDesc, &r.CommoditySpec, &r.CommoditySize)
	for _, a := range r.CommodityAttrs {
		a.Normalize()
	}
	r.CommodityHash = utils.Hash(r.CommodityName, r.CommoditySpec, r.CommoditySize)
}
func (r *RefCommodity) GetHash() string {
	r.CommodityHash = utils.Hash(r.CommodityName, r.CommoditySpec, r.CommoditySize)
	return r.CommodityHash
}

// RefQueryCommodity - 查询商品
type RefQueryCommodity struct {
	// CommodityName 姓名
	CommodityName string `json:"commodity_name,omitempty" where:"regex,commodity_name,omitempty"`
	// CommodityDesc 备注
	CommodityDesc string `json:"commodity_desc,omitempty" where:"regex,commodity_desc,omitempty"`
	// CommoditySpec 规格
	CommoditySpec string `json:"commodity_spec,omitempty" where:"regex,commodity_spec,omitempty"`
	// CommoditySize 尺寸
	CommoditySize string `json:"commodity_size,omitempty" where:"regex,commodity_size,omitempty"`
	// CommodityCount 数量
	CommodityCount RangeX[uint32] `json:"commodity_count,omitempty" where:"range,commodity_count,omitempty"`
	// CommodityHash 商品哈希值
	CommodityHash string `json:"commodity_hash,omitempty" where:"eq,commodity_hash,omitempty"`
}

// Normalize -
func (q *RefQueryCommodity) Normalize() {
	utils.PstrTrims(&q.CommodityName, &q.CommodityDesc, &q.CommoditySpec, &q.CommoditySize, &q.CommodityHash)
}
