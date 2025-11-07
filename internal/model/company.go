package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/utils"
)

// CompanyCommodity 商品
type CompanyCommodity struct {
	// Xid 主键
	Xid uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	// ID 商品ID
	ID string `json:"id" xorm:"char(36) unique not null 'id'"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" xorm:"created 'created_at'"`
	// UpdatedAt 更新时间
	UpdatedAt time.Time `json:"updated_at" xorm:"updated 'updated_at'"`
	// Commodity 商品
	Commodity `json:",inline" xorm:"extends"`
	// Company 公司代号
	Company string `json:"company" xorm:"varchar(36) not null index('idx_company') 'company'"`
}

// TableName 商品表名
func (u *CompanyCommodity) TableName() string {
	return featc.GetTableName(featc.CompanyCommodity)
}
func (u *CompanyCommodity) GetFeature() string {
	return featc.CompanyCommodity
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

// RefCompany 引用企业
type RefCompany struct {
	CompanyID   string `json:"company_id" xorm:"char(36) 'company_id'"`
	CompanyName string `json:"company_name" xorm:"varchar(100) 'company_name'"`
}

// Normalize -
func (r *RefCompany) Normalize() {
	utils.PstrTrims(&r.CompanyID, &r.CompanyName)
}
