// Package model 用户数据类型定义
package model

import (
	"time"

	"github.com/icdb37/bfsm/internal/model"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

const (
	// TableUser 用户表名
	TableUser = "user"
)

// EntireUser 完整用户信息
type EntireUser struct {
	Xid       uint32                `json:"xid" xorm:"pk autoincr 'xid'"`
	ID        string                `json:"id" xorm:"char(36) unique not null 'id'"`
	Name      string                `json:"name"  xorm:"varchar(30) 'name'"`
	Desc      string                `json:"desc"  xorm:"varchar(100) 'desc'"`
	Password  string                `json:"password" xorm:"varchar(30) 'password'"`
	Phone     string                `json:"phone" xorm:"varchar(11) 'phone'"`
	IC        *coModel.IdentityCard `json:"ic" xorm:"json 'ic'"`
	BC        *coModel.BankCard     `json:"bc" xorm:"json 'bc'"`
	Tags      []*coModel.Tag        `json:"tags" xorm:"json 'tags'"`
	Contacts  []*coModel.Contact    `json:"contacts" xorm:"json 'contacts'"`
	CreatedAt time.Time             `json:"created_at" xorm:"created 'created_at'"`
	UpdatedAt time.Time             `json:"updated_at" xorm:"updated 'updated_at'"`
}

func (u *EntireUser) TableName() string {
	return TableUser
}

// Normalize 标准化用户信息
func (u *EntireUser) Normalize() {
	utils.PstrTrims(&u.Name, &u.Desc, &u.Phone, &u.Password)
	ns := []any{&u.IC, &u.BC}
	for i, size := 0, len(u.Contacts); i < size; i++ {
		ns = append(ns, &(u.Contacts[i]))
	}
	for i, size := 0, len(u.Tags); i < size; i++ {
		ns = append(ns, &(u.Tags[i]))
	}
	for _, n := range ns {
		if pn, ok := n.(model.Normalizer); ok {
			pn.Normalize()
		}
	}
}

// QueryUser 查询用户信息
type QueryUser struct {
	// Name 姓名
	Name string `json:"name" where:"regex,name,omitempty"`
	// Desc 备注
	Desc string `json:"desc" where:"regex,desc,omitempty"`
	// Phone 手机号
	Phone string `json:"phone" where:"regex,phone,omitempty"`
	// CreatedAt 范围搜索
	CreatedAt coModel.RangeX[time.Time] `json:"created_at" where:"range,created_at,omitempty"`
	// UpdatedAt 范围搜索
	UpdatedAt coModel.RangeX[time.Time] `json:"updated_at" where:"range,updated_at,omitempty"`
}

func (q *QueryUser) Normalize() {
	utils.PstrTrims(&q.Name, &q.Desc, &q.Phone)
}

// SearchRequest 搜索用户请求
type SearchRequest = coModel.SearchRequest[QueryUser]

// SearchResponse 搜索用户应答
type SearchResponse = coModel.SearchResponse[EntireUser]
