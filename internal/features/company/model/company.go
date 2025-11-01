package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

// EntireCompany 完整公司信息
type EntireCompany struct {
	Xid       uint32             `json:"xid" xorm:"pk autoincr 'xid'"`
	ID        string             `json:"id" xorm:"char(36) unique not null 'id'"`
	Name      string             `json:"name" xorm:"varchar(100) 'name'" validate:"required" cfpx:"name"`
	Desc      string             `json:"desc" xorm:"varchar(200) 'desc'" validate:"required" cfpx:"desc"`
	Address   string             `json:"address" xorm:"varchar(200) 'address'" validate:"required"`
	Contacts  []*coModel.Contact `json:"contacts" xorm:"json 'contacts'" cfpx:"contact"`
	CreatedAt time.Time          `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt time.Time          `json:"updated_at" xorm:"updated 'updated_at'"`
}

// TableName 表名
func (c *EntireCompany) TableName() string {
	return featc.GetTableName(featc.CompanyCompany)
}

// GetFeature 特征
func (u *EntireCompany) GetFeature() string {
	return featc.CompanyCompany
}

func (c *EntireCompany) Normalize() {
	utils.PstrTrims(&c.Name, &c.Desc, &c.Address)
	for _, contact := range c.Contacts {
		contact.Normalize()
	}
}

// SimpleCompany 简单公司信息
type SimpleCompany struct {
	Xid  uint32 `json:"xid" xorm:"pk autoincr 'xid'"`
	ID   string `json:"id" xorm:"char(36) unique not null 'id'"`
	Name string `json:"name" xorm:"varchar(100) 'name'" validate:"required"`
}

// TableName 表名
func (c *SimpleCompany) TableName() string {
	return featc.GetTableName(featc.CompanyCommodity)
}

type QueryCompany struct {
	// Name 姓名
	Name string `json:"name" where:"regex,name,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// Address 地址
	Address string `json:"address" where:"regex,address,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}
