package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// EntireCommodity 商品
type EntireCommodity struct {
	coModel.Commodity `json:",inline" xorm:"extends"`
	CreatedAt         time.Time `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt         time.Time `json:"updated_at" xorm:"updated 'updated_at'"`
}

// TableName 商品表名
func (u *EntireCommodity) TableName() string {
	return featc.GetTableName(featc.CommodityCommodity)
}
func (u *EntireCommodity) GetFeature() string {
	return featc.CommodityCommodity
}

// QueryCommodity 商品查询参数
type QueryCommodity struct {
	coModel.QueryCommodity `json:",inline" where:",,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}
