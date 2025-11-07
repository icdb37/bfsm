package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// PurchaseGoods 采购商品
type PurchaseGoods struct {
	// Xid 主键
	Xid uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	// ID 商品ID
	ID string `json:"id" xorm:"char(36) unique not null 'id'"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" xorm:"created 'created_at'"`
	// PurchaseID 采购标识
	PurchaseID string `json:"purchase_id" xorm:"char(36) unique not null 'purchase_id'"`
	// PurchaseName 采购名称
	PurchaseName string `json:"purchase_name" xorm:"varchar(200) 'purchase_name'" cfpx:"name"`
	// Goods 商品信息
	coModel.Goods `json:"goods" xorm:"extends 'goods'" cfpx:"goods"`
	// Company 企业信息
	coModel.SimpleCompany `json:"company" xorm:"extends 'company'" cfpx:"company"`
}

// TableName 表名
func (*PurchaseGoods) TableName() string {
	return featc.GetTableName(featc.PurchaseGoods)
}

// GetFeature 特征
func (*PurchaseGoods) GetFeature() string {
	return featc.PurchaseGoods
}

// QueryGoods 商品查询参数
type QueryGoods struct {
	// StatusCode 状态码，精确匹配
	StatusCode enum.StatusCode `json:"status_code" where:"eq,status_code,omitempty"`
	// TotalAmount 订单金额，分
	TotalAmount coModel.RangeX[int32] `json:"total_amount" where:"range,total_amount,omitempty"`
	// CompanyName 公司名称，模糊匹配
	CompanyName string `json:"company_name" where:"regex,commodities,omitempty"`
	// CommodityName 商品名称，模糊匹配
	CommodityName string `json:"commodity_name" where:"regex,commodities,omitempty"`
	// PurchaseName 采购名称，模糊匹配
	PurchaseName string `json:"purchase_name" where:"regex,purchase_name,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
	// Name 商品名称
	Name string `json:"name" where:"regex,name,omitempty"`
	// Desc 商品备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// Spec 商品规格
	Spec string `json:"spec" where:"regex,spec,omitempty"`
	// Size 商品大小
	Size string `json:"size" where:"regex,size,omitempty"`
	// Count 商品数量
	Count coModel.RangeX[int32] `json:"count" where:"range,count,omitempty"`
	// Price 商品价格
	Price coModel.RangeX[int32] `json:"price" where:"range,price,omitempty"`
}
