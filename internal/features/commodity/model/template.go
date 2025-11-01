package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

const (
	TableTemplate = "commodity_template"
)

// EntireTemplate 商品
type EntireTemplate struct {
	Xid         uint32               `json:"xid" xorm:"pk autoincr 'xid'"`
	ID          string               `json:"id" xorm:"char(36) unique not null 'id'"`
	Name        string               `json:"name" xorm:"varchar(100) 'name'" validate:"required" cfpx:"name"`
	Desc        string               `json:"desc" xorm:"varchar(200) 'desc'" validate:"required" cfpx:"desc"`
	CreatedAt   time.Time            `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt   time.Time            `json:"updated_at" xorm:"updated 'updated_at'"`
	Commodities []*coModel.Commodity `json:"commodities" xorm:"json 'commodities'"`
}

// TableName 商品表名
func (u *EntireTemplate) TableName() string {
	return TableTemplate
}
func (u *EntireTemplate) GetFeature() string {
	return featc.CommodityTemplate
}

// QueryCommodity 商品查询参数
type QueryTemplate struct {
	// TemplateID 模板ID
	TemplateID string `json:"template_id"  where:"eq,template_id,omitempty"`
	// Name 姓名
	Name string `json:"name" where:"regex,name,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}
