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
	Hash     string           `json:"hash" xorm:"char(32) 'commodity_hash'"`
	Name     string           `json:"name" xorm:"varchar(100) 'commodity_name'" cfpx:"name"`
	Desc     string           `json:"desc" xorm:"varchar(200) 'commodity_desc'" cfpx:"desc"`
	Spec     string           `json:"spec" xorm:"varchar(100) 'commodity_spec'" cfpx:"spec"`
	Size     string           `json:"size" xorm:"varchar(100) 'commodity_size'" cfpx:"size"`
	Validity int32            `json:"validity" xorm:"tinyint 'commodity_validity'" cfpx:"validity"`
	Price    int32            `json:"price" xorm:"int 'commodity_price'" cfpx:"price"`
	Count    int32            `json:"count" xorm:"count 'commodity_count'" cfpx:"count"`
	Attrs    []*CommodityAttr `json:"attrs" xorm:"json 'commodity_attrs'"`
}

// Normalize -
func (r *RefCommodity) Normalize() {
	utils.PstrTrims(&r.Name, &r.Desc, &r.Spec, &r.Size)
	for _, a := range r.Attrs {
		a.Normalize()
	}
	r.Hash = utils.Hash(r.Name, r.Spec, r.Size)
}

// RefQueryCommodity - 查询商品
type RefQueryCommodity struct {
	// Name 姓名
	Name string `json:"name,omitempty" where:"regex,commodity_name,omitempty"`
	// Desc 备注
	Desc string `json:"desc,omitempty" where:"regex,commodity_desc,omitempty"`
	// Spec 规格
	Spec string `json:"spec,omitempty" where:"regex,commodity_spec,omitempty"`
	// Size 尺寸
	Size string `json:"size,omitempty" where:"regex,commodity_size,omitempty"`
	// Count 数量
	Count RangeX[uint32] `json:"count,omitempty" where:"range,commodity_count,omitempty"`
	// Hash 商品哈希值
	Hash string `json:"hash,omitempty" where:"eq,commodity_hash,omitempty"`
}

// Normalize -
func (q *RefQueryCommodity) Normalize() {
	utils.PstrTrims(&q.Name, &q.Desc, &q.Spec, &q.Size, &q.Hash)
}
