package model

import (
	"time"
)

// Builder 建设者模型
// 修改Builder模型
type Builder struct {
	ID       int64     `json:"id" xorm:"pk autoincr 'id'"`
	Alias    string    `json:"alias" xorm:"varchar(100) 'alias'"`
	Desc     string    `json:"desc" xorm:"text 'desc'"`
	Phone    string    `json:"phone" xorm:"varchar(20) 'phone'"`
	Password string    `json:"password,omitempty" xorm:"varchar(255) 'password'"`
	Created  time.Time `json:"created" xorm:"created 'created'"`
	Updated  time.Time `json:"updated" xorm:"updated 'updated'"`

	// 内联存储，无需额外表
	IC       *IdentityCard `json:"ic" xorm:"json 'ic_data'"`
	BC       *BankCard     `json:"bc" xorm:"json 'bc_data'"`
	Contacts []*Contact    `json:"contacts" xorm:"json 'contacts_data'"`
	Tags     []string      `json:"tags" xorm:"json 'tags'"`
}

// TableName 表名
func (Builder) TableName() string {
	return "builders"
}

// BuilderSearchRequest 建设者搜索请求
type BuilderSearchRequest struct {
	Page  int                 `json:"page"`
	Size  int                 `json:"size"`
	Sorts []string            `json:"sorts"`
	Query *BuilderSearchQuery `json:"query"`
}

// BuilderSearchQuery 建设者搜索条件
type BuilderSearchQuery struct {
	Alias    string             `json:"alias"`
	Desc     string             `json:"desc"`
	Phone    string             `json:"phone"`
	IC       *IdentityCardQuery `json:"ic"`
	BC       *BankCardQuery     `json:"bc"`
	Tags     string             `json:"tags"`
	Contacts *ContactQuery      `json:"contacts"`
	Created  *TimeRangeQuery    `json:"created"`
	Updated  *TimeRangeQuery    `json:"updated"`
}

// IdentityCardQuery 身份证查询条件
type IdentityCardQuery struct {
	No      string    `json:"no"`
	Name    string    `json:"name"`
	Sex     uint8     `json:"sex"`
	Birth   time.Time `json:"birth"`
	Address string    `json:"address"`
}

// BankCardQuery 银行卡查询条件
type BankCardQuery struct {
	No   string `json:"no"`
	Name string `json:"name"`
	Bank string `json:"bank"`
}

// ContactQuery 联系人查询条件
type ContactQuery struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// TimeRangeQuery 时间范围查询
type TimeRangeQuery struct {
	Beg time.Time `json:"beg"`
	End time.Time `json:"end"`
}

// BuilderCreateRequest 建设者创建请求
type BuilderCreateRequest struct {
	Alias    string        `json:"alias" validate:"required"`
	Desc     string        `json:"desc"`
	Phone    string        `json:"phone" validate:"required"`
	Password string        `json:"password" validate:"required"`
	IC       *IdentityCard `json:"ic"`
	BC       *BankCard     `json:"bc"`
	Tags     []string      `json:"tags"`
	Contacts []*Contact    `json:"contacts"`
}

// BuilderUpdateRequest 建设者更新请求
type BuilderUpdateRequest struct {
	Alias    string        `json:"alias" validate:"required"`
	Desc     string        `json:"desc"`
	Phone    string        `json:"phone" validate:"required"`
	Password string        `json:"password,omitempty"`
	IC       *IdentityCard `json:"ic"`
	BC       *BankCard     `json:"bc"`
	Tags     []string      `json:"tags"`
	Contacts []*Contact    `json:"contacts"`
}

// BuilderResponse 建设者响应
type BuilderResponse struct {
	ID      int64     `json:"id"`
	Alias   string    `json:"alias"`
	Desc    string    `json:"desc"`
	Phone   string    `json:"phone"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	IC      *struct {
		No      string `json:"no"`
		Name    string `json:"name"`
		Sex     uint8  `json:"sex"`
		Address string `json:"address"`
	} `json:"ic,omitempty"`
	Tags []string `json:"tags"`
}

// DeleteResponse 删除响应
type DeleteResponse struct {
	ID    int64  `json:"id"`
	Alias string `json:"alias"`
}
