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
	// Status 状态码，精确匹配
	Status enum.StatusCode `json:"status" where:"eq,status,omitempty"`
	// AmountTotal 订单金额，分
	AmountTotal *coModel.RangeX[int32] `json:"total_amount" where:"range,total_amount,omitempty"`
	// CompanyName 公司名称，模糊匹配
	CompanyName string `json:"company_name" where:"regex,commodities,omitempty"`
	// CommodityName 商品名称，模糊匹配
	CommodityName string `json:"commodity_name" where:"regex,commodities,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt *coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt *coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
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
	// Status 状态数值
	Status enum.StatusCode `json:"status" xorm:"int 'status'" cfpx:"status"`
	// AmountGoods 商品金额，分
	AmountGoods int32 `json:"amount_goods" xorm:"int 'amount_goods'" cfpx:"amount_goods"`
	// AmountExtra 额外金额，分
	AmountExtra int32 `json:"amount_extra" xorm:"int 'amount_extra'" cfpx:"amount_extra"`
	// AmountTotal 采购订单金额，分
	AmountTotal int32 `json:"amount_total" xorm:"int 'amount_total'" cfpx:"amount_total"`
}

func (e *SimplePurchase) Normalize() {
	e.AmountGoods = 0
	e.AmountExtra = 0
	e.AmountTotal = 0
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
	// Desc 企业采购描述
	Desc string `json:"desc" xorm:"varchar(200) 'desc'" cfpx:"desc"`
	// AmountGoods 企业商品金额，分
	AmountGoods int32 `json:"amount_goods" xorm:"int 'amount_goods'"`
	// AmountExtra 企业额外费用金额，分
	AmountExtra int32 `json:"amount_extra" xorm:"int 'amount_extra'"`
	// AmountTotal 企业采购总金额，分
	AmountTotal int32 `json:"amount_total" xorm:"int 'amount_total'"`
	// AmountClear 企业已清金额，分
	AmountClear int32 `json:"amount_clear" xorm:"int 'amount_clear'"`
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
		e.AmountExtra += t.Amount
	}
	for _, c := range e.Companies {
		c.AmountGoods = 0
		for _, g := range c.Goods {
			g.Amount = g.Count * g.Price
			c.AmountGoods += g.Amount
		}
		c.AmountTotal = c.AmountGoods + c.AmountExtra
		e.AmountExtra += c.AmountExtra // 采购额外费用金额
		e.AmountGoods += c.AmountGoods // 采购商品费用金额
	}
	e.AmountTotal = e.AmountGoods + e.AmountExtra // 采购总金额
}
