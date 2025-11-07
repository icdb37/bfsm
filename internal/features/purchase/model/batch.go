package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// ExtraExpense 额外费用
type ExtraExpense struct {
	// Name 名称
	Name string `json:"name" xorm:"varchar(200) 'name'" cfpx:"name"`
	// Desc 备注
	Desc string `json:"desc" xorm:"varchar(500) 'desc'" cfpx:"desc"`
	// Amount 额外费用，分
	Amount int32 `json:"amount" xorm:"int 'amount'"`
}

// QueryPurchase 商品查询参数
type QueryPurchase struct {
	// ID 采购标识
	ID string `json:"id" where:"eq,id,omitempty"`
	// Name 名称
	Name string `json:"name" where:"regex,name,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// StatusCode 状态码，精确匹配
	StatusCode enum.StatusCode `json:"status_code" where:"eq,status_code,omitempty"`
	// TotalAmount 订单金额，分
	TotalAmount *coModel.RangeX[int32] `json:"total_amount" where:"range,total_amount,omitempty"`
	// CompanyName 公司名称，模糊匹配
	CompanyName string `json:"company_name" where:"regex,commodities,omitempty"`
	// CommodityName 商品名称，模糊匹配
	CommodityName string `json:"commodity_name" where:"regex,commodities,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt *coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt *coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}

// UpdateBatchStatus 更新采购订单
type UpdateBatchStatus struct {
	// ID 采购标识
	ID string `json:"id" where:"eq,id,omitempty"`
	// UpdatedAt 更新时间
	UpdatedAt time.Time `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	// StatusCode 状态码
	StatusCode enum.StatusCode `json:"status_code" xorm:"int 'status_code'" cfpx:"status_code"`
}

// SimplePurchase 采购订单
type SimplePurchase struct {
	Xid uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" xorm:"created 'created_at'"`
	// UpdatedAt 更新时间
	UpdatedAt time.Time `json:"updated_at" xorm:"updated 'updated_at'" cfpx:"updated_at"`
	// ID 采购标识
	ID string `json:"id" xorm:"char(36) unique not null 'id'"`
	// Name 采购名称
	Name string `json:"name" xorm:"varchar(200) 'name'" cfpx:"name"`
	// Desc 采购描述
	Desc string `json:"desc" xorm:"varchar(500) 'desc'" cfpx:"desc"`
	// GoodsAmount 商品金额，分
	GoodsAmount int32 `json:"goods_amount" xorm:"int 'goods_amount'" cfpx:"goods_amount"`
	// ExtraAmount 额外金额，分
	ExtraAmount int32 `json:"extra_amount" xorm:"int 'extra_amount'" cfpx:"extra_amount"`
	// TotalAmount 采购订单金额，分
	TotalAmount int32 `json:"total_amount" xorm:"int 'total_amount'" cfpx:"total_amount"`
	// StatusFlow 状态流程
	StatusFlow uint16 `json:"status_flow" xorm:"int 'status_flow'" cfpx:"status_flow"`
	// StatusCode 状态数值
	StatusCode enum.StatusCode `json:"status_code" xorm:"int 'status_code'" cfpx:"status_code"`
}

func (e *SimplePurchase) Normalize() {
	e.GoodsAmount = 0
	e.ExtraAmount = 0
	e.TotalAmount = 0
}

// TableName 采购订单表名
func (*SimplePurchase) TableName() string {
	return featc.GetTableName(featc.PurchaseBatch)
}

// GetFeature 获取功能名称
func (*SimplePurchase) GetFeature() string {
	return featc.PurchaseBatch
}

// PurchaseCompany 企业采购
type PurchaseCompany struct {
	// Goods 企业商品列表
	Goods []*coModel.Goods `json:"goods" xorm:"json 'goods'" cfpx:"goods"`
	// Company 企业信息
	Company coModel.RefCompany `json:"company" xorm:"json 'company'" cfpx:"company"`
	// GoodsAmount 企业商品金额，分
	GoodsAmount int32 `json:"goods_amount" xorm:"int 'goods_amount'"`
	// ExtraAmount 企业额外费用金额，分
	ExtraAmount int32 `json:"extra_amount" xorm:"int 'extra_amount'"`
	// TotalAmount 企业采购总金额，分
	TotalAmount int32 `json:"total_amount" xorm:"int 'total_amount'"`
}

// PurchaseBatch 采购批次
type PurchaseBatch struct {
	SimplePurchase `json:",inline" xorm:"extends"`
	// Extras 额外费用列表
	Extras []*ExtraExpense `json:"extras" xorm:"json 'extras'" cfpx:"extras"`
	// Companies 企业采购列表
	Companies []*PurchaseCompany `json:"companies" xorm:"json 'companies'"`
}

func (e *PurchaseBatch) GetBatch() coModel.RefBatch {
	return coModel.RefBatch{
		BatchID:    e.ID,
		BatchName:  e.Name,
		BatchDesc:  e.Desc,
		SourceCode: enum.SourceCodePurchaseProduce,
	}
}

// Normalize -
func (e *PurchaseBatch) Normalize() {
	e.SimplePurchase.Normalize()
	for _, t := range e.Extras {
		e.ExtraAmount += t.Amount
	}
	for _, c := range e.Companies {
		c.GoodsAmount = 0
		for _, g := range c.Goods {
			g.Amount = g.Count * g.Price
			c.GoodsAmount += g.Amount
		}
		c.TotalAmount = c.GoodsAmount + c.ExtraAmount
		e.ExtraAmount += c.ExtraAmount // 采购额外费用金额
		e.GoodsAmount += c.GoodsAmount // 采购商品费用金额
	}
	e.TotalAmount = e.GoodsAmount + e.ExtraAmount // 采购总金额
}
