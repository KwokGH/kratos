package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Bank struct {
	ID    string `gorm:"column:id;primary_key"`           // 银行代码
	Name  string `gorm:"column:name"`                     // 银行名称
	Logo  string `gorm:"column:logo"`                     // 银行LOGO地址
	Limit int    `gorm:"column:limit;default:0;NOT NULL"` // 限额，单位元
}

func (m *Bank) TableName() string {
	return "bank"
}

type IBankDao interface {
	dbo.DataAccesser
}

type BankDao struct {
	dbo.BaseDA
}

var (
	_BankOnce sync.Once
	_BankDao  IBankDao
)

func GetBankDao() IBankDao {
	_BankOnce.Do(func() {
		_BankDao = &BankDao{}
	})
	return _BankDao
}

type BankCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c BankCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c BankCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c BankCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
