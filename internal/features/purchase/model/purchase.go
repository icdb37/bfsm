package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// ExtraCost 额外费用
type ExtraCost struct {
	Name string `json:"name" xorm:"varchar(200) 'name'" cfpx:"name"`
	Desc string `json:"desc" xorm:"varchar(500) 'desc'" cfpx:"desc"`
	Cost int32  `json:"cost" xorm:"int 'cost'"` //额外费用，分
}

// CompanyCommodity 公司商品
type CompanyCommodity struct {
	Commodities []*coModel.ProduceCommodity `json:"commodities" xorm:"json 'commodities'" cfpx:"commodities"` //商品费用
	Extras      []*ExtraCost                `json:"extras" xorm:"json 'extras'" cfpx:"extras"`                //额外费用
	TotalAmount uint32                      `json:"total_amount" xorm:"int 'total_amount'"`                   //总共金额，分
	ClearAmount uint32                      `json:"clear_amount" xorm:"int 'clear_amount'"`                   //已结算金额，分
}

// EntirePurchase 采购订单
type EntirePurchase struct {
	Xid             uint32              `json:"xid" xorm:"pk autoincr 'xid'"`
	CreatedAt       time.Time           `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt       time.Time           `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	ID              string              `json:"id" xorm:"char(36) unique not null 'id'"`
	Name            string              `json:"name" xorm:"varchar(200) 'name'" cfpx:"name"`                            //订单名称
	Desc            string              `json:"desc" xorm:"varchar(500) 'desc'" cfpx:"desc"`                            //订单描述
	Commodities     []*CompanyCommodity `json:"commodities" xorm:"json 'commodities'" cfpx:"commodities"`               //商品
	CommodityAmount uint32              `json:"commodity_amount" xorm:"int 'commodity_amount'" cfpx:"commodity_amount"` //商品金额，分
	TotalAmount     uint32              `json:"total_amount" xorm:"int 'total_amount'" cfpx:"total_amount"`             //订单金额，分
	ClearAmount     uint32              `json:"clear_amount" xorm:"int 'clear_amount'" cfpx:"clear_amount"`             //已结算金额，分
	StatusFlow      uint16              `json:"status_flow" xorm:"int 'status_flow'" cfpx:"status_flow"`                //状态流程
	StatusCode      enum.StatusCode     `json:"status_code" xorm:"int 'status_code'" cfpx:"status_code"`                //状态码
}

// TableName 采购订单表名
func (*EntirePurchase) TableName() string {
	return featc.GetTableName(featc.PurchasePurchase)
}

// GetFeature 获取功能名称
func (*EntirePurchase) GetFeature() string {
	return featc.PurchasePurchase
}

// QueryPurchase 商品查询参数
type QueryPurchase struct {
	// Name 名称
	Name string `json:"name" where:"regex,name,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// StatusCode 状态码，精确匹配
	StatusCode enum.StatusCode `json:"status_code" where:"eq,status_code,omitempty"`
	// TotalAmount 订单金额，分
	TotalAmount coModel.RangeX[uint32] `json:"total_amount" where:"range,total_amount,omitempty"`
	// CompanyName 公司名称，模糊匹配
	CompanyName string `json:"company_name" where:"regex,commodities,omitempty"`
	// CommodityName 商品名称，模糊匹配
	CommodityName string `json:"commodity_name" where:"regex,commodities,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}
